package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Net Http!"))
	})
	http.ListenAndServe(":8080", nil)
}
