package main

import (
	"fmt"
	"net/http"
)

import (
	"github.com/jroes/gorum/models"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var userStore models.UserStore
	userStore = models.NewUserGobStore("users/")
	user, err := userStore.FindUser("jon@example.com")
	if err != nil {
		fmt.Fprintf(w, "Had trouble finding jon@example.com, attempting creation...\n")
		user = models.NewUser("jon@example.com", "password1!")
		err = userStore.SaveUser(*user)
		if err != nil {
			fmt.Fprintf(w, "Had trouble creating jon@example.com. Ouch: %v\n", err)
			return
		}
	}

	fmt.Fprintf(w, "Hello world, %v!\n", user)
	err = user.HasPassword(r.URL.Path[1:])
	if err != nil {
		fmt.Fprintf(w, "Looks like you have the wrong password!\n")
		return
	}
	fmt.Fprintf(w, "Looks like you have the matching password!\n")
}

func main() {
	fmt.Println("Starting server...")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
