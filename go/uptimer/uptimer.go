package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Println(req.UserAgent())
		fmt.Println(strings.ToLower(req.Method))
		fmt.Println(req.Proto)

		res.WriteHeader(404)
		res.Write([]byte("hello"))
	})

	http.ListenAndServe(":8080", nil)
}  
 