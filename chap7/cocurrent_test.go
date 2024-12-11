package chap7

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"runtime"
	"sync"
	"testing"
	"time"
)

func sum(id int) {
	var x int64
	for i := 0; i < math.MaxUint32; i++ {
		x += int64(i)
	}
	fmt.Println(id, x)
}
func TestConcurrent(t *testing.T) {
	wg := new(sync.WaitGroup)
	wg.Add(2)

	for i := 0; i < 2; i++ {
		go func(id int) {
			defer wg.Done()
			sum(id)
		}(i)
	}
	wg.Wait()
}

func TestCocurrent2(t *testing.T) {
	wg := new(sync.WaitGroup)
	wg.Add(1)

	go func() {
		defer wg.Done()
		defer fmt.Println("A.defer")

		func() {
			defer fmt.Println("B.defer")
			runtime.Goexit()
			fmt.Println("B")

		}()
		fmt.Println("A")

	}()
	wg.Wait()
}

func TestCocurrent3(t *testing.T) {
	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 6; i++ {
			fmt.Println(i)
			if i == 3 {
				runtime.Gosched()
			}
		}
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Hello, world!")
	}()
	wg.Wait()
}

func TestCocurrent4(t *testing.T) {
	data := make(chan int)
	exit := make(chan bool)

	go func() {
		for d := range data {
			fmt.Println(d)
		}
		fmt.Println("recv over")
		exit <- true
	}()

	data <- 1
	data <- 2
	data <- 3
	close(data)
	fmt.Println("send over")
	<-exit
}

func TestCocurrent5(t *testing.T) {
	data := make(chan int, 3)
	exit := make(chan bool)
	data <- 1
	data <- 2
	data <- 3
	go func() {
		for d := range data {
			fmt.Println(d)
			fmt.Printf("len(data): %d, cap(data): %d\n", len(data), cap(data))
		}
		exit <- true
	}()
	data <- 4
	data <- 5
	data <- 6
	close(data)
	// data <- 7,向closed channel发送数据会panic
	<-exit
}

func TestCocurrent6(t *testing.T) {
	c := make(chan int, 2)
	var send chan<- int = c
	var recv <-chan int = c
	send <- 1
	// <-send, send是只写channel，不能接收数据
	<-recv
	// recv <- 2，recv是只读channel，不能发送数据
	// d := (chan int)(send) send是只写channel，不能转换为读写channel
	// f := (chan int)(recv) recv是只读channel，不能转换为读写channel
}

func TestCocurrent7(t *testing.T) {
	a, b := make(chan int, 3), make(chan int)

	go func() {
		v, ok, s := 0, false, ""
		for {
			select {
			case v, ok = <-a:
				s = "a"
			case v, ok = <-b:
				s = "b"
			}
			if ok {
				fmt.Println(s, v)
			} else {
				os.Exit(0)
			}
		}

	}()
	for i := 0; i < 5; i++ {
		select {
		case a <- i:
		case b <- i:
		}
	}
	close(a)

	select {}
}

func newTest() chan int {
	c := make(chan int)
	rand.Seed(time.Now().UnixNano())
	go func() {
		time.Sleep(time.Second)
		c <- rand.Int()
	}()
	return c
}

func TestCocurrent8(t *testing.T) {
	p := newTest()
	fmt.Println(<-p)
}

func TestCocurrent9(t *testing.T) {
	wg := new(sync.WaitGroup)
	wg.Add(3)

	sem := make(chan int, 1) //信号量，控制协程并发数
	for i := 0; i < 3; i++ {
		go func(id int) {
			defer wg.Done()
			sem <- 1
			for x := 0; x < 3; x++ {
				fmt.Println(id, x)
			}
			<-sem
		}(i)
	}
	wg.Wait()
}

func TestCocurrent10(t *testing.T) {
	var wg sync.WaitGroup
	quit := make(chan bool)

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			task := func() {
				fmt.Println(id, time.Now().Nanosecond())
				time.Sleep(time.Second)
			}
			for {
				select {
				case <-quit:
					return
				default:
					task()
				}
			}
		}(i)
	}
	time.Sleep(5 * time.Second) //模拟业务处理时间
	close(quit)                 //finish all tasks
	wg.Wait()
}

func TestCocurrent11(t *testing.T) {
	w := make(chan bool)
	c := make(chan int, 2)
	go func() {

		select {
		case v := <-c:
			fmt.Println(v)
		case <-time.After(time.Second * 3):
			fmt.Println("timeout")

		}
		w <- true

	}()
	c <- 1

	<-w
}

type Request struct {
	data []int
	ret  chan int //channel作为结构体字段
}

func NewRequest(data ...int) *Request {
	return &Request{data, make(chan int, 1)}
}

func Process(req *Request) {
	sum := 0
	for _, v := range req.data {
		sum += v
	}
	req.ret <- sum
}

func TestCocurrent12(t *testing.T) {
	req := NewRequest(1, 2, 3)
	go Process(req)
	fmt.Println(<-req.ret)
}
