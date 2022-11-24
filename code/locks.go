package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

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

func TestString() {
	const (
		ONE = "hello0000"
		TWO = "TOW"
	)
	str := ""
	go func() {
		for {
			time.Sleep(10)
			fmt.Println(str)
		}

	}()
	for i := 1; true; i = 1 - i {
		if i == 0 {
			str = ONE
		} else {
			str = TWO
		}
		time.Sleep(10)
	}
}

func main() {
	TestString()
}
