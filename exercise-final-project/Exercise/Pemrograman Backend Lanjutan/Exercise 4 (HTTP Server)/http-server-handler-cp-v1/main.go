package main

import (
	"net/http"
	"time"
	"strconv"
)

func GetHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		currentTime := time.Now()
		timeFormat := currentTime.Weekday().String() + ", " + strconv.Itoa(currentTime.Day()) + " " + currentTime.Month().String() + " " + strconv.Itoa(currentTime.Year())
		writer.Write([]byte(timeFormat))
	} 
}

func main() {
	http.ListenAndServe("localhost:8080", GetHandler())
}
