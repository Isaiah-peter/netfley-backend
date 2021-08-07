package routes

import (
	"net/http"

	"github.com/Isaiah-peter/netfley-backend/pkg/controllers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var Movies = func(router *mux.Router) {
	router.HandleFunc("/movie", controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/random/", controllers.GetRandomMovie).Methods("GET")
	router.HandleFunc("/series/", controllers.GetMovieByType).Methods("GET")
	router.HandleFunc("/movies/", controllers.GetMovieByTypeMovie).Methods("GET")
	router.HandleFunc("/movie/{id}", controllers.GetMovieById).Methods("GET")
	router.HandleFunc("/movie/{id}", controllers.UpdateMovie).Methods("PUT")
	router.HandleFunc("/movie/{id}", controllers.DeleteMovie).Methods("DELETE")

	handler := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "DELETE", "PUT"},
	}).Handler(router)

	http.ListenAndServe("Localhost:8000", handler)
}
