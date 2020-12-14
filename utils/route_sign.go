package utils

import (
	"gobackend/chitChat/data"
	"net/http"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "login.layout", "public.navbar", "signup")
}

func SignUpAccount(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	name := r.PostFormValue("name")
	email := r.PostFormValue("email")
	p := r.PostFormValue("password")
	password := data.MdSalt(p)
	err := data.CreateUser(name, email, password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	http.Redirect(w, r, "/login", 302)
}
