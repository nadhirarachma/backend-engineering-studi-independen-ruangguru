package main

import (
	"net/http"
)

func StudentHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to Student page"))
	}
}

func RequestMethodGet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(405)
			w.Write([]byte("Method is not allowed"))
			return
		}

		next.ServeHTTP(w, r)
	}) 
}

func main() {
	mux := http.DefaultServeMux
    mux.HandleFunc("/student", StudentHandler())

    var handler http.Handler = mux
    handler = RequestMethodGet(handler)

	server := new(http.Server)
    server.Handler = handler

	http.ListenAndServe("localhost:8080", nil)
}
