package controller

import (
	"context"
	"fmt"
	"log"

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

