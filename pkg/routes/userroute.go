package routes

import (
	"net/http"

	"github.com/Isaiah-peter/netfley-backend/pkg/controllers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var CrudUser = func(router *mux.Router) {
	router.HandleFunc("/user", controllers.GetUser).Methods("GET")
	router.HandleFunc("/user/{id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{id}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/user/{id}", controllers.DeleteUser).Methods("DELETE")

	handler := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "DELETE", "PUT"},
	}).Handler(router)

	http.ListenAndServe("Localhost:8000", handler)

}
