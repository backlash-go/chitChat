package data

import "gobackend/chitChat/models"

type UserInfos struct {
	S           []string
	UserThreads []models.Threads
}

type UserCache struct {
	Id    uint64
	Uuid  string
	Name  string
	Email string
}
