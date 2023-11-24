package main

import (
	"log"
	"net/http"

	"Guruprasad/pkg/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()                       // use the gorilla mux to create the server and listen at the port 8080
	routes.RegisterBookStoreRoutes(r)          // we have the method RegisteBookStoreRoutes in the route folder so in which it require the router so we pass the 'r' which is consist of router
	http.Handle("/", r)                        // handle the root and return the router
	log.Fatal(http.ListenAndServe(":8080", r)) // listen at the port 8080
}


