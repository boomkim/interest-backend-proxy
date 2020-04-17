package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("Default: ", r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("param: ", r.Form["test_param"])
	// parameter 전체 출력
	for k, v := range r.Form{
		fmt.Println("key: ", k)
		fmt.Println("val: ", strings.Join(v,""))
	}
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	http.HandleFunc("/", defaultHandler)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("err:", err)
	} else {
		fmt.Println("server started :9090")
	}
}