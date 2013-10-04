package main

import (
	"fmt"
	"net/http"
	"github.com/jroes/gorum/models"
)

func handler(w http.ResponseWriter, r *http.Request) {
	u := models.NewUser("jon@example.com", "password1!")
	fmt.Fprintf(w, "Hello world, %v!", u)
	err := u.HasPassword(r.URL.Path[1:])
	if err != nil {
		fmt.Fprintf(w, "Looks like you have the wrong password!")
		return
	}
	fmt.Fprintf(w, "Looks like you have the matching password!")
}

func main() {
	fmt.Println("Starting server...")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
