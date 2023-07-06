package main

import (
	"log"
	"net/http"
	"time"
)

// 接口超时控制
func home0(w http.ResponseWriter, r *http.Request) {
	var resp string

	done := make(chan struct{}, 1)
	go func() {
		resp = readDB()
		done <- struct{}{}
	}()

	select {
	case <-done:
	case <-time.After(100 * time.Millisecond):
		resp = "timeout"
	}
	// fmt.Fprint(w, []byte(resp))
	w.Write([]byte(resp))
}

func readDB() string {
	time.Sleep(200 * time.Millisecond)
	return "ok"
}

func baseHTTP() {
	http.HandleFunc("/", home)
	http.HandleFunc("/home0", home0)
	http.HandleFunc("/body/once", readBodyOnce)
	http.HandleFunc("/body/multi", getBodyIsNil)
	http.HandleFunc("/url/query", queryParams)
	http.HandleFunc("/header", header)
	http.HandleFunc("/wholeUrl", wholeURL)
	http.HandleFunc("/form", form)
	//if err := http.ListenAndServe(":8080", nil); err != nil {
	//	panic(err)
	//}
	// http.ListenAndServe("127.0.0.1:8085", nil)
	log.Fatal(http.ListenAndServe(":8085", nil))
}

func main() {
	baseHTTP()
}

// func main() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		_, err := fmt.Fprintln(w, "hello, go world")
// 		if err != nil {
// 			return
// 		}
// 	})
// 	// 如果启动 一个goroutine去监听端口，那么主goroutine退出，其它goroutine也退出
// 	// 需要select{} 阻塞主goroutine退出
// 	go func() {
// 		if err := http.ListenAndServe(":8888", nil); err != nil {
// 			log.Fatal(err)
// 		}
// 	}()

// 	select {}
// }
