package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /test/first/second/third/fourth/fifth", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("INDEX PAGE"))
	})

	mux.HandleFunc("GET /test/{first}/{$}", func(w http.ResponseWriter, r *http.Request) {
		first := r.PathValue("first")
		w.Write([]byte(fmt.Sprintf("Hello %s", first)))
	})

	mux.HandleFunc("GET /test/{first}/{second}/{$}", func(w http.ResponseWriter, r *http.Request) {
		first := r.PathValue("first")
		second := r.PathValue("second")
		w.Write([]byte(fmt.Sprintf("Hello %s %s", first, second)))
	})

	mux.HandleFunc("GET /test/{first}/{second}/{third}/{fourth}/{fifth}/{$}", func(w http.ResponseWriter, r *http.Request) {
		first := r.PathValue("first")
		second := r.PathValue("second")
		third := r.PathValue("third")
		fourth := r.PathValue("fourth")
		fifth := r.PathValue("fifth")
		w.Write([]byte(fmt.Sprintf("Hello %s %s %s %s %s", first, second, third, fourth, fifth)))
	})

	fmt.Println("Listenning on http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("err:", err)
		return
	}
}
