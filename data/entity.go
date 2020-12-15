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
	Id    uint64
	Uuid  string
	Name  string
	Email string
}

type ThreadInfos struct {
	ID        uint64     `gorm:"column:id"form:"id" json:"id"`
	Topic     string     `gorm:type:text","column:topic", form:"topic",json:"topic"`
	UserId    uint64     `gorm:"column:user_id" form:"user_id" json:"user_id"`
	UserName  string     `json:"user_name"`
	CreatedAt *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at"`

}

type ThreadList struct {
	ID        uint64     `gorm:"column:id"form:"id" json:"id"`
	Topic     string     `gorm:type:text","column:topic", form:"topic",json:"topic"`
	UserId    uint64     `gorm:"column:user_id" form:"user_id" json:"user_id"`
	CreatedAt *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at"`
}
