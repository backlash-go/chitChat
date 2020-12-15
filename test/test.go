package test

import (
	"fmt"
	"log"
	"net/http"
)


//test form x-www-form-urlencode
func ValidateUserLogin(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	log.Println(r.Form.Get("hello"))
	log.Println(r.Form.Get("post"))
	log.Println(r.PostForm.Get("hello") + "PostForm")
	log.Println(r.PostForm.Get("post") + "PostForm")
	log.Println(r.Form)
	log.Println(r.Form.Get("thread")+"thread")

	/*
	log.Println(r.MultipartForm.Value["hello"][0])

	    log.Println(r.PostFormValue("hello"))
		fmt.Fprintln(w, r.PostFormValue("hello"))
	 */



	//操作文件
	/*
	fileHeader := r.MultipartForm.File["upload"][0]
		log.Println(fileHeader)
		file, _ := fileHeader.Open()

		s, _ := ioutil.ReadAll(file)

		//log.Println(r.PostFormValue("hello"))
		fmt.Fprintln(w, string(s))
	 */

	fmt.Fprintln(w, r.PostForm)
}


