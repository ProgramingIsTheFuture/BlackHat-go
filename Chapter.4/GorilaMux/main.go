package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/foo", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello mux world!")
	}).Methods("GET")

	r.HandleFunc("/users/{user}", func(w http.ResponseWriter, req *http.Request) {
		user := mux.Vars(req)["user"]
		fmt.Fprintf(w, "Hello %s", user)
	})

	http.ListenAndServe(":8000", r)
}
