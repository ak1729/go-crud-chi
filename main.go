package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func home(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Home endpoint hitted")

	fmt.Fprint(w, "\n Hello World! \n")
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	err := godotenv.Load(".env")

	if err!=nil{
		fmt.Println("env not found")
	}

	r.Get("/", home)

	r.Mount("/books", BookRoutes())

	fmt.Println("Server started at port 3000!")
	http.ListenAndServe(":"+os.Getenv("Port"), r)
}

func BookRoutes() chi.Router {
	r := chi.NewRouter()

	bookHandler := BookHandler{}

	r.Get("/", bookHandler.ListBooks)
	r.Post("/", bookHandler.CreateBook)
	r.Get("/{id}", bookHandler.GetBooks)
    r.Put("/{id}", bookHandler.UpdateBook)
    r.Delete("/{id}", bookHandler.DeleteBook)

    return r
}