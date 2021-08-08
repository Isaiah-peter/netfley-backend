package routes

import (
	"github.com/Isaiah-peter/netfley-backend/pkg/controllers"
	"github.com/gorilla/mux"
)

var Movies = func(router *mux.Router) {
	router.HandleFunc("/movie", controllers.CreateMovie).Methods("POST")
}
