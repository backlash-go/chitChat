package utils

import (
	"fmt"
	"gobackend/chitChat/data"
	"log"
	"net/http"
	"strconv"
)

func ReadThread(w http.ResponseWriter, r *http.Request) {
	threadID := r.URL.Query().Get("id")
	thID, _ := strconv.ParseInt(threadID, 10, 64)

	item := &data.ThreadPostUserList{}

	thread, err := data.GetThread(thID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	user, err := data.SelectUserInfo(thread.UserId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	item.UserID = thread.UserId
	item.Topic = thread.Topic
	item.UserName = user.Name
	item.ThreadCreatedAt = thread.CreatedAt
	item.ThreadID = thread.ID

	userList, err := data.GetUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	postList, err := data.GetPosts(thread.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	m1 := make(map[int64]string)
	for _, va := range userList {
		m1[va.ID] = va.Name
	}

	for _, v := range userList {
		m1[v.ID] = v.Name
	}

	for _, value := range postList {
		item.Posts = append(item.Posts, &data.PostsList{
			ID:        value.ID,
			UserId:    value.UserID,
			UserName:  m1[value.UserID],
			Body:      value.Body,
			CreatedAt: value.CreatedAt,
		})
	}
	cookie, err := r.Cookie("to_chitChat")
	if err == http.ErrNoCookie {
		generateHTML(w, item, "layout", "public.navbar", "public.thread")
		return
	}
	log.Printf("readthread cookie value is %s", cookie.Value)
	uidExisted, err := data.IsExisted(cookie.Value)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	if uidExisted == 0 {
		http.Redirect(w, r, "/login", 302)
	} else {
		generateHTML(w, item, "layout", "private.navbar", "private.thread")
	}

}

func PostThread(w http.ResponseWriter, r *http.Request) {
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

	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	body := r.PostFormValue("body")
	userid, _ := strconv.ParseInt(userInfo[0], 10, 64)
	threadId, _ := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	log.Printf("body is %s , userid is %d threadId is %d", body, userid, threadId)
	if err := data.CreatePost(userid, threadId, body); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	url := fmt.Sprintf("/thread/read?id=%d", threadId)
	http.Redirect(w, r, url, 302)

}
