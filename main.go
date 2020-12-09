package main

import (
	"gobackend/chitChat/utils"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	//access static file-server
	files := http.FileServer(http.Dir("public"))
	//mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.Handle("/", files)
	//add router relative function

	mux.HandleFunc("/health", utils.Health)
	mux.HandleFunc("/login", utils.Login)
	mux.HandleFunc("/index", utils.Index)

	mux.HandleFunc("/authenticate", utils.ValidateUserLogin)
	http.ListenAndServe(":8080", mux)
}
