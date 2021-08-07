package routes

import (
	"net/http"

	"github.com/Isaiah-peter/netfley-backend/pkg/controllers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var RegisterNetflixUser = func(router *mux.Router) {
	router.HandleFunc("/register/", controllers.NewUser).Methods("POST")
	router.HandleFunc("/login", controllers.Login).Methods("POST")

	handler := cors.New(cors.Options{
		AllowedMethods: []string{"POST"},
	}).Handler(router)

	http.ListenAndServe("Localhost:8000", handler)
}
