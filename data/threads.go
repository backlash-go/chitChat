package data

import (
	"gobackend/chitChat/db"
	"gobackend/chitChat/models"
)

func GetThreads() (threads []models.Threads, err error) {
	err = db.GetDB().Find(&threads).Error
	return
}
