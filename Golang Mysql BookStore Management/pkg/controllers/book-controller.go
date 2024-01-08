package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"Guruprasad/pkg/models"
	"Guruprasad/pkg/utils"

	"github.com/gorilla/mux"
)

var NewBook models.Book // in the model folder there is the struct Book so here we crete the new variable of the Book Struct called NewBooks

func GetBook(w http.ResponseWriter, r *http.Request) { // we create the functions to handle the routes in which it require the response write and request and w , r are the objects
	newBooks := models.GetBook()                       // in the model folder ther is the function call GetBook so this is now stored in the newBooks variable
	res, _ := json.Marshal(newBooks)                   // we decode the data into the json using the marshal
	w.Header().Set("Content-Type", "application/json") // declaring the headers
	w.WriteHeader(http.StatusOK)                       // returning the status
	w.Write(res)                                       // showing the response which is stored in the res variable
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)                       // here the we are storing the database variables into the var means that to get the variable form the database we use this statement and store it to another variable
	bookId := vars["bookId"]                  // so from the database we take the bookId named variable and stored it into the bookId
	ID, err := strconv.ParseInt(bookId, 0, 0) // we are converting the bookId into the string and stored in the ID variable and handleing the error below if any error occur
	if err != nil {
		fmt.Println("Error While Parsing")
	}

	bookDetails, _ := models.GetBookById(ID)           // now there is the function GetBookById in the model folder which require the book and db instance so to store the data we make the variable bookDetails and _ for nothing
	res, _ := json.Marshal(bookDetails)                // in this we decode the data into the json and stored it into the res variable
	w.Header().Set("Content-Type", "application/json") // set the headers
	w.WriteHeader(http.StatusOK)                       // return the status
	w.Write(res)                                       // showing the resposne
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{} // geting the book reference from the model folder
	if err := utils.ParseBody(r, CreateBook); err != nil {
		// Handle the error, e.g., return a bad request response
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	b := CreateBook.CreateBook() // after getting the instance of the Book struct then we call the create book function and stored the funtion inot the 'b'
	res, _ := json.Marshal(b)    // decode the data into the json
	w.WriteHeader(http.StatusOK) // send the status ok
	w.Write(res)                 // show the response
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r) // create the variable vars to store the database variables / instance

	bookId := vars["bookId"] // search the variable bookId in the database and stored in the bookId

	ID, err := strconv.ParseInt(bookId, 0, 0) // convert the bookId into integer
	if err != nil {
		fmt.Println("Error While Parsing") // catch the error while parsing
	}

	book := models.DeleteBook(ID) // calling the function DeleteBook form the model package and give the parameter id ans store the result into the book variable

	res, err := json.Marshal(book) // decode the data and store into the res
	if err != nil {
		fmt.Println("There is the error to marshal") // handle the error
	}
	w.Header().Set("Content-Type", "application/json") // showing the response into the json
	w.WriteHeader(http.StatusOK)                       // ok status
	w.Write(res)                                       // shoiwing the response
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{} // get the reference of the Book from the model package

	utils.ParseBody(r, updateBook) // use the parsebody function form the utils package and give the parameters to it
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {

		fmt.Println("Error while parsing")

	}

	bookDetails, db := models.GetBookById(ID)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}

	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
