package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zhaslanbeksultan/GoProjects/tsis1/pkg/handlers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/barca", handlers.WelcomeHandler)
	r.HandleFunc("/health", handlers.HealthCheck).Methods("GET")
	r.HandleFunc("/barca/all", handlers.PrintAllPlayers).Methods("GET")
	r.HandleFunc("/barca/{name}", handlers.GetPlayerByName).Methods("GET")
	http.ListenAndServe(":5500", r)
}
