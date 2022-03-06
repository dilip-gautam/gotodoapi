package router

import (
	"go-server/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/task", middleware.GetAllTask).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/task", middleware.CreateTask).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/deleteTask/{id}", middleware.DeleteTask).Methods("DELETE", "OPTIONS")
	return router
}
