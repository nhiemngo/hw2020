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
	//mux := http.NewServeMux()

	app := application{
		&sellerDB{
			DB: db,
		},
	}

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("templates/css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("templates/js"))))
	http.HandleFunc("/", functionHandler)
	http.HandleFunc("/view/", app.viewHandler)
	http.HandleFunc("/testseller/", viewTestSellerHandler)
	http.HandleFunc("/create/", app.createHandler)
	http.HandleFunc("/login/", login)
	http.HandleFunc("/logout/", logout)
	http.HandleFunc("/secret/", secret)

	http.ListenAndServe(":2000", nil)
}
