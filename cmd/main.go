package main

import (
	"crudGo/pkg/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//well all package names need to be same in same folder and all files are patched in same file
// and we can use same file variables etc without importing .
func main(){
	router := mux.NewRouter()

	router.HandleFunc("/books", handlers.GetAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", handlers.GetBook).Methods(http.MethodGet)
	router.HandleFunc("/books", handlers.AddBook).Methods(http.MethodPost)
	router.HandleFunc("/books/{id}", handlers.UpdateBook).Methods(http.MethodPut)
	router.HandleFunc("/books/{id}", handlers.DeleteBook).Methods(http.MethodDelete)

	log.Println("API is running!")
	http.ListenAndServe(":4000", router)
}