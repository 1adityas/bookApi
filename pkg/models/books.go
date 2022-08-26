package models

import "go.mongodb.org/mongo-driver/bson/primitive"
type Book struct {
	Id     primitive.ObjectID   `json:"id,omitempty" bson:"id,omitempty"`
	Title  string `json:"title,omitempty"`
	Author string `json:"author,omitempty"`
	Desc   string `json:"desc,omitempty"`
}

