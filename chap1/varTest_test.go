package chap1

import (
	"fmt"
	"testing"
)

func TestMultiVar(t *testing.T) {
	MultiVar()

}

func TestTst(t *testing.T) {
	a, _ := Tst()
	fmt.Println(a)
}

func TestConstTest(t *testing.T) {
	ConstTest()
}

func TestImVar(t *testing.T) {
	ImVar()
}

func TestConstType(t *testing.T) {
	ConstTypeTest()
}

func TestStringTest(t *testing.T) {
	StringTest()
}

func TestUnicodeTest(t *testing.T) {
	UnicodeTest()
}

func TestChangeString(t *testing.T) {
	ChangeString()
}

func TestTraverseString(t *testing.T) {
	TraverseString()
}

func TestPointer(t *testing.T) {
	PointerTest()
}

func TestPointerArithmetic(t *testing.T) {
	PointerArithmetic()
}

func TestUnnameStruct(t *testing.T) {
	UnnamedStruct()
}

func TestStructTest(t *testing.T) {
	StructEx()
}
