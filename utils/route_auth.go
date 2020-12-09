package utils

import (
	"html/template"
	"net/http"
)

//func Login(w http.ResponseWriter, r *http.Request) {
//    t, err := template.ParseFiles("./template/login.html")
//    if err != nil {
//    	w.WriteHeader(http.StatusInternalServerError)
//    	w.Write([]byte(err.Error()))
//	}
//	t.Execute(w, nil)
//}
func Login(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/test.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	t.Execute(w, nil)
}