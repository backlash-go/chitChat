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
	mux.HandleFunc("/test", utils.Test)

	mux.HandleFunc("/logout", utils.Logout)
	mux.HandleFunc("/login", utils.Login)
	mux.HandleFunc("/signup",utils.SignUp)
	mux.HandleFunc("/signup_account",utils.SignUpAccount)

	mux.HandleFunc("/authenticate", utils.ValidateUserLogin)



	// defined in route_thread.go
	mux.HandleFunc("/thread/new", utils.NewThread)
	mux.HandleFunc("/thread/create", utils.CreateThread)
	mux.HandleFunc("/thread/post", utils.PostThread)
	mux.HandleFunc("/thread/read", utils.ReadThread)
	http.ListenAndServe("0.0.0.0:8080", mux)
}
