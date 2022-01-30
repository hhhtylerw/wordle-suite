package main

import "net/http"

func One(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
