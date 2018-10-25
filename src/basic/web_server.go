package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	var h Hello
	err := http.ListenAndServe("localhost:4000", h)
	if err != nil {
		log.Fatal(err)
	}
}

type Hello struct{}

/**
ServeHTTP is a Handler function

type Handler interface {
	ServeHTTP(w ResponseWriter, r *Request)
}

*/
func (h Hello) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	log.Println(r.Header)
	fmt.Fprint(w, "hello")
}
