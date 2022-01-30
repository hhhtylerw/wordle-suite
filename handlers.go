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
	fmt.Println(username, password)
	if username == "" || password == "" || len(username) > 20 || len(password) > 20 {
		respondWithJson(w, http.StatusBadRequest, map[string]string{"message": "Username and password required"})
		return
	}

	// Verify username is available
	if !usernameAvailable(username) {
		respondWithJson(w, http.StatusBadRequest, map[string]string{
			"message": "Username taken",
		})
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
	userStruct := getAccount(username)
	respondWithJson(w, http.StatusOK, userStruct)

	/*respondWithJson(w, http.StatusOK, map[string]interface{}{
		"message": "Successfully created account",
	})*/
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

func Test(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, http.StatusOK, map[string]string{"message": "Hello, World!"})
}
