package chap3

import (
	"errors"
	"fmt"
	"sync"
	"testing"
)

func TestMultiReturn(t *testing.T) {

	a, b := MultiReturnTest(10, 20, "RET:%d\n")
	fmt.Println(a)
	fmt.Println(b)
}

func TestFNTest(t *testing.T) {
	TestFn()
}

func TestBc(t *testing.T) {
	fmt.Println(BcTest("Sum: %d", 1, 2, 3))
}

func TestSliceBC(t *testing.T) {
	s := []int{1, 2, 7}
	fmt.Println(BcTest("Sum: %d", s...))
}

func TestNameRet(t *testing.T) {
	fmt.Println(NameReturn(2, 5))
}

func TestNameReturnDefer(t *testing.T) {
	fmt.Println(NameReturnDefer(2, 5))

}

func TestNameReturnDefer2(t *testing.T) {
	fmt.Println(NameReturnDefer2(2, 5))
}

func TestAnoyFunc(t *testing.T) {
	fn := func() { fmt.Println("Hello, world!") }
	fn()

	fns := [](func(x int) int){
		func(x int) int { return x + 1 },
		func(x int) int { return x * 2 },
		func(x int) int { return x * x },
	}
	for _, f := range fns {
		fmt.Println(f(3))
	}

	d := struct {
		fn func() string
	}{
		fn: func() string { return "Hello, world!" },
	}
	fmt.Println(d.fn())

	fc := make(chan func() string, 2)
	fc <- func() string { return "2222Hello, world!" }
	fmt.Println((<-fc)())
}

func TestBB(t *testing.T) {
	f := BB()
	f()
}

func TestDeferTest(t *testing.T) {
	DeferTest(0)
}

func TestDeferTest2(t *testing.T) {
	DeferTest2()
}

var lock sync.Mutex

func locktest() {
	lock.Lock()
	lock.Unlock()
}

func testdefer() {
	lock.Lock()
	defer lock.Unlock()
}

func BenchmarkLock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		locktest()
	}
}

func BenchmarkDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testdefer()
	}
}

func TestPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover from panic", r.(string))
		}
	}()
	panic("panic test")
}

func TestRecover(t *testing.T) {
	defer func() {
		fmt.Println(recover())
	}()

	defer func() {
		panic("panic in defer")
	}()
	panic("panic test")
}

func TestPanic2(t *testing.T) {
	PanicTest(1, 0)
}

func Divide(a, b int) (int, error) {
	var ErrDivideByZero = errors.New("divide by zero")
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}

func TestDivide(t *testing.T) {
	a, b := 10, 1
	if _, err := Divide(a, b); err != nil {
		t.Error(err)
	}
}
