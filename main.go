package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type application struct {
	db *sellerDB
}

func main() {
	db, err := sql.Open("mysql", "user:password@/sellerdb")
	if err != nil {
		log.Fatalln(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	app := application{
		&sellerDB{
			DB: db,
		},
	}

	rt := mux.NewRouter()

	rt.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("templates/css"))))
	rt.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("templates/js"))))

	rt.HandleFunc("/", functionHandler)

	rt.HandleFunc("/create", app.createHandler)
	rt.HandleFunc("/create/", app.createHandler)

	rt.HandleFunc("/view/{id}", app.viewHandler)
	rt.HandleFunc("/view/{id}/", app.viewHandler)

	rt.HandleFunc("/option/{id}", optionHandler)
	rt.HandleFunc("/option/{id}/", optionHandler)

	rt.HandleFunc("/order/{id}", orderHandler)
	rt.HandleFunc("/order/{id}/", orderHandler)

	rt.HandleFunc("/testseller/", viewTestSellerHandler)

	http.ListenAndServe(":2000", rt)
}
