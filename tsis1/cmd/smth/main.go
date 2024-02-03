package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/barca", WelcomeHandler)
	r.HandleFunc("/health", HealthCheck).Methods("GET")
	r.HandleFunc("/barca/all", PrintAllPlayers).Methods("GET")
	r.HandleFunc("/barca/{name}", GetPlayerByName).Methods("GET")
	http.ListenAndServe(":5500", r)
}
