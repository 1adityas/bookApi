package controller

import (
	"context"
	"crudGo/pkg/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectingStr = "mongodb+srv://7aditya7:7aditya7@cluster0.nftzibi.mongodb.net/?retryWrites=true&w=majority"
const dbName = "bookApi"
const colName = "library"

// most imp
var collection *mongo.Collection

// connect with mongo now . 

func init(){
	// it runs for first time and one one time 
clientOptions:= options.Client().ApplyURI(connectingStr)

//connect to mongo 	
client, err:= mongo.Connect(context.TODO(),clientOptions)

if err!=nil{
	log.Fatal(err)
}
fmt.Println("connection successfull")

collection=(*mongo.Collection)(client.Database(dbName).Collection(colName))
//collection instance is ready 
fmt.Printf("collection instance is ready ")

}

func insertBook(book models.Book){
	inserted , err:=collection.InsertOne(context.Background(),book)

	if err!=nil{fmt.Print(err)}
	fmt.Println(inserted.InsertedID) 
}
func getAllBooks()[] primitive.M{
	cur,err:=collection.Find(context.Background(),bson.D{{}})
	//also do check for errror

	var books [] primitive.M 
	
	for cur.Next(context.Background()){
		var book bson.M
		err=cur.Decode(&book)
		if err!=nil{
			log.Fatal(err)
		}
		books=append(books,book)
	}
	defer cur.Close(context.Background())
	return books

}
func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(mocks.Books)
	allBooks:=getAllBooks()
	json.NewEncoder(w).Encode(allBooks)
}
func AddBook(w http.ResponseWriter, r *http.Request){
	w.Header().Add("Content-Type", "application/json")
	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)
	insertBook(book)
	json.NewEncoder(w).Encode(book)
}