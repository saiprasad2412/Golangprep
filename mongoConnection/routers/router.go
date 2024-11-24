package router

import (
	controller "main/controllers"

	"github.com/gorilla/mux"
)

func Router()*mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies",controller.GetAllMoviesController).Methods("GET")
	router.HandleFunc("/api/movies/{id}",controller.MarkAsWatchedController).Methods("PUT")

	return router
}