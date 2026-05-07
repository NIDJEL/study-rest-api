package main

import (
	"fmt"
	"net/http"
	"strconv"
)

var counter int

func GetHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		fmt.Fprintln(w, "Counter равен", strconv.Itoa(counter))
	} else {
		fmt.Fprintln(w, "Поддерживается только метод GET")
	}
}

func PostHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		counter++
		fmt.Fprintln(w, "Counter увеличен на 1")
	} else {
		fmt.Fprintln(w, "Поддерживается только метод POST")
	}
}

func main() {
	http.HandleFunc("/get", GetHandler)
	http.HandleFunc("/post", PostHandler)

	http.ListenAndServe("localhost:8080", nil)
}
