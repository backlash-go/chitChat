package utils

import (
	"fmt"
	"gobackend/chitChat/data"
	"gorm.io/gorm"
	"html/template"
	"log"
	"net/http"
	"strconv"
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
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")
	log.Printf("email is %s and password is %s", email, password)
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
		log.Printf("uuid is %s", session.Uuid)
		err := data.SetHashValues(keyFileValue)
		if err != nil {
			log.Println(err.Error())
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
		http.Redirect(w, r, "/", 302)
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	t := parseTemplateFiles("login.layout", "public.navbar", "login")
	t.Execute(w, nil)
}

//health check
func Health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "go web servers is ok")
}

func Test(w http.ResponseWriter, r *http.Request) {
	p := &[]struct {
		Name string
		Age  int
	}{
		{"longshuai", 22},
		{"xxb", 33},
	}

	t1, err := template.ParseFiles("templates/test.html")
	if err != nil {
		panic(err)
	}
	t1.Execute(w, p)
}

func Index(w http.ResponseWriter, r *http.Request) {
	userList, err := data.GetUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	threadsList, err := data.GetThreads()
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
	log.Printf("Index cookie value is %s", cookie.Value)
	keyFiles := []interface{}{cookie.Value, "id", "email", "name"}
	userInfo, err := data.GetHashValues(keyFiles)
	log.Printf("userInfo is %s len is %d\n", userInfo, len(userInfo))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	m1 := make(map[int64]string)
	for _, va := range userList {
		m1[va.ID] = va.Name
	}

	items := make([]*data.ThreadInfos, 0, 0)
	for _, value := range threadsList {
		items = append(items, &data.ThreadInfos{
			ID:        value.ID,
			Topic:     value.Topic,
			UserId:    value.UserId,
			UserName:  m1[value.UserId],
			CreatedAt: value.CreatedAt,
		})
	}
	log.Printf("data is %s", items)
	log.Printf("threads is %s", items)
	//log.Printf("threads is %s", threadsInfo[1].Topic)

	if userInfo[0] == "" || userInfo[1] == "" || userInfo[2] == "" {
		generateHTML(w, items, "layout", "public.navbar", "index")
	} else {
		generateHTML(w, items, "layout", "private.navbar", "index")
	}

	//generateHTML(w, threads, "layout", "private.navbar", "index")
}

func Err(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	//generateHTML(writer, vals.Get("msg"), "layout", "public.navbar", "error")
	generateHTML(writer, vals.Get("msg"), "layout", "private.navbar", "error")

}

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("to_chitChat")
	if err != http.ErrNoCookie {
		http.Redirect(w, r, "/", 302)
	}
	if _, err := data.DelHash(cookie.Value); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	http.Redirect(w, r, "/", 302)

}

func NewThread(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("to_chitChat")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	log.Printf("PostThread cookie value is %s", cookie.Value)
	keyFiles := []interface{}{cookie.Value, "id", "email", "name"}
	userInfo, err := data.GetHashValues(keyFiles)
	log.Printf("userInfo is %s len is %d\n", userInfo, len(userInfo))
	if userInfo[0] == "" || userInfo[1] == "" || userInfo[2] == "" {
		http.Redirect(w, r, "/login", 302)
	}

	generateHTML(w, userInfo[1], "layout", "private.navbar", "new.thread")

	return
}

func CreateThread(w http.ResponseWriter, r *http.Request) {
	userId, _ := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	topic := r.PostFormValue("name")

	err := data.CreateThread(topic, userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	http.Redirect(w, r, "/", 302)

}
