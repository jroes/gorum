package main

import (
	"fmt"
	"net/http"
	"github.com/jroes/gorum/models"
)

func handler(w http.ResponseWriter, r *http.Request) {
	u := models.User{"jon@example.com", "password1!"}
	fmt.Fprintf(w, "Hello world, %v!", u)
}

func main() {
	fmt.Println("Starting server...")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}