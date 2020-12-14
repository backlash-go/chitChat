package main

import (
	"gobackend/chitChat/db"
	"gobackend/chitChat/utils"
	"net/http"
)

func main() {

    //initial db
	db.Init()

	mux := http.NewServeMux()
	//access static file-server
	files := http.FileServer(http.Dir("public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	//mux.Handle("/", files)
	//add router relative function
	mux.HandleFunc("/", utils.Index)
	mux.HandleFunc("/err", utils.Err)
	mux.HandleFunc("/health", utils.Health)

	mux.HandleFunc("/login", utils.Login)
	mux.HandleFunc("/signup",utils.SignUp)
	mux.HandleFunc("/signup_account",utils.SignUpAccount)


	mux.HandleFunc("/authenticate", utils.ValidateUserLogin)
	http.ListenAndServe("0.0.0.0:8080", mux)
}
