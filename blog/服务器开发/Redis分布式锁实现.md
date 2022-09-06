# 基于 Redis 的分布式锁实现

## 简介

在并发场景下，为了保证数据的一致性，需要通过加锁来实现。单机场景下，多线程修改本地数据，直接通过加锁互斥访问即可实现。但在分布式场景下，如果是多个实例修改一个数据，例如电商中秒杀，可能同时会有多个实例要访问数据库修改商品数量，这时就需要用到分布式锁，让多个实例互斥访问数据库。

### 本地锁

使用go自带的互斥锁实现多线程数据修改。

```go
package main
import (
	"fmt"
	"sync"
)
type Count struct {
	num int
	mtx sync.Mutex
}

func (c *Count) Change(v int) {
	c.mtx.Lock()
	c.num += v
	c.mtx.Unlock()
}
func main() {
	cnt := Count{num: 0}
	ch := make(chan int, 10)
	N := 1000
	for i := 0; i < N; i++ {
		go func() {
			cnt.Change(1)
			ch <- 1 //线程完成任务后通知主线程
		}()
	}
	for i := 0; i < N; i++ {
		<-ch //每次从chan中拿出一个，拿完N个时，新建的所有线程都执行完毕。
		//注意此时新建线程不一定都结束了，而是线程的内容执行完了
	}
	fmt.Println(cnt)
	/* !+output
	不加锁的情况下：995左右波动（可以自己运行尝试）
	加锁情况下： 1000
	!-output
	*/
}
```

### 分布式锁

分布式锁就是在分布式环境下抢占一把锁，是实现分布式系统同步访问共享资源的一种方式。

分布式锁相比本地锁的难点在于需要考虑网络。

基于 Redis 分布式锁需要考虑的问题：

+   加锁失败/或超时怎么处理？
+   用户应该只能释放自己加的锁，如何实现
+   分布式锁的过期时间怎么设置

## 实现

### 加锁

基于 Redis 中的 setnx(set if not exists) 来实现加锁，setnx 命令在指定的 key 不存在时，为 key 设置指定的值，设置成功，返回 1 。 设置失败，返回 0 。

使用key作为锁名，同时为了保证只能释放自己加的锁，将val设置为token，只用加锁的用户知道，在释放锁时，如果用户提供的token与Redis中设置的token不同则无法释放锁。

```go
var client *redis.Client

// Lock 实现分布式加锁
// key代表加锁的名称，token是该锁的一个令牌、标识，用来区分是谁加的锁 可以使用 uuid / 生成随机字符串等标识
func Lock(key, token string, expiration time.Duration) (bool, error) {
	ok, err := client.SetNX(key, token, expiration).Result()
	if err != nil {
		ok, err := client.SetNX(key, token, expiration).Result() //如果失败再尝试一次
		return ok, err                                           //无论尝试是否成功，直接返回结果
	}
	return ok, nil
}
```

### 释放

获取要释放锁的token，与用户提供的token对比，如果相同释放锁，不同则不释放锁，返回error

```go
// Unlock 实现分布式解锁
// 解锁需要指定锁名和token，通过token可以辨别当前锁是谁持有，不能释放其他人的锁
func Unlock(key, token string) error {
	val, err := client.Get(key).Result() //查询锁key对应的token
	if err != nil {
		return err
	}
	if val != token { //判断是否是该用户加的锁
		return errors.New("解锁失败，不能释放其他人的锁")
	}
	client.Del(key) // 将锁释放
	return nil
}
```

这里有一个问题，当判断完锁是当前用户持有时（第8行），还没来得及执行释放锁，恰好锁超时过期啦，且其他用户获取了该锁，这时释放锁就会导致释放了他人的锁。

要避免这种情况，可以将查询、判断、释放修改为原子操作。考虑到Redis是单线程服务，使用Redis脚本实现以上三个操作，达到原子效果。

```go
func Unlock1(key, token string) error {
	UnlockScript := redis.NewScript(`
		if redis.call("GET", KEYS[1]) = KEYS[2] then
			return redis.call("del", KEYS[1]) else 
			return 0 
		end`) //新建一个Redis脚本，由于Redis是单线程，所以脚本中的内容一定可以原子执行。

	n, err := UnlockScript.Run(client, []string{key, token}).Result()
	if err != nil {
		return err
	}
	if n == 1 {
		return nil
	}
	return errors.New("解锁失败")
}
```

### 续约

通常无法确定一个锁的超时时间，同时也不能设置锁的超时时间为无限长，即没有超时时间，这样当程序异常崩溃时，锁将无法被获取。因此通过续约机制来实现锁的超时时间动态延长。

```go
// 续约机制
func AutoRefresh(key, token string, expiration time.Duration, interval time.Duration) {
	go func() {
		ticker := time.NewTicker(interval)
		for { //一直循环
			select {
			case <-ticker.C:
				val, err := client.Get(key).Result()
				if err != nil || val != token { //判断是否是该用户加的锁
					return //如果不是证明用户的锁已经释放
				}
				_, err = client.Expire(key, expiration).Result() //修改过期时间
				if err != nil {
					return
				}
			}
		}
	}()
}
```

以上就是基于Redis实现的一个简单分布式锁。
