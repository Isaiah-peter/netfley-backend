package main

import (
	"log"
	"net/http"

	"github.com/Isaiah-peter/netfley-backend/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterNetflixUser(r)
	routes.CrudUser(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("Localhost:8000", r))
}
