package chap4

import (
	"fmt"
	"testing"
)

func TestData(t *testing.T) {
	a := [3]int{1, 2}           //初始化，未初始化的为0
	b := [...]int{1, 2, 3, 4}   //通过初始值确定数组长度
	c := [5]int{2: 100, 4: 200} //通过索引赋值

	d := [...]struct {
		name string
		age  uint8
	}{
		{"Alice", 20},
		{"Bob", 25},
	}
	fmt.Println(a, b, c, d)
	e := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	f := [...][2]int{{1, 1}, {2, 2}, {3, 3}}
	fmt.Println(e, f)
}
func test(x [2]int) {
	fmt.Printf("x: %p\n", &x)
	x[1] = 1000
}
func TestData2(t *testing.T) {

	a := [2]int{}
	fmt.Printf("a: %p\n", &a)
	test(a)
	fmt.Println(a) //值拷贝，地址不同
}

func TestData3(t *testing.T) {
	data := [...]int{0, 1, 2, 3, 4, 5}
	s := data[2:4]
	s[0] += 100
	s[1] += 200
	fmt.Println(s)
	fmt.Println(data)
}

func TestSlice(t *testing.T) {
	s1 := []int{0, 1, 2, 3, 8: 100}
	fmt.Println(s1, len(s1), cap(s1))

	s2 := make([]int, 6, 8) //通过make函数创建切片，指定长度和容量
	fmt.Println(s2, len(s2), cap(s2))

	s3 := make([]int, 6) //省略cap
	fmt.Println(s3, len(s3), cap(s3))

	s4 := []int{0, 1, 2, 3, 4, 5}
	p := &s4[2]
	*p += 100
	fmt.Println(s4)

	data := [][]int{
		[]int{1, 2, 3},
		[]int{100, 200},
		[]int{300, 400, 500, 600},
	}
	for _, row := range data {
		for _, col := range row {
			fmt.Printf("%d ", col)
		}
		fmt.Println()
	}
}

func TestSlice2(t *testing.T) {
	data := [5]struct {
		x int
	}{}
	s := data[:]
	data[1].x = 10
	s[2].x = 20

	fmt.Println(data)
	fmt.Printf("s: %p, %p\n", &data, &data[0])
}

func TestReslice(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	s1 := s[2:5]
	s1[2] = 100
	fmt.Println(s1)
	fmt.Println(s)
	s2 := s1[2:6]
	s2[3] = 200
	fmt.Println(s2)
	fmt.Println(s1)
	fmt.Println(s) //新对象依旧指向底层数组
}

func TestAppend(t *testing.T) {
	s := make([]int, 0, 5)
	fmt.Printf("%p\n", &s)

	s2 := append(s, 1)
	fmt.Printf("%p\n", &s2)

	fmt.Println(s, s2)

	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s4 := data[:3]
	s5 := append(s4, 100, 200)
	fmt.Println(data)
	fmt.Println(s4)
	fmt.Println(s5)
}

func TestSlice3(t *testing.T) {
	s := make([]int, 0, 1)
	c := cap(s)
	//通过2倍容量扩容
	for i := 0; i < 50; i++ {
		s = append(s, i)
		fmt.Printf("i: %d, len: %d, cap: %d\n", i, len(s), cap(s))
		if n := cap(s); n > c {
			fmt.Printf("cap: %d -> %d\n", c, n)
			c = n
		}
	}
}

func TestMap(t *testing.T) {
	m := map[int]struct {
		name string
		age  int
	}{
		1: {"Alice", 20},
		2: {"Bob", 25},
	}

	fmt.Println(m[1].name)
}

func TestMap2(t *testing.T) {
	m := map[string]int{"a": 1}
	if v, ok := m["a"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("key not found")
	}

	fmt.Println(m["b"])

	m["b"] = 2

	delete(m, "c")

	fmt.Println(len(m))
	for k, v := range m {
		fmt.Printf("%s: %d\n", k, v)
	}
}

func TestMap3(t *testing.T) {
	type user struct {
		name string
	}
	m := map[int]user{
		1: {"user1"},
	}
	u := m[1]
	u.name = "user2"
	m[1] = u
	fmt.Println(m)

	m2 := map[int]*user{
		1: &user{"user1"},
	}
	m2[1].name = "Tom"
	fmt.Println(m2[1].name)
}

func TestMap4(t *testing.T) {
	for i := 0; i < 5; i++ {
		m := map[int]string{
			0: "a", 1: "a", 2: "a", 3: "a", 4: "a",
			5: "a", 6: "a", 7: "a", 8: "a", 9: "a",
		}
		for k := range m {
			m[k+k] = "x"
			delete(m, k)
		}
		fmt.Println(m)
	}
}

func TestStruct(t *testing.T) {
	type Node struct {
		_    int
		id   int
		data *byte
		next *Node
	}
	n1 := Node{
		id:   1,
		data: nil,
	}
	n2 := Node{
		id:   2,
		data: nil,
		next: &n1,
	}
	fmt.Printf("n1: %p, %v, n2: %p, %v\n", &n1, n1, &n2, n2)
}

func TestStruct2(t *testing.T) {
	type File struct {
		name string
		size int64
		attr struct {
			perm  int
			owner int
		}
	}
	f := File{
		name: "test.txt",
		size: 1025,
		// attr: {0755, 1},
	}
	f.attr.perm = 0644
	f.attr.owner = 1000
	fmt.Println(f)

	var attr = struct {
		perm  int
		owner int
	}{2, 0755}
	f.attr = attr
	fmt.Println(f)
}

func TestStruct3(t *testing.T) {
	type User struct {
		name string
	}
	type Manager struct {
		User
		title string
	}

	m := Manager{
		User:  User{"Tom"},
		title: "Administrator",
	}
	fmt.Println(m)

}

func TestStruct4(t *testing.T) {

	type Resource struct {
		id int
	}
	type User struct {
		Resource
		name string
	}
	type Manager struct {
		User
		title string
	}
	var m Manager
	m.id = 1
	m.name = "Tom"
	m.title = "Manager"
	fmt.Println(m)
}

func TestStruct5(t *testing.T) {
	type Resource struct {
		id   int
		name string
	}
	type Classify struct {
		id int
	}
	type User struct {
		Resource
		Classify
		name string
	}
	u := User{
		Resource{1, "Tom"},
		Classify{100},
		"Jerry",
	}
	fmt.Println(u.name)
	fmt.Println(u.Resource.name)

	fmt.Println(u.Classify.id)
}

func TestStruct6(t *testing.T) {
	type User struct {
		id   int
		name string
	}

	type Manager struct {
		User
		title string
	}

	m := Manager{User{1, "Tom"}, "Administrator"}
	var u User = m.User
	fmt.Println(u.name)

}

func appendSlice(s []int, x int) []int {
	s = append(s, x)
	fmt.Printf("s: %p, %v\n", &s, s)
	return s
}
func TestSlice4(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	appendSlice(s, 6)
	fmt.Printf("s: %p, %v\n", &s, s)
}
