package handlers

// import (
// 	"crudGo/pkg/mocks"
// 	"encoding/json"
// 	"net/http"
// 	"strconv"

// 	"github.com/gorilla/mux"
// )

// func DeleteBook(w http.ResponseWriter,r *http.Request) {
// 	vars := mux.Vars(r)
// 	id, _ := strconv.Atoi(vars["id"])

// 	// Iterate over all the mock Books
// 	for index, book := range mocks.Books {
// 		if book.Id == id {
// 			// Delete book and send response if the book Id matches dynamic Id
// 			mocks.Books = append(mocks.Books[:index], mocks.Books[index+1:]...) //syntax from doc

// 			w.Header().Add("Content-Type", "application/json")
// 			w.WriteHeader(http.StatusOK)
// 			json.NewEncoder(w).Encode("Deleted")
// 			break
// 		}
// 	}


// }