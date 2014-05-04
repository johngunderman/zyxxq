package zyxxq

import (
	"net/http"
)

func init() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/resources/", resourceHandler)
}

func resourceHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

func handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "resources/index.html")
}
