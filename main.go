package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("INDEX PAGE"))
	})

	mux.HandleFunc("GET /test/{name}/{$}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello " + r.PathValue("name")))
	})

	mux.HandleFunc("GET /test/{name}/{another}/{$}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello " + r.PathValue("name") + r.PathValue("another")))
	})

	fmt.Println("Listenning on http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("err:", err)
		return
	}
}
