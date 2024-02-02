package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zhaslanbeksultan/GoProjects/tsis1/pkg/handlers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/barca", handlers.WelcomeHandler)
	http.ListenAndServe(":5500", r)
}
