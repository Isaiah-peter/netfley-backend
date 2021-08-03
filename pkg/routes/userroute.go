package routes

import (
	"github.com/Isaiah-peter/netfley-backend/pkg/controllers"
	"github.com/gorilla/mux"
)

var CrudUser = func(router *mux.Router) {
	// router.HandleFunc("/user", controllers.GetUser).Methods("GET")
	router.HandleFunc("/user/{id}", controllers.UpdateUser).Methods("PUT")
}
