package chap5

import (
	"fmt"
	"testing"
)

type Queue struct {
	elements []interface{}
}

func NewQueue() *Queue {
	return &Queue{
		elements: make([]interface{}, 10),
	}
}

func (*Queue) Push(element interface{}) error {
	panic("not implemented")
}
func (self *Queue) length() int {
	return len(self.elements)
}

type Data struct {
	x int
}

func (self Data) ValueTest() {
	fmt.Printf("Value: %p\n", &self)
}

func (self *Data) PointerTest() {
	fmt.Printf("Pointer: %p\n", self)
}
func TestMethod1(t *testing.T) {
	d := Data{}
	p := &d
	fmt.Printf("Data: %p\n", p)

	d.ValueTest()
	d.PointerTest()

	p.ValueTest()
	p.PointerTest()

}

type User struct {
	id   int
	name string
}
type Manager struct {
	User
	title string
}

func (self *User) ToString() string {
	return fmt.Sprintf("User: %p, %v\n", self, self)

}
func (self *Manager) ToString() string {
	return fmt.Sprintf("Manager: %p, %v\n", self, self)
}
func TestMethod2(t *testing.T) {
	m := Manager{
		User{1, "Alice"},
		"Manager",
	}
	fmt.Printf("Manager: %p\n", &m)
	fmt.Println(m.ToString())
	fmt.Println(m.User.ToString())
}

func (self *User) Test() {
	fmt.Printf("User: %p, %v\n", self, self)
}

func TestMethod3(t *testing.T) {
	u := User{1, "Alice"}
	u.Test() //method value

	mValue := u.Test
	mValue()

	mExpr := (*User).Test //method expression
	mExpr(&u)
}

func (self *User) TestPointer() {
	fmt.Printf("TestPointer: %p, %v\n", self, self)
}

func (self User) TestValue() {
	fmt.Printf("TestValue: %p, %v\n", &self, self)
}

func TestMethod4(t *testing.T) {
	u := User{1, "Alice"}
	fmt.Printf("User: %p, %v\n", &u, u)
	mv := User.TestValue
	mv(u)

	mp := (*User).TestPointer
	mp(&u)

	mp2 := (*User).TestValue
	mp2(&u)
}

type Data2 struct{}

func (Data2) TestValue() {
}

func (*Data2) TestPointer() {}

func TestMethod5(t *testing.T) {
	var p *Data2 = nil
	p.TestPointer()

	(*Data2)(nil).TestPointer()
	(*Data2).TestPointer(nil)

}
