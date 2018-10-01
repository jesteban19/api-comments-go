package routes

import (
	"github.com/gorilla/mux"
	"github.com/jesteban19/edcomments/controllers"
)

// router para login
func SetLoginRouter(router *mux.Router) {
	router.HandleFunc("/api/login", controllers.Login).Methods("POST")
}
