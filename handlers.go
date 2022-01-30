package main

import "net/http"

func SignUp(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
func Login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
