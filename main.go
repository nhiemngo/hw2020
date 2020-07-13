package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", functionHandler)
	mux.HandleFunc("/view/", viewHandler)
	mux.HandleFunc("/create/", createHandler)

	http.ListenAndServe(":2000", mux)
}
