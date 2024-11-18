package main

import (
	"log"
	"net/http"

	"github.com/Chinzzii/CRUD-API-with-MySQL-and-Go/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to the Bookstore API! Use /book/ to access the API endpoints."))
	})

	routes.RegisterBookStoreRoutes(r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))

}
