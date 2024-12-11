package chap6

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

type Stringer interface {
	String() string
}

type Printer interface {
	Stringer
	Print()
}

type User struct {
	id   int
	name string
}

func (self *User) String() string {
	return fmt.Sprintf("user: %d, name: %s", self.id, self.name)
}

func (self *User) Print() {
	fmt.Println(self.String())
}

func Print(v interface{}) {
	fmt.Printf("%T: %v\n", v, v)
}
func TestInterface(t *testing.T) {
	var p Printer = &User{1, "Alice"}
	p.Print()
}

func TestInterface2(t *testing.T) {
	Print(1)
	Print("hello, world")
}

type Tester struct {
	s interface {
		String() string
	}
}

type MyUser struct {
	id   int
	name string
}

func (self *MyUser) String() string {
	return fmt.Sprintf("user: %d, name: %s", self.id, self.name)
}

func (self *MyUser) Print() {
	fmt.Println(self.String())
}

func TestInterface3(t *testing.T) {
	var p Printer = &MyUser{1, "Alice"}
	p.Print()
}

func TestInterface4(t *testing.T) {
	u1 := User{1, "Alice"}
	var i interface{} = u1 //数据指针持有的是目标对象的只读复制品，复制完整对象和指针
	u1.id = 2
	u1.name = "Bob"
	fmt.Printf("%v\n", u1)
	fmt.Printf("%v\n", i.(User))
}

func TestInterface5(t *testing.T) {
	u := User{1, "Alice"}
	var vi, pi interface{} = u, &u
	//vi.(User).id = 2 //vi.(User)是只读的，不能修改,接口转型返回临时对象，只有使用指针才能修改其状态
	pi.(*User).name = "Bob"
	fmt.Printf("%v\n", vi.(User))
	fmt.Printf("%v\n", pi.(*User))
}

type iface struct {
	itab, data uintptr
}

func TestInterface6(t *testing.T) {
	var a interface{} = nil         //tab为nil，data为nil
	var b interface{} = (*int)(nil) //tab包含*int类型信息，data为nil

	ia := *(*iface)(unsafe.Pointer(&a))
	ib := *(*iface)(unsafe.Pointer(&b))

	fmt.Println(a == nil, ia)
	fmt.Println(b == nil, ib, reflect.ValueOf(b).IsNil())
}

func TestInterface7(t *testing.T) {
	var o interface{} = &User{1, "Alice"}
	if i, ok := o.(fmt.Stringer); ok {
		fmt.Println(i)
	}
	// u := o.(*User)
	u := o.(*User) //o是(*User)类型
	fmt.Println(u)
}

func TestInterface8(t *testing.T) {
	var o interface{} = &User{1, "Alice"}
	switch v := o.(type) {
	case nil:
		fmt.Println("nil interface")
	case fmt.Stringer:
		fmt.Println("Stringer interface")
		fmt.Println(v)
		//fallthrough,不能使用fallthrough，否则会导致类型断言失败
	case func() string:
		fmt.Println("func() string interface")
		fmt.Println(v())
	case *User:
		fmt.Printf("%d, %s\n", v.id, v.name)
	default:
		fmt.Println("unknown type")
	}
}

type MyPrinter interface {
	String() string
	Print()
}

func TestInterface9(t *testing.T) {
	var p MyPrinter = &User{1, "Alice"}
	var o Stringer = p
	fmt.Println(o.String())
}

type Tester2 interface {
	Do()
}

type FuncDo func()

func (a FuncDo) Do() {
	a()
}

func TestInterface10(t *testing.T) {
	var td Tester2 = FuncDo(func() {
		fmt.Println("Hello, world")
	}) //将一个匿名函数包装为接口实例
	td.Do() //通过调用接口调用该函数
}
