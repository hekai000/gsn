package chap3

import "fmt"

func MultiReturnTest(x, y int, s string) (int, string) {
	n := x + y
	return n, fmt.Sprintf(s, n)
}

func testFN(fn func() int) int {
	return fn()
}

type FormatFunc func(s string, x, y int) string

func format(fn FormatFunc, s string, x, y int) string {
	return fn(s, x, y)
}
func TestFn() {
	s1 := testFN(func() int { return 100 })
	s2 := format(func(s string, x, y int) string {
		return fmt.Sprintf(s, x, y)
	}, "AA: %d, BB: %d", 10, 20)

	fmt.Println(s1, "\n", s2)
}
func BcTest(s string, n ...int) string {
	var x int
	for _, i := range n {
		x += i
	}
	return fmt.Sprintf(s, x)
}
func NameReturn(x, y int) (z int) {
	z = x + y
	return
}

func NameReturnDefer(x, y int) (z int) {
	defer func() {
		fmt.Println("Defer called")
		z += 10
	}()
	z = x + y
	return
}

func NameReturnDefer2(x, y int) (z int) {
	defer func() {
		fmt.Println(z)
		z += 10
	}()
	z = x + y
	return 200 + z
}

func BB() func() {
	x := 100
	fmt.Printf("x (%p) = %d\n", &x, x)
	return func() {
		fmt.Printf("x (%p) = %d\n", &x, x)
	}
}

// 闭包怎么理解，匿名函数

//defer 语句的执行顺序: FILO

func DeferTest(x int) {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	defer func() {
		fmt.Println(1 / x)
	}()
	fmt.Println("Hello, world!")
}

func DeferTest2() {
	x, y := 10, 20
	defer func(i int) {
		fmt.Println("defer: ", i, y)
	}(x) // 延迟调用的参数的注册时求值，即x为10
	x += 10
	y += 100
	fmt.Println("x, y: ", x, y)
}

func PanicTest(x, y int) {
	var z int
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in f", r)
				z = 0
			}
		}()
		z = x / y
	}()

	fmt.Println("Result: ", z)
}
