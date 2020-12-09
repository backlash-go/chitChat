package utils

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func ValidateUserLogin(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "r.ParseForm is failed")
	}
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")
	log.Printf("email is:  %s  password is %s", email, password)
	fmt.Fprintf(w, "go web servers is ok")

}

func Login(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/test.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	log.Println(t.Name())
	t.Execute(w, nil)
}

//health check
func Health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "go web servers is ok")
}

func Index(w http.ResponseWriter, r *http.Request) {
	files := []string{"./templates/public.navbar.html", "./templates/index.html"}
	t := template.Must(template.ParseFiles(files...))
	t.ExecuteTemplate(w, "layout", nil)

}
