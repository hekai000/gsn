package chap9

import (
	"fmt"
	"reflect"
	"testing"
	"time"
	"unsafe"
)

func TestDeep1(t *testing.T) {
	x := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	y := x[1:3:6]

	fmt.Println(y, len(y), cap(y))
}

type data struct {
	x [1024 * 100]byte
}

func test() uintptr {
	p := &data{}
	return uintptr(unsafe.Pointer(p))
}

func TestDeep2(t *testing.T) {
	const N = 10000
	cache := new([N]uintptr)
	for i := 0; i < N; i++ {
		cache[i] = test()
		time.Sleep(time.Microsecond)
	}
}

type User struct {
	Username string
}
type Admin struct {
	User
	title string
}

func TestDeep3(t *testing.T) {
	var u Admin
	p := reflect.TypeOf(u)

	for i, n := 0, p.NumField(); i < n; i++ {
		field := p.Field(i) // Field returns a StructField
		fmt.Println(field.Name, field.Type)
	}
}

func TestDeep4(t *testing.T) {
	u := new(Admin)
	p := reflect.TypeOf(u)
	if p.Kind() == reflect.Ptr {
		p = p.Elem()
	}
	for i, n := 0, p.NumField(); i < n; i++ {
		field := p.Field(i) // Field returns a StructField
		fmt.Println(field.Name, field.Type)
	}
}

type MyUser struct{}
type Admin2 struct {
	MyUser
}

func (*MyUser) ToString() {}

func (Admin2) test1() {}

func TestDeep5(t *testing.T) {
	var u Admin2
	methods := func(r reflect.Type) {
		for i, n := 0, r.NumMethod(); i < n; i++ {
			method := r.Method(i) // Method returns a Method
			fmt.Println(method.Name, method.Type)
		}
	}
	fmt.Println("----Value interface----")
	methods(reflect.TypeOf(u))
	fmt.Println("----Pointer interface----")
	methods(reflect.TypeOf(&u))
}

type MyUser2 struct {
	Username string
	age      int
}

type Admin3 struct {
	MyUser2
	title string
}

func TestDeep6(t *testing.T) {
	var u Admin3
	p := reflect.TypeOf(u)
	f, _ := p.FieldByName("title")
	fmt.Println(f.Name, f.Type)

	f, _ = p.FieldByName("MyUser2")
	fmt.Println(f.Name, f.Type)

	f, _ = p.FieldByName("Username")
	fmt.Println(f.Name, f.Type)

	f = p.FieldByIndex([]int{0, 1})
	fmt.Println(f.Name, f.Type)
}

var (
	Int    = reflect.TypeOf(0)
	String = reflect.TypeOf("")
)

// 从基本类型获取对应复合类型
func TestDeep7(t *testing.T) {
	c := reflect.ChanOf(reflect.SendDir, String)
	fmt.Println(c)

	m := reflect.MapOf(String, Int)
	fmt.Println(m)

	s := reflect.SliceOf(Int)
	fmt.Println(s)

	tt := struct{ Name string }{}
	p := reflect.PtrTo(reflect.TypeOf(tt))
	fmt.Println(p)

	w := reflect.TypeOf(make(chan int)).Elem()
	fmt.Println(w)
}

type Data struct {
}

func (*Data) String() string {
	return ""
}

func TestDeep8(t *testing.T) {
	var d Data
	q := reflect.TypeOf(d)
	fmt.Println(q)
	it := reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	fmt.Println(q.Implements(it))
}

type User3 struct {
	Username string
	age      int
}

func TestDeep9(t *testing.T) {
	u := User3{"Jack", 23}
	v := reflect.ValueOf(u)
	p := reflect.ValueOf(&u)

	fmt.Println(v.CanSet(), v.FieldByName("Username").CanSet())
	fmt.Println(p.CanSet(), p.Elem().FieldByName("Username").CanSet())
}

func TestDeep10(t *testing.T) {
	u := User3{"Jack", 23}
	p := reflect.ValueOf(&u).Elem()
	p.FieldByName("Username").SetString("Tom")

	f := p.FieldByName("age")
	fmt.Println(f.CanSet())

	if f.CanAddr() {
		age := (*int)(unsafe.Pointer(f.UnsafeAddr()))
		*age = 24
	}
	fmt.Println(u, p.Interface().(User3))
}

func TestDeep11(t *testing.T) {
	s := make([]int, 0, 10)
	v := reflect.ValueOf(&s).Elem()

	v.SetLen(2)
	v.Index(0).SetInt(100)
	v.Index(1).SetInt(200)

	fmt.Println(v.Interface(), s)

	v2 := reflect.Append(v, reflect.ValueOf(300))
	v2 = reflect.AppendSlice(v2, reflect.ValueOf([]int{400, 500}))

	fmt.Println(v2.Interface())

	fmt.Println("--------")

	m := map[string]int{"a": 1}
	v = reflect.ValueOf(&m).Elem()

	v.SetMapIndex(reflect.ValueOf("a"), reflect.ValueOf(100))
	v.SetMapIndex(reflect.ValueOf("b"), reflect.ValueOf(200))

	fmt.Println(v.Interface(), m)
}

type Data2 struct {
}

func (*Data2) Test(x, y int) (int, int) {
	return x + 100, y + 100
}
func (*Data2) Sum(s string, x ...int) string {
	c := 0
	for _, n := range x {
		c += n
	}
	return fmt.Sprintf(s, c)
}

func info(m reflect.Method) {
	t := m.Type
	fmt.Println(m.Name)

	for i := 0; i < t.NumIn(); i++ {
		fmt.Printf(" in[%d] %v\n", i, t.In(i))
	}
	for i := 0; i < t.NumOut(); i++ {
		fmt.Printf(" out[%d] %v\n", i, t.Out(i))
	}
}

func TestDeep12(t *testing.T) {
	d := new(Data2)
	v := reflect.ValueOf(d)
	test, _ := v.Type().MethodByName("Test")
	info(test)

	sum, _ := v.Type().MethodByName("Sum")
	info(sum)
}

func TestDeep13(t *testing.T) {
	d := new(Data2)
	v := reflect.ValueOf(d)

	exec := func(name string, in []reflect.Value) {
		m := v.MethodByName(name)
		out := m.Call(in)
		for _, v := range out {
			fmt.Println(v.Interface())
		}

	}

	exec("Test", []reflect.Value{reflect.ValueOf(1), reflect.ValueOf(2)})

	fmt.Println("--------")

	exec("Sum", []reflect.Value{reflect.ValueOf("sum: %d"), reflect.ValueOf(1), reflect.ValueOf(2)})
}

func Make(T reflect.Type, fptr interface{}) {
	swap := func(in []reflect.Value) []reflect.Value {

		return []reflect.Value{
			reflect.MakeSlice(
				reflect.SliceOf(T),
				int(in[0].Int()),
				int(in[1].Int()),
			),
		}
	}

	fn := reflect.ValueOf(fptr).Elem()

	v := reflect.MakeFunc(fn.Type(), swap)
	fn.Set(v)
}

func TestDeep14(t *testing.T) {
	var makeints func(int, int) []int
	var makestrings func(int, int) []string
	Make(Int, &makeints)
	Make(String, &makestrings)

	x := makeints(3, 4)
	fmt.Printf("%#v\n", x)

	s := makestrings(3, 4)
	fmt.Printf("%#v\n", s)

}
