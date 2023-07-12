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

func todos(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("todos.html")

	if err != nil { // handle error
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Template parsing error:", err)
	}

	err = t.Execute(w, nil) // applies parsed template to specified data obj
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/todos/", todos)
	fmt.Println("Server is running on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil)) // if something goes wrong, handle
}