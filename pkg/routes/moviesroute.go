package routes

import (
	"github.com/Isaiah-peter/netfley-backend/pkg/controllers"
	"github.com/gorilla/mux"
)

var Movies = func(router *mux.Router) {
	router.HandleFunc("/movie", controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/random", controllers.GetRandomMovie).Methods("GET")
	router.HandleFunc("/movie/{id}", controllers.GetMovieById).Methods("GET")
	router.HandleFunc("/movie/{id}", controllers.UpdateMovie).Methods("PUT")
	router.HandleFunc("/movie/{id}", controllers.DeleteMovie).Methods("DELETE")
}
