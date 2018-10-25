package main

import (
	"fmt"
	"net/http"
)

func main() {
	//先注册路径和对应的处理器
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"hello", ":", "Go"})
	//在开启服务器进行监听
	//此时注册服务器时，handler可以为nil
	http.ListenAndServe("localhost:4001", nil)
}

type String string

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func (str String) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, str)
}

func (s *Struct) String() string {
	return fmt.Sprintf("Greet:%v%v, from %v", s.Greeting, s.Punct, s.Who)
}

func (s *Struct) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, s)
}
