package main

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis"
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

func localLock() {
	cnt := Count{num: 0}
	ch := make(chan int, 10)
	N := 1000
	for i := 0; i < N; i++ {
		go func() {
			cnt.Change(1)
			ch <- 1 //协程完成任务通知主线程
		}()
	}
	for i := 0; i < N; i++ {
		<-ch //每次从chan中拿出一个，拿完N个时，新建的所有协程都执行完毕。
		//注意此时新建协程不一定都结束了，而是协程的内容执行完了
	}
	fmt.Println(cnt)
	/* !+output
	不加锁的情况下：995左右波动（可以自己运行尝试）
	加锁情况下： 100
	!-output
	*/
}

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

// Unlock 实现分布式解锁
// 解锁需要指定锁名和token，通过token可以辨别当前锁是谁持有，不能释放其他人的锁
func Unlock(key, token string) error {
	val, err := client.Get(key).Result()
	if err != nil {
		return err
	}
	if val != token { //判断是否是该用户加的锁
		return errors.New("解锁失败，不能释放其他人的锁")
	}
	client.Del(key) // 将锁释放
	// 这里有一个问题，当我判断完锁是我持有时，还没来得及执行释放锁，恰好锁超时过期啦，且另外获取了该锁，这时释放锁就会导致释放了他人的锁
	return nil
}

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

func main() {

}
