package models
type Book struct {
	Id     int    `json:"id" bson:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"	desc"`
}
