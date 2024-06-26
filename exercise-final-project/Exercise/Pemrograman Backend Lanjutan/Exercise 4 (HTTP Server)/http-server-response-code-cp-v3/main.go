package main

import (
	"net/http"
)

var students = []string{
	"Aditira",
	"Dito",
	"Afis",
	"Eddy",
}

func IsNameExists(name string) bool {
	for _, n := range students {
		if n == name {
			return true
		}
	}

	return false
}

func CheckStudentName() http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(405) 
			w.Write([]byte("Method is not allowed"))
		} else {
			name := r.URL.Query().Get("name")

			if IsNameExists(name) {
				w.WriteHeader(200) 
				w.Write([]byte("Name is exists"))
			} else {
				w.WriteHeader(404) 
				w.Write([]byte("Data not found"))
			}
		}
	}
}

func GetMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/students", CheckStudentName())
	return mux
}

func main() {
	http.ListenAndServe("localhost:8080", GetMux())
}
