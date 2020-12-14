package utils

import (
	"fmt"
	"gobackend/chitChat/data"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")
	name := r.PostFormValue("name")
	if email == "" || password == "" || name == "" {
		w.WriteHeader(http.StatusInternalServerError)
		s := "账号邮箱或者密码错误"
		w.Write([]byte(s))
		return
	}
	err := data.CreateUser(name, email, password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	s := "注册成功"
	w.Write([]byte(s))
	return
}

func ValidateUserLogin(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	cookie, err := r.Cookie("to_chitChat")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	getKeyFile := []interface{}{cookie.Value, "id", "email", "name"}
	value, err := data.GetHashValues(getKeyFile)
	if len(value) == 3 {
		//todo login success
		w.WriteHeader(http.StatusOK)
		s := "登陆成功"
		w.Write([]byte(s))
	}
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	if email == "" || password == "" {
		w.WriteHeader(40003)
		s := "账号或者密码错误"
		w.Write([]byte(s))

	}

	user, err := data.UserIsExisted(email)
	if err == gorm.ErrRecordNotFound && err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	if data.MdSalt(password) == user.Password {
		session, _ := data.CreateSession(user)
		keyFileValue := []interface{}{session.Uuid, "id", session.Id, "email", session.Email, "name", session.Name,}
		err := data.SetHashValues(keyFileValue)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		if err := data.SetKeyTtl(session.Uuid, 3*3600); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		cookie := &http.Cookie{
			Name:     "to_chitChat",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)
		return
	}

	log.Printf("email is:  %s  password is %s", email, password)
	fmt.Fprintf(w, "go web servers is ok")

}

func Login(w http.ResponseWriter, r *http.Request) {
	t := parseTemplateFiles("login.layout", "public.navbar", "login")
	t.Execute(w, nil)
}

//health check
func Health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "go web servers is ok")
}

func Index(w http.ResponseWriter, r *http.Request) {
	threads, err := data.GetThreads()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	cookie, err := r.Cookie("to_chitChat")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	if cookie == nil {
		generateHTML(w, threads, "layout", "public.navbar", "index")
	}
	keyFiles := []interface{}{cookie.Value, "id", "email", "name"}
	userInfo, err := data.GetHashValues(keyFiles)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	data := data.UserInfos{S: userInfo, UserThreads: threads}

	if len(userInfo) == 0 {
		generateHTML(w, threads, "layout", "public.navbar", "index")
	} else {
		generateHTML(w, data, "layout", "private.navbar", "index")

	}

	//generateHTML(w, threads, "layout", "private.navbar", "index")
}

func Err(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	//generateHTML(writer, vals.Get("msg"), "layout", "public.navbar", "error")
	generateHTML(writer, vals.Get("msg"), "layout", "private.navbar", "error")

}


