package data

import (
	"gobackend/chitChat/db"
	"gobackend/chitChat/models"
)

func GetThreads() (ThreadList []*ThreadList, err error) {
	err = db.GetDB().Model(&models.Threads{}).Scan(&ThreadList).Error
	return
}

func GetThread(id int64) (Thread models.Threads, err error) {
	err = db.GetDB().Where("id = ?", id).First(&Thread).Error
	return
}

func CreateThread(topic string, userid int64) error {
	md := &models.Threads{Topic: topic, UserId: userid}
	err := db.GetDB().Model(&models.Threads{}).
		Create(md).Error
	return err

}
