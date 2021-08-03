package routes

import (
	"github.com/Isaiah-peter/netfley-backend/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterNetflixUser = func(router *mux.Router) {
	router.HandleFunc("/register/", controllers.NewUser).Methods("POST")
	router.HandleFunc("/login", controllers.Login).Methods("POST")
}
