package handlers

import (
	"bytes"
	"crudGo/pkg/mocks"
	"crudGo/pkg/models"
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net/http"
)

func AddBook(w http.ResponseWriter, r *http.Request){
	defer r.Body.Close()// its called when we want any function to run in the end and if more than more defer are used then LIFO;
	var body bytes.Buffer
	_, err := io.Copy(&body,r.Body)//dont use ioutil.ReadAll .. as you can google again lol

	/*Response of the body can be read using any method that could read data from incoming byte stream.
	 Simplest of them is ReadAll function provided in ioutil package.*/

	if err != nil {
		log.Fatalln(err)
	}

	var book models.Book
	json.Unmarshal(body.Bytes(), &book)

	// Append to the Book mocks
	book.Id = rand.Intn(100) 
	mocks.Books = append(mocks.Books, book)

	// Send a 201 created response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
	
	//Marshal and Unmarshal convert a string into JSON and vice versa.
	// Encoding and decoding convert a stream into JSON and vice versa.
	/*Encoder and decoder write struct to slice of a stream or read data from
	a slice of a stream and convert it into a struct. Internally, it also
	implements the marshal method. The only difference is if you want to play with string or bytes use marshal,
	and if any data you want to read or write to some writer interface, use encodes and decode.*/
}