package main

import (
	"fmt"
	"github.com/jroes/gorum/models"
	"net/http"
)

func signInHandler(w http.ResponseWriter, r *http.Request) {
	var userStore models.UserStore
	userStore = models.NewUserGobStore("users/")
	email := r.FormValue("email")
	password := r.FormValue("password")
	user, err := userStore.FindUser(email)
	if err != nil {
		fmt.Fprintf(w, "Sorry, couldn't find a user with that email address.\n")
		return
	}

	err = user.HasPassword(password)
	if err != nil {
		fmt.Fprintf(w, "Looks like you have the wrong password!\n")
		return
	}
	fmt.Fprintf(w, "Looks like you have the matching password!\n")
}

func signUpHandler(w http.ResponseWriter, r *http.Request) {
	var userStore models.UserStore
	userStore = models.NewUserGobStore("users/")
	email := r.FormValue("email")
	password := r.FormValue("password")
	user := models.NewUser(email, password)
	err := userStore.SaveUser(*user)
	if err != nil {
		fmt.Fprintf(w, "Had trouble creating %s. Error: %v\n", email, err)
		return
	}

	fmt.Fprintf(w, "Your user account has been created!\n")
}

func main() {
	http.HandleFunc("/sign_up", signUpHandler)
	http.HandleFunc("/sign_in", signInHandler)
	fmt.Println("Listening for requests...")
	http.ListenAndServe(":8080", nil)
}
