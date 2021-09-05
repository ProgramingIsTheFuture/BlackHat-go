package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	listenAddr string
	wsAddr     string
	jsTemplate *template.Template
)

func init() {

	flag.StringVar(&listenAddr, "listen-addr", "", "Address to listen on")
	flag.StringVar(&wsAddr, "ws-addr", "", "Address for websocket connection")

	flag.Parse()

	var err error

	jsTemplate, err = template.ParseFiles("logger.js")
	if err != nil {
		panic(err)
	}

}

func serverWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "", 500)
		return
	}

	defer conn.Close()
	fmt.Printf("Connection from %s\n", conn.RemoteAddr().String())

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}

		fmt.Printf("From %s: %s\n", conn.RemoteAddr().String(), string(msg))
	}
}

func serveFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/javascript")
	jsTemplate.Execute(w, wsAddr)
}

func htmlFile(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("index.html")
	if err != nil {
		return
	}

	t.Execute(w, nil)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ws", serverWS)
	r.HandleFunc("/k.js", serveFile).Methods("GET", "OPTIONS")
	r.HandleFunc("/", htmlFile)
	log.Fatal(http.ListenAndServe(listenAddr, r))
}
