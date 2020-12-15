package data

import (
	"gobackend/chitChat/db"
	"gobackend/chitChat/models"
)

func GetThreads() (ThreadList []*ThreadList, err error) {
	err = db.GetDB().Model(&models.Threads{}).Scan(&ThreadList).Error
	return
}


