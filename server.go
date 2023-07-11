package main

import (
	"fmt"
	"log"
	"net/http"
)

func home (w http.ResponseWriter, r *http.Request) { // those two arguments required
	fmt.Println("Home!")
}

func main() {
	http.HandleFunc("/", home)
	fmt.Println("Server is running on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil)) // if something goes wrong, handle
}