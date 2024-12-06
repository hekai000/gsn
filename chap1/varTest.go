package chap1

import (
	"fmt"
	"reflect"
	"unsafe"
)

// const x, y int = 1, 2
// const s = "abc"
// const (
// 	a = 1
// 	b = 2
// )

// const (
// 	d = "abc"
// 	e = len(d)
// 	f = unsafe.Sizeof(d)
// )

// const (
// 	_        = iota
// 	KB int64 = 1 << (10 * iota)
// 	MB
// 	GB
// 	TB
// )

const (
	A = iota
	B
	C = "c"
	D
	E = iota
	F
)

type Color int

const (
	RED Color = iota
	BLUE
	GREEN
)

func test(c Color) { fmt.Println(c) }
func ConstTypeTest() {
	c := BLUE
	test(c)

	x := 2
	test(Color(x))

	test(1)
}
func ConstTest() {
	fmt.Println(A, B, C, D, E, F)
}

// 变量声明方式
// var s = "abc"
// var s1 string = "abc"
// var x int

// 一次定义多个变量
// var q, y, t int
// var u, v = 1, 2
// var (
// 	a int
// 	b float32
// )

func Tst() (string, int) {
	// 变量赋值方式
	ss := "def"
	return ss, 1
}

func MultiVar() {
	// 变量赋值
	data, i := [3]int{1, 2, 3}, 0
	i, data[i] = 2, 100
	fmt.Println(i)
	for j := 0; j < len(data); j++ {
		fmt.Println(data[j])
	}
	Tst()
}

//基本类型

//bool , byte, rune, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, float32, float64, complex64, complex128

//复合类型
//array, struct, pointer, function, interface, slice, map, channel, string

//空指针
//nil

//引用类型
//slice, map , channel

//类型：命名(bool, int, string)，未命名(array, slice, map)

func ImVar() {
	a := []int{0, 0, 0}
	a[0] = 1
	fmt.Println(a)

	b := make([]int, 3)
	b[1] = 100
	fmt.Println(b)

	c := new([]int)
	*c = make([]int, 3) //new([]int) 分配了一个指向整型切片的指针，但此时切片本身还未被初始化，即其内部容量和长度均为零
	(*c)[1] = 300
	fmt.Println(*c)

}

func StringTest() {
	s := "abc"
	fmt.Println(s[0] == '\x61', s[1] == '\x62', s[2] == '\x63')

	q := `a
	b\r\n\x00
	c`
	fmt.Println(q)

	t := "中文" + "测试"
	fmt.Println(t)

	pp := "hello, world!"
	pp1 := pp[:5]
	pp2 := pp[7:]
	pp3 := pp[1:5]
	fmt.Println(pp1, "\n", pp2, "\n", pp3)
}

func UnicodeTest() {
	fmt.Printf("%T\n", 'a')

	var c1, c2 rune = '\u6211', '们'
	fmt.Println(c1 == '我', string(c2) == "\xe4\xbb\xac")
}

func ChangeString() {
	s := "abcd"
	bs := []byte(s)
	bs[0] = 'A'
	fmt.Println(string(bs))

	u := "电脑"
	us := []rune(u)
	us[1] = '话'
	fmt.Println(string(us))
}

func TraverseString() {
	s := "abc汉字"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%d %c\n", i, s[i])
	}
	for _, c := range s {
		fmt.Printf("%c\n", c)
	}
}

func PointerTest() {
	type data struct{ a int }
	var d = data{1234}
	var p *data
	p = &d
	fmt.Printf("%p, %v\n", p, p.a)
	x := 0x12345678
	q := unsafe.Pointer(&x)
	n := (*[4]byte)(q)
	for i := 0; i < len(n); i++ {
		fmt.Printf("%x %p\n", n[i], &n[i])
	}
}

func PointerArithmetic() {
	d := struct {
		s string
		x int
	}{"abc", 100}

	p := uintptr(unsafe.Pointer(&d))
	fmt.Printf("%v\n", p)
	fmt.Println(unsafe.Offsetof(d.x))
	p += unsafe.Offsetof(d.x)
	fmt.Printf("%v\n", p)

	p2 := unsafe.Pointer(p)
	px := (*int)(p2)
	*px = 200

	fmt.Printf("%#v\n", d)
}

func UnnamedStruct() {
	var a struct {
		x int `a`
	}
	var b struct {
		x int `b`
	}
	fmt.Println(reflect.DeepEqual(a, b))

	type bigint int64
	var y bigint = 100
	fmt.Println(y)

	z := 1234
	var bb bigint = bigint(z)
	var bb2 int64 = int64(bb)
	fmt.Println(bb2)
	type myslice []int
	var s myslice = []int{1, 2, 3}
	var s2 []int = s
	fmt.Println(s2)

}

func StructEx() {
	type person struct {
		name string
		age  int
	}
	p := person{"Alice", 25}
	fmt.Println(p.name, p.age)
	p.age = 30
	fmt.Println(p.name, p.age)

}
