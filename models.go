package main

type Book struct {	
	ID 		string 	`json:"id"`
	Title 	string	`json:"title"`
	Author 	string	`json:"author"`
}

var books = []*Book {
	{
		ID: "1",
		Title: "Atomic Habits",
		Author: "William",
	},
}

func listBooks() []*Book {
	return books
}