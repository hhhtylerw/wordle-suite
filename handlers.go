package main

import (
	"fmt"
	"net/http"
	"strconv"

	"cloud.google.com/go/firestore"
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

	respondWithJson(w, http.StatusOK, map[string]string{
		"message": "Successfully created account",
		"friends": "h3x,",
		"plays":   "0",
		"wins":    "0",
		"losses":  "0",
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
	// Get username and password from request
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	fmt.Println(username, password)

	// Verify account exists
	if usernameAvailable(username) {
		respondWithJson(w, http.StatusBadRequest, map[string]string{"message": "Account does not exist"})
		return
	}

	// Verify password is correct
	userAccount, _ := client.Collection("accounts").Doc(username).Get(ctx)
	if (getSHA256(password + salt)) != userAccount.Data()["password"].(string) {
		respondWithJson(w, http.StatusBadRequest, map[string]string{"message": "Incorrect password"})
		return
	} else {
		respondWithJson(w, http.StatusOK, map[string]string{
			"message": "Successfully logged in",
			"friends": "friends3x",
			"plays":   strconv.FormatInt(userAccount.Data()["plays"].(int64), 10),
			"wins":    strconv.FormatInt(userAccount.Data()["wins"].(int64), 10),
			"losses":  strconv.FormatInt(userAccount.Data()["losses"].(int64), 10),
		})
	}

}

func Account(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form.Get("username")
	if username == "" {
		respondWithJson(w, http.StatusBadRequest, map[string]string{"message": "Username required"})
		return
	}

	userAccount, _ := client.Collection("accounts").Doc(username).Get(ctx)
	if userAccount.Exists() {
		fmt.Println("User account exists")
	}
}

func Friends(w http.ResponseWriter, r *http.Request) {
	// Get username from request
	r.ParseForm()
	username := r.Form.Get("username")
	if username == "" {
		respondWithJson(w, http.StatusBadRequest, map[string]string{"message": "Username required"})
		return
	}

	if usernameAvailable(username) {
		respondWithJson(w, http.StatusBadRequest, map[string]string{"message": "Account does not exist"})
		return
	}

	userAccount, _ := client.Collection("accounts").Doc(username).Get(ctx)
	friends := userAccount.Data()["friends"].([]interface{})
	fmt.Println(friends)
	friendsMap := make(map[string]map[string]string)

	for i := 0; i < len(friends); i++ {
		friend := friends[i].(string)
		friendAccount, _ := client.Collection("accounts").Doc(friend).Get(ctx)
		friendsMap[friend] = map[string]string{
			"plays":  strconv.FormatInt(friendAccount.Data()["plays"].(int64), 10),
			"wins":   strconv.FormatInt(friendAccount.Data()["wins"].(int64), 10),
			"losses": strconv.FormatInt(friendAccount.Data()["losses"].(int64), 10),
		}
	}

	fmt.Println(friendsMap)
	respondWithJson(w, http.StatusOK, friendsMap)
}

func AddFriend(w http.ResponseWriter, r *http.Request) {
	// Get username and friend from request
	r.ParseForm()
	username := r.Form.Get("username")
	friend := r.Form.Get("friend")
	fmt.Println(username, friend)

	// Verify account exists
	if usernameAvailable(username) {
		respondWithJson(w, http.StatusBadRequest, map[string]string{"message": "Account does not exist"})
		return
	}

	// Verify friend exists
	if usernameAvailable(friend) {
		respondWithJson(w, http.StatusBadRequest, map[string]string{"message": "Friend does not exist"})
		return
	}

	// Add friend to user account
	userAccount, _ := client.Collection("accounts").Doc(username).Get(ctx)
	friends := userAccount.Data()["friends"].([]interface{})
	if contains(friends, friend) { // Return if friend is already in friends list
		fmt.Println("Friend already exists")
		respondWithJson(w, http.StatusBadRequest, map[string]string{"message": "Friend already exists"})
		return
	}

	friends = append(friends, friend)
	client.Collection("accounts").Doc(username).Set(ctx, map[string]interface{}{
		"friends": friends,
	}, firestore.MergeAll)

	respondWithJson(w, http.StatusOK, map[string]string{"message": "Successfully added friend"})
}

func AddWin(w http.ResponseWriter, r *http.Request) {
	// Get username from request
	r.ParseForm()
	username := r.Form.Get("username")
	if username == "" {
		respondWithJson(w, http.StatusBadRequest, map[string]string{"message": "Username required"})
		return
	}

	// Verify account exists
	if usernameAvailable(username) {
		respondWithJson(w, http.StatusBadRequest, map[string]string{"message": "Account does not exist"})
		return
	}

	// Add win to user account
	userAccount, _ := client.Collection("accounts").Doc(username).Get(ctx)
	plays := userAccount.Data()["plays"].(int64)
	wins := userAccount.Data()["wins"].(int64)
	wins++
	plays++
	client.Collection("accounts").Doc(username).Set(ctx, map[string]interface{}{
		"plays": plays,
		"wins":  wins,
	}, firestore.MergeAll)

	respondWithJson(w, http.StatusOK, map[string]string{"message": "Successfully added win"})

}

func AddLoss(w http.ResponseWriter, r *http.Request) {
	// Get username from request
	r.ParseForm()
	username := r.Form.Get("username")
	if username == "" {
		respondWithJson(w, http.StatusBadRequest, map[string]string{"message": "Username required"})
		return
	}

	// Verify account exists
	if usernameAvailable(username) {
		respondWithJson(w, http.StatusBadRequest, map[string]string{"message": "Account does not exist"})
		return
	}

	// Add loss to user account
	userAccount, _ := client.Collection("accounts").Doc(username).Get(ctx)
	plays := userAccount.Data()["plays"].(int64)
	losses := userAccount.Data()["losses"].(int64)
	losses++
	plays++
	client.Collection("accounts").Doc(username).Set(ctx, map[string]interface{}{
		"plays":  plays,
		"losses": losses,
	}, firestore.MergeAll)

	respondWithJson(w, http.StatusOK, map[string]string{"message": "Successfully added loss"})

}

func Test(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, http.StatusOK, map[string]string{"message": "Hello, World!"})
}
