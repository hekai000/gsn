package chap2

import "fmt"

// 保留字：package, if, else, break, continue, for, switch, case, func, go, chan, const, type, var, return,
//import, interface, map, struct, defer, select, const, goto, map, range, fallthrough, default

//位运算

// & | ^
//&^ a&^b= a&b^a=0100&^1011= 0100&1011^0100=1011,按照右侧数清楚标志位，右侧数为1清楚，为0则保留左侧数

//标志位操作
// a:=0
// a|=1<<2
// a|=1<<6
// a=a&^(1<<6)

// p++,p--是语句，不是表达式
// var aa = []int{
// 	1,
// 	2,
// 	3,
// }

// 不支持三元操作 a > b? a : b不支持

func RangeCopy() {
	a := [3]int{1, 2, 3}

	for i, v := range a {
		if i == 0 {

			a[1], a[2] = 999, 999
			fmt.Println(a)
		}
		a[i] = v + 100
	}
	fmt.Println(a)
}

func RangeTest() {
	s := []int{1, 2, 3, 4, 5}
	for i, v := range s {
		if i == 0 {
			s = s[:3]
			s[2] = 100
			fmt.Println(s)
		}
		fmt.Println(i, v)
	}
	fmt.Println(s)
}

func SwitchTest() {
	x := []int{1, 2, 3}
	i := 2
	switch i {
	case x[i]:
		fmt.Println("a")
	case 1, 3:
		fmt.Println("b")
	default:
		fmt.Println("c")
	}
}

func GotoTest() {
	var i int
	for {
		fmt.Println(i)
		i++
		if i > 2 {
			goto BREAK
		}
	}
BREAK:
	fmt.Println("break")
}
