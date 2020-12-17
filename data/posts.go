package data

import (
	"gobackend/chitChat/db"
	"gobackend/chitChat/models"
)

func GetPosts(threadID int64) (Posts []*models.Posts, err error) {
	err = db.GetDB().Model(&models.Posts{}).Where("threads_id = ?", threadID).Find(&Posts).Error
	return
}

func CreatePost(userId, threadId int64, body string) error {
	md :=&models.Posts{UserID: userId, ThreadsID: threadId, Body: body}
	err := db.GetDB().Model(&models.Posts{}).
		Create(md).Error
	return err

}

