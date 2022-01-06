package main

import (
	"net/http"
)

func profileHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Profile Page"))
	w.WriteHeader(http.StatusOK)

	path := r.URL.Path
	w.Write([]byte(path))
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Net Http!"))
	})

	http.HandleFunc("/profile", profileHandler)

	http.ListenAndServe(":8080", nil)
}
