GO语言只强人锁难

## 锁的优化

+   减少持有时间

```go
// 临界区很长导致所持有时间也很长
func f(){
    // do something
    mu.Lock()
    defer mu.UnLock()
    ...
    // do many thing 
    ...
    
    // add new
}
// 可行的方法
func f(){
    // do something
    func(){
        mu.Lock()
    	defer mu.UnLock()
        // do something
    }
    ...
    // do many thing 
    ...
    
    // add new
}
```

​	较少临界区的长度从而减少锁的持有时间，达到优化并发的效果

+   优化锁的粒度

    +   分段加锁：map[int]string，定义N个锁，读写第 i 时，向第 i%N 个锁申请加锁。这样能够保证锁同时被 N 个线程并发访问。
    +   无法分段的数据类型：申请多个该数据类型，每个对应一个锁。有点线程池，连接池的感觉。

    ```go
    type Rander struct {
    	pos     uint32
    	randers [16]*rand.Rand
    	locks   [16]*sync.Mutex
    }
    
    func (r *Rander) Intn(n int) int {
    	x := atomic.AddUint32(&r.pos, 1)
    	x = x % 16
    	r.locks[x].Lock()
    	defer r.locks[x].Unlock()
    	n = r.randers[x].Intn(n)
    	return n
    }
    ```

+   读写分离

    +   使用 `sync.RWMutex`

    ```go
    type Counter struct {
    	n int
    	// mu sync.Mutex
    	mu sync.RWMutex
    }
    
    func (c *Counter) Write() {
    	c.mu.Lock()
    	defer c.mu.Unlock()
    	c.n++
    }
    
    func (c *Counter) Read() {
    	// c.mu.Lock()
    	c.mu.RLock()
    	defer c.mu.RUnlock()
    	c.n++
    }
    ```

+   使用原子操作



## 锁的坑

+   **不要**拷贝锁

```go
func Worker(mu sync.Mutex){
    mu.Lock()
    defer mu.UnLock()
    // DO something
}
func main(){
    var mu sync.Mutex
    go Worker(mu)
    go Worker(mu)
}
```

go 语言中参数传递是值传递，两个goroutine中的两个锁是全新的锁。

非要传递锁的话可以使用指针类型传递。

+   Mutex 不可重入
+    Atomic.Value 应存入自读对象
+    race detector 
    +   go run -race main.go > err.log  // 发现 潜在的问题 