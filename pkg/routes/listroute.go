package routes

import (
	"github.com/Isaiah-peter/netfley-backend/pkg/controllers"
	"github.com/gorilla/mux"
)

var List = func(router *mux.Router) {
	router.HandleFunc("/list", controllers.CreateList).Methods("POST")
}
