package main

import (
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

func waitGroup() {

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("goroutine done!")
		wg.Done()
	}()

	wg.Wait()
}

func Once() {

	var loadIconsOnce sync.Once
	//func (o *Once) Do(f func()) {}

	f := func() {
		fmt.Println("Once done!")
	}

	loadIconsOnce.Do(f)
	loadIconsOnce.Do(f)
	loadIconsOnce.Do(f)
}

func syncMap() {

	var m = sync.Map{}
	//Store、Load、LoadOrStore、Delete、Range

	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m.Store(key, n)
			value, _ := m.Load(key)
			fmt.Printf("k=:%v,v:=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

var x int64
var l sync.Mutex
var wg sync.WaitGroup

// 普通版加函数
func add() {
	// x = x + 1
	x++ // 等价于上面的操作
	wg.Done()
}

// 互斥锁版加函数
func mutexAdd() {
	l.Lock()
	x++
	l.Unlock()
	wg.Done()
}

// 原子操作版加函数
func atomicAdd() {
	atomic.AddInt64(&x, 1)
	wg.Done()
}

func Atomic() {
	start := time.Now()
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		//go add()       // 普通版add函数 不是并发安全的
		//go mutexAdd()  // 加锁版add函数 是并发安全的，但是加锁性能开销大
		go atomicAdd() // 原子操作版add函数 是并发安全，性能优于加锁版
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(x)
	fmt.Println(end.Sub(start))
}

func main() {
	waitGroup()
	Once()
	syncMap()
	Atomic()
}
