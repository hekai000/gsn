package main

/*
	#include <stdlib.h>
	#include <stdio.h>

	void test(char* s) {
		printf("%s\n", s);
	}
	char* cstr(){
		return "abcde";
	}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	s := "Hello, World!"
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	C.test(cs)
	cs = C.cstr()
	fmt.Println(C.GoString(cs))
	fmt.Println(C.GoStringN(cs, 2))
	fmt.Println(C.GoBytes(unsafe.Pointer(cs), 2))
}

// func sum(id int) {
// 	var x int64
// 	for i := 0; i < math.MaxUint32; i++ {
// 		x += int64(i)
// 	}
// 	fmt.Println(id, x)
// }

// func main() {
// 	wg := new(sync.WaitGroup)
// 	wg.Add(2)

// 	for i := 0; i < 2; i++ {
// 		go func(id int) {
// 			defer wg.Done()
// 			sum(id)
// 		}(i)
// 	}
// 	wg.Wait()
// }

// var now = time.Now()

// func init() {
// 	fmt.Println("init now:", now)
// }

// func init() {
// 	fmt.Println("init:", int(time.Now().Sub(now).Seconds()))
// 	w := make(chan bool)

// 	go func() {
// 		time.Sleep(time.Second * 3)
// 		w <- true
// 	}()
// 	<-w
// }

// func main() {
// 	fmt.Println("main:", int(time.Now().Sub(now).Seconds()))
// }

// func main() {
// 	http.HandleFunc("/qyuhen/test", handler)
// 	http.ListenAndServe(":80", nil)
// }

//	func handler(w http.ResponseWriter, r *http.Request) {
//		fmt.Fprint(w, `<meta name="go-import" content="test.com/qyuhen/test git https://github.com/qyuhen/test">`)
//	}
// type data struct {
// 	x [1024 * 100]byte
// }

// func test() uintptr {
// 	p := &data{}
// 	return uintptr(unsafe.Pointer(p))
// }

// func main() {
// 	const N = 10000
// 	cache := new([N]uintptr)
// 	for i := 0; i < N; i++ {
// 		cache[i] = test()
// 		time.Sleep(time.Microsecond)
// 	}
// }
