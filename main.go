package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

var (
	ctx    context.Context = context.Background()
	client *firestore.Client
	salt   string
)

/*type Account struct {
	Username string   `json:"username"`
	Friends  []string `json:"friends"`
	Plays    int      `json:"plays"`
	Wins     int      `json:"wins"`
	Losses   int      `json:"losses"`
}*/

func init() {
	godotenv.Load()
	salt = os.Getenv("SALT")
}

func main() {
	// Initialize Firestore client
	var err error
	opt := option.WithCredentialsFile("wordle-suite-firebase-adminsdk-n1yxj-ed63a51e95.json")
	client, err = firestore.NewClient(ctx, "wordle-suite", opt)
	if err != nil {
		fmt.Println(err)
	}
	defer client.Close()

	// Initialize HTTP server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.Handle("/", http.FileServer(http.Dir("wordle-suite")))
	http.HandleFunc("/signup", SignUp)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/account", Account)
	http.HandleFunc("/friends", Friends)
	http.HandleFunc("/addfriend", AddFriend)
	http.HandleFunc("/addwin", AddWin)
	http.HandleFunc("/addloss", AddLoss)
	http.HandleFunc("/test", Test)
	http.ListenAndServe(":"+port, nil)
}

func getSHA256(input string) string {
	hasher := sha256.New()
	hasher.Write([]byte(input))
	output := hex.EncodeToString(hasher.Sum(nil))
	return string(output)
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func contains(s []interface{}, e interface{}) bool {
	for _, a := range s {
		if a.(string) == e.(string) {
			return true
		}
	}
	return false
}
