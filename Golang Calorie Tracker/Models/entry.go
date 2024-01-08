package models

import (
	config "Guruprasad/Config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Entry struct {
	gorm.Model

	Dish        string  `json:"dish"`
	Fat         float64 `json:"fat"`
	Ingredients string  `json:"ingerdients"`
	Calories    string  `json:"calories"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Entry{})
}

func (e *Entry) CreateEntry() *Entry {
	db.NewRecord(e)
	db.Create(e)
	return e
}

func GetEntry() []Entry {

	var entries []Entry
	db.Find(&entries)
	return entries
}

func GetEntryById(id int64) (*Entry, *gorm.DB) {
	var entry Entry
	db := db.Where("id =?", id).Find(&entry)
	return &entry, db
}

func GetByIngredients(ing string) []Entry {
	var entry []Entry
	db.Where("ingredients =?", ing).Find(&entry)
	return entry
}

func DeleteEntry(id int64) Entry {
	var entry Entry
	db.Where("id =?", id).Delete(entry)
	return entry
}

// func DeleteAll() error {

// 	err := db.Exec("DELETE FROM entries")
// 	if err != nil {
// 		fmt.Println("Erorr while deleting the row")
// 	}

// 	return err.Error

// }
