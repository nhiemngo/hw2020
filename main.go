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
	// database stuffs
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

	//server stuffs
	rt := mux.NewRouter()

	rt.HandleFunc("/", functionHandler)

	rt.HandleFunc("/view/{id}", app.viewHandler)
	rt.HandleFunc("/view/{id}/", app.viewHandler)

	rt.HandleFunc("/testseller/", viewTestSellerHandler)

	rt.HandleFunc("/create/", app.createHandler)
	rt.HandleFunc("/login/", login)
	rt.HandleFunc("/logout/", logout)
	rt.HandleFunc("/secret/", secret)

	http.ListenAndServe(":2000", rt)
}
