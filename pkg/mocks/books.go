package mocks

import "crudGo/pkg/models" // this is right of importng.

var Books = []models.Book{
	{
		// Id:     1,
		Title:  "Golang",
		Author: "Gopher",
		Desc:   "A book for Go",
	},
}