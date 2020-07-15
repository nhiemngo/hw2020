package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
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
	mux := http.NewServeMux()

	app := application{
		&sellerDB{
			DB: db,
		},
	}
	mux.HandleFunc("/", functionHandler)
	fs := http.FileServer(http.Dir("templates"))
	http.Handle("/css/", fs)
	mux.HandleFunc("/view/", app.viewHandler)
	mux.HandleFunc("/testseller/", viewTestSellerHandler)
	mux.HandleFunc("/create/", app.createHandler)
	mux.HandleFunc("/login/", login)
	mux.HandleFunc("/logout/", logout)
	mux.HandleFunc("/secret/", secret)

	http.ListenAndServe(":2000", mux)
}
