package middleware

import (
	"fmt"
	"net/http"

	handler "github.com/MarceloLima11/food-nearby/server/handlers"
	"github.com/gorilla/mux"
)

func initializeRoutes() {
	r := mux.NewRouter()

	r.Use(initializeCors)

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	routes(r)

	fmt.Print("Server is ready")
	err := server.ListenAndServe()

	if err != nil {
		fmt.Print("Erro ao iniciar servidor: ", err)
	}
}

func routes(router *mux.Router) {
	handler.InitializeHandler()

	router.HandleFunc("/restaurants", handler.ListRestaurants).Methods("GET")
	router.HandleFunc("/restaurant/{id}", handler.ShowRestaurant).Methods("GET")
	router.HandleFunc("/restaurants", handler.CreateRestaurant).Methods("POST")
	router.HandleFunc("/restaurant/{id}", handler.UpdateRestaurant).Methods("UPDATE")
	router.HandleFunc("/restaurant/{id}", handler.DeleteRestaurant).Methods("DELETE")
}
