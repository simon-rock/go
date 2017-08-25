package concurrency

import (
	"fmt"
	"sync"
	"time"
)

// 1 goroutine, create_pthread
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
func goroutinetest() {
	go say("world")
	say("hello")
}

// 2 channel
// 默认情况下，在另一端准备好之前，发送和接收都会阻塞。这使得 goroutine 可以在没有明确的锁或竞态变量的情况下进行同步
func sum(a []int, c chan int) {
	sum := 0
	for i, v := range a {
		sum += v
		fmt.Println("thread ", i)
		time.Sleep(100 * time.Millisecond)
	}
	c <- sum
}
func channeltest() {
	a := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	fmt.Println("main")
	x, y := <-c, <-c // 3个线程,其中主线程 从管道中读取数据是阻塞的
	fmt.Println(x, y, x+y)
}

// 3 channel with buffer, buffer is queue
func channeltest_2() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2 // if buffer is full, will block!!!
	fmt.Println(<-ch)
	fmt.Println(<-ch) // if buffer is empty, will block!!!
}

// 4 range and close
// 发送者可以 close 一个 channel 来表示再没有值会被发送了
// 接收者可以通过赋值语句的第二参数来测试 channel 是否被关闭
// v, ok :=<-ch, ok is false if ch has been closed
// 循环 `for i := range c` 会不断从 channel 接收值，直到它被关闭
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}
func closechannel() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c) // 由于管道的读取和放入是阻塞的所以可以第一个参数传入一个比较大值
	for i := range c {
		fmt.Println(i)
	}
}

// 5 select
// select 会阻塞，直到条件分支中的某个可以继续执行，这时就会执行那个条件分支。当多个都准备好的时候，会随机选择一个
func fibonacci_2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
func selecttest() {
	fmt.Println("--select test--")
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci_2(c, quit)
}

// 6 select 3 , defaut branch
// 当 select 中的其他条件分支都没有准备好的时候，default 分支会被执行。
// 为了非阻塞的发送或者接收，可使用 default 分支：
func selecttest_2() {
	fmt.Println("--6 select test--")
	tick := time.Tick(100 * time.Millisecond)  // 定时事件
	boom := time.After(500 * time.Millisecond) // 定时事件
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("  .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

// 7 8 练习 skip
// 9 sync.Mutex
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.v[key]++
}
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}
func mutex_test() {
	fmt.Println("--9 mutex test--")
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 100; i++ {
		go c.Inc("sk")
	}
	time.Sleep(time.Second)
	fmt.Println(c.Value("sk"))
}

// Out is
func Out() {
	// 1
	goroutinetest()
	// 2
	channeltest()
	// 3
	channeltest_2()
	// 4
	closechannel()
	// 5
	selecttest()
	// 6
	selecttest_2()
	// 9
	mutex_test()
}
