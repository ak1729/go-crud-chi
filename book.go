package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type BookHandler struct {
}

func (b BookHandler) ListBooks(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(listBooks())

	if err!=nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}
}


func (b BookHandler) GetBooks(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	book := getBook(id)

	if book == nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	err := json.NewEncoder(w).Encode(book)

	if err!=nil{
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}
}

func (b BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book Book

	err := json.NewDecoder(r.Body).Decode(&book)

	if err!=nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	storeBook(book)

	err1 := json.NewEncoder(w).Encode(book)

	if err1 !=nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}
}

func (b BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var book Book

	err := json.NewDecoder(r.Body).Decode(&book)

	if err!=nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	updatedBook := updateBook(id, book)

	if updatedBook == nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	err1 := json.NewEncoder(w).Encode(updatedBook)

	if err1!=nil{
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

func (b BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	book := deleteBook(id)

	if book==nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func getBook(id string) *Book {
	for _, book := range books {
		if book.ID==id {
			return book
		}
	}
	return nil
}

func storeBook(book Book)  {
	books = append(books, &book)
}

func deleteBook(id string) *Book {
	for i, book := range books {
		if book.ID==id {
			books = append(books[:i], (books)[i+1:]...)
			return &Book{}
		}
	}
	return nil
}

func updateBook(id string, bookUpdate Book) *Book {
	for i, book := range books {
		if book.ID==id {
			books[i]=&bookUpdate
			return book
		}
	}
	return nil
}