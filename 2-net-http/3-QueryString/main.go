package main

import "net/http"

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/query", queryFunc)

	http.ListenAndServe(":8080", mux)
}

func queryFunc(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	surname := r.FormValue("surname")

	data := "/query?name=" + name + "&surname=" + surname

	w.Write([]byte(data))
}
