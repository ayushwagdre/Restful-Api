package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tutorials/go/crud/pkg/models"
)

func (h handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Read request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updatedBook *models.Book
	json.Unmarshal(body, &updatedBook)

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	// book.Title = updatedBook.Title
	// book.Author = updatedBook.Author
	// book.Desc = updatedBook.Desc
	if book.Config == nil {
		book.Config = updatedBook.Config
	}
	for key, value := range updatedBook.Config {
		//loop over value
		log.Println(key, value)
		getEachElement := value.(map[string]interface{})
		if getEachElement["priceXPath"] == nil || getEachElement["priceXPath"] == "" {
			log.Println("priceXPath is empty")
		}
		if getEachElement["renderToPath"] == nil || getEachElement["renderToPath"] == "" {
			log.Println("render is empty")
		}
		log.Println(reflect.TypeOf(getEachElement["style"]))
		if getEachElement["style"] == nil {

			log.Println("style is empty")
		}
	}

	h.DB.Save(&book)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Updated")
}
