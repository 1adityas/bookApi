package handlers

// import (
// 	"crudGo/pkg/mocks"
// 	"crudGo/pkg/models"
// 	"encoding/json"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"strconv"

// 	"github.com/gorilla/mux"
// )

// func UpdateBook(w http.ResponseWriter, r *http.Request) {
// 	// Read dynamic id parameter
// 	vars := mux.Vars(r)
// 	id, _ := strconv.Atoi(vars["id"])

// 	// Read request body
// 	defer r.Body.Close()
// 	body, err := ioutil.ReadAll(r.Body)//i used its alternate at add books

// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	var updatedBook models.Book
// 	json.Unmarshal(body, &updatedBook)

// 	// Iterate over all the mock Books
// 	for index, book := range mocks.Books {
// 		if book.Id == id {
// 			// Update and send response when book Id matches dynamic Id
// 			book.Title = updatedBook.Title
// 			book.Author = updatedBook.Author
// 			book.Desc = updatedBook.Desc

// 			mocks.Books[index] = book

// 			w.Header().Add("Content-Type", "application/json")
// 			w.WriteHeader(http.StatusOK)
// 			json.NewEncoder(w).Encode("Updated")
// 			break
// 		}
// 	}
// }