package main

import (
	"flag"
	"net/http"
)

var (
	port = flag.String("port", ":8080", "port to listen to")
)

func main() {
	flag.Parse()
	http.HandleFunc("/", static)
	println("Listen on port " + *port)
	http.ListenAndServe(*port, nil)
}

func static(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./"+r.URL.Path)
}
