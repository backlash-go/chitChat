package data

import (
	"gobackend/chitChat/db"
	"gobackend/chitChat/models"
)

func GetThreads() (threads []models.Threads, err error) {
	err = db.GetDB().Model(&models.Threads{}).Find(&threads).Error
	return
}
