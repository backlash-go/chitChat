package data

import (
	"gobackend/chitChat/models"
	"time"
)

type UserInfos struct {
	S           []string
	UserThreads []*models.Threads
}

type UserCache struct {
	ID    int64
	Uuid  string
	Name  string
	Email string
}

type ThreadInfos struct {
	ID        int64      `gorm:"column:id"form:"id" json:"id"`
	Topic     string     `gorm:type:text","column:topic", form:"topic",json:"topic"`
	UserId    int64      `gorm:"column:user_id" form:"user_id" json:"user_id"`
	UserName  string     `json:"user_name"`
	CreatedAt *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at"`
}

type ThreadList struct {
	ID        int64      `gorm:"column:id"form:"id" json:"id"`
	Topic     string     `gorm:type:text","column:topic", form:"topic",json:"topic"`
	UserId    int64      `gorm:"column:user_id" form:"user_id" json:"user_id"`
	CreatedAt *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at"`
}

type UserList struct {
	ID   int64
	Name string
}

type ThreadPostUserList struct {
	UserID          int64
	Topic           string
	UserName        string
	ThreadID        int64
	ThreadCreatedAt *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at"`
	Posts           []*PostsList
}

type PostsList struct {
	ID        int64  `gorm:"column:id"form:"id" json:"id"`
	UserId    int64  `gorm:"column:user_id" form:"user_id" json:"user_id"`
	UserName  string `json:"user_name"`
	Body      string
	CreatedAt time.Time `gorm:"column:created_at" form:"created_at" json:"created_at"`
}
