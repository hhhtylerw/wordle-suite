package main

import (
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	http.Handle("/", http.FileServer(http.Dir("wordle-suite")))
	http.HandleFunc("/signup", SignUp)
	http.HandleFunc("/login", Login)
	http.ListenAndServe(":"+port, nil)
}
