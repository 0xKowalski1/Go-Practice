package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type User struct {
	ID   int
	Name string
}

type Users []User

var users = Users{
	User{ID: 1, Name: "user1"},
	User{ID: 2, Name: "user2"},
	User{ID: 3, Name: "user3"},
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	jsonUsers, err := json.Marshal(users)
	if err != nil {
		http.Error(w, "Something went wrong!", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonUsers)
}

func HandleShow(w http.ResponseWriter, r *http.Request, userIDString string) {
	var requestedUser *User
	userID, err := strconv.Atoi(userIDString)
	if err != nil {
		http.Error(w, "Bad userId given.", http.StatusBadRequest)
		return
	}

	for _, user := range users {
		if user.ID == userID {
			requestedUser = &user
			break
		}
	}

	if requestedUser == nil {
		http.Error(w, "No user found.", http.StatusNotFound)
		return

	}

	jsonUser, err := json.Marshal(requestedUser)

	if err != nil {
		http.Error(w, "Something went wrong!", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonUser)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// net/http does not handle params
		if r.URL.Path != "/" {
			userID := strings.Split(r.URL.Path, "/")[1] // The URL pattern is /:id
			HandleShow(w, r, userID)
		} else {
			HandleIndex(w, r)
		}
	})

	log.Println("Listening on 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
