package handlers

// import (
// 	"crudGo/pkg/mocks"
// 	"encoding/json"
// 	"net/http"
// 	"strconv"

// 	"github.com/gorilla/mux"
// )

// func GetBook(w http.ResponseWriter, r *http.Request) {
// 	// Read dynamic id parameter
// 	vars := mux.Vars(r)
// 	id, _ := strconv.Atoi(vars["id"])// as here we want id to be int not string. and it returns 

// 	// Iterate over all the mock books
// 	for _, book := range mocks.Books {
// 		// if book.Id == id {
// 		// 	// If ids are equal send book as response
// 		// 	w.Header().Add("Content-Type", "application/json")
// 		// 	w.WriteHeader(http.StatusOK)
// 		// 	json.NewEncoder(w).Encode(book)
// 		// 	break
// 		// }
// 	}
// }