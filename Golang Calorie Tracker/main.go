package main

import (
	routes "Guruprasad/Routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	http.Handle("/", r)
	routes.CalorieTrackingRoutes(r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
