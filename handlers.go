package main

import (
	"fmt"
	"net/http"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	// Get username and password from request
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")

	// Verify username is available
	if !usernameAvailable(username) {
		w.Write([]byte("Username taken"))
		return
	}

	// Create new user account
	client.Collection("accounts").Doc(username).Set(ctx, map[string]interface{}{
		"username": username,
		"password": getSHA256(password + salt),
		"friends":  []string{"h3x"},
		"plays":    0,
		"wins":     0,
		"losses":   0,
	})
	w.Header().Set("Session", username)
	respondWithJson(w, http.StatusOK, map[string]interface{}{
		"message": "Successfully created account",
	})
}

func usernameAvailable(username string) bool { // Check if username exists in database
	userAccount, err := client.Collection("accounts").Doc(username).Get(ctx)
	if err != nil {
		fmt.Println(err)
		return true
	}
	return !userAccount.Exists()
}

func Login(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello, World!"))
}
