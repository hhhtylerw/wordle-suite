package main

import (
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	http.Handle("/", http.FileServer(http.Dir("wordle-suite")))
	http.HandleFunc("/test", One)
	http.ListenAndServe(":"+port, nil)
}
