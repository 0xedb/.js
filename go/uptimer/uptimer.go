package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Println(req.UserAgent())

		res.Write([]byte("hello"))
	})

	http.ListenAndServe(":8080", nil)
}
