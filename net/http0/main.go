package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello, go world")
	})
	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatal(err)
	}
}
