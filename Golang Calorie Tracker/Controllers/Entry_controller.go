package controllers

import (
	models "Guruprasad/Models"
	utils "Guruprasad/Utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var NewEntry models.Entry

func CreateEntry(w http.ResponseWriter, r *http.Request) {
	CreateEntry := &models.Entry{}
	if err := utils.ParseBody(r, CreateEntry); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	b := CreateEntry.CreateEntry()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func GetEntries(w http.ResponseWriter, r *http.Request) {

	entries := models.GetEntry()
	res, _ := json.Marshal(entries)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetEntryById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	entryId := vars["id"]
	id, err := strconv.ParseInt(entryId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	entry, _ := models.GetEntryById(id)
	res, _ := json.Marshal(entry)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func GetEntryByIngredient(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	ingredient := vars["ingredients"]
	entry := models.GetByIngredients(ingredient)
	res, _ := json.Marshal(entry)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateEntry(w http.ResponseWriter, r *http.Request) {

	var updateEntry = &models.Entry{}
	if err := utils.ParseBody(r, updateEntry); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	entryId := vars["id"]
	id, err := strconv.ParseInt(entryId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	entry, db := models.GetEntryById(id)
	if updateEntry.Dish != "" {

		entry.Dish = updateEntry.Dish
	}

	if updateEntry.Fat != 0 {

		entry.Fat = updateEntry.Fat
	}

	if updateEntry.Ingredients != "" {

		entry.Ingredients = updateEntry.Ingredients
	}

	if updateEntry.Calories != "" {

		entry.Calories = updateEntry.Calories
	}

	db.Save(entry)
	res, _ := json.Marshal(entry)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func UpdateIngredients(w http.ResponseWriter, r *http.Request) {

	var updateEntry = &models.Entry{}
	if err := utils.ParseBody(r, updateEntry); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	entryId := vars["id"]
	id, err := strconv.ParseInt(entryId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	entry, db := models.GetEntryById(id)
	if updateEntry.Ingredients != "" {
		entry.Ingredients = updateEntry.Ingredients
	}

	db.Save(entry)

	res, _ := json.Marshal(entry)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteEntry(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	entryId := vars["id"]
	id, err := strconv.ParseInt(entryId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	entry := models.DeleteEntry(id)

	res, err := json.Marshal(entry)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// func DeleteAll(w http.ResponseWriter, r *http.Request) {

// 	entry := models.DeleteAll()
// 	res, err := json.Marshal(entry)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte(err.Error()))
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)

// }
