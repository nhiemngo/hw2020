package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", functionHandler)
	mux.HandleFunc("/view/", viewHandler)
	mux.HandleFunc("/testseller/", viewTestSellerHandler)
	mux.HandleFunc("/create/", createHandler)
	mux.HandleFunc("/login/", login)
	mux.HandleFunc("/logout/", logout)
	mux.HandleFunc("/secret/", secret)

	http.ListenAndServe(":2000", mux)
}
