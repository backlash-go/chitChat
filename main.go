package main

import (
	"fmt"
	"gobackend/chitChat/utils"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/health", Health)
	mux.HandleFunc("/login", utils.Login)
	http.ListenAndServe(":8888", mux)
}

//health check
func Health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "go web servers is ok")
}
