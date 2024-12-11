package main

import (
	"fmt"
	"math"
	"time"
)

func sum(id int) {
	var x int64
	for i := 0; i < math.MaxUint32; i++ {
		x += int64(i)
	}
	fmt.Println(id, x)
}

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

var now = time.Now()

// func init() {
// 	fmt.Println("init now:", now)
// }

func init() {
	fmt.Println("init:", int(time.Now().Sub(now).Seconds()))
	w := make(chan bool)

	go func() {
		time.Sleep(time.Second * 3)
		w <- true
	}()
	<-w
}

func main() {
	fmt.Println("main:", int(time.Now().Sub(now).Seconds()))
}

// func main() {
// 	http.HandleFunc("/qyuhen/test", handler)
// 	http.ListenAndServe(":80", nil)
// }

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, `<meta name="go-import" content="test.com/qyuhen/test git https://github.com/qyuhen/test">`)
// }
