package routes

import (
	controllers "Guruprasad/Controllers"

	"github.com/gorilla/mux"
)

var CalorieTrackingRoutes = func(router *mux.Router) {

	router.HandleFunc("/entry/create", controllers.CreateEntry).Methods("POST")
	router.HandleFunc("/entries", controllers.GetEntries).Methods("GET")
	router.HandleFunc("/entry/{id}", controllers.GetEntryById).Methods("GET")
	router.HandleFunc("/ingredient/{ingredients}", controllers.GetEntryByIngredient).Methods("GET")

	router.HandleFunc("/entry/update/{id}", controllers.UpdateEntry).Methods("PUT")
	router.HandleFunc("/ingredient/update/{id}", controllers.UpdateIngredients).Methods("PUT")
	router.HandleFunc("/entry/delete/{id}", controllers.DeleteEntry).Methods("DELETE")
	//router.HandleFunc("/entry/deleteall}", controllers.DeleteAll).Methods("DELETE")

}
