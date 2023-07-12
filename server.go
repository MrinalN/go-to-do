package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func home (w http.ResponseWriter, r *http.Request) { // those two arguments required
	fmt.Println("Home!")
}

// some linters need a comment above each type
// Todo is a todo with a title and content
type Todos struct { // typically at top of file
	Title string
	Content string
}

// PageVariables are variables sent to the html template
type PageVariables struct{
	PageTitle string
	PageTodos []Todos
}

var todos []Todos

func getTodos(w http.ResponseWriter, r *http.Request) {
	pageVariables := PageVariables{
		PageTitle: "Get To Dos",
		PageTodos: todos,
	}
	t, err := template.ParseFiles("todos.html")

	if err != nil { // handle error
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Template parsing error:", err)
	}

	// err = t.Execute(w, nil) // applies parsed template to specified data obj. Inital build with nil
	err = t.Execute(w, pageVariables)
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/todos/", getTodos)
	fmt.Println("Server is running on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil)) // if something goes wrong, handle
}