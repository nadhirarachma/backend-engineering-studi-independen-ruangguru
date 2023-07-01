package main

import (
	"net/http"
	"time"
	"strconv"
)

func TimeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		currentTime := time.Now()
		timeFormat := currentTime.Weekday().String() + ", " + strconv.Itoa(currentTime.Day()) + " " + currentTime.Month().String() + " " + strconv.Itoa(currentTime.Year())
		w.Write([]byte(timeFormat))
	} 
}

func SayHelloHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")

		if name == "" {
			w.Write([]byte("Hello there"))
		} else {
			w.Write([]byte("Hello, " + name + "!"))
		}
	} 
}

func GetMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/time", TimeHandler())
	mux.HandleFunc("/hello", SayHelloHandler())
	return mux
}

func main() {
	http.ListenAndServe("localhost:8080", GetMux())
}
