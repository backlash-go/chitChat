package models

import (
	"time"
)

//评论表
type Posts struct {
	//ID        int64      `gorm:"column:id" form:"id" json:"id"`
	Body      string `gorm:type:text","column:body", form:"body",json:"body"`
	UserID    int64  `gorm:"column:user_id" form:"user_id" json:"user_id"`
	ThreadsID int64  `gorm:"column:threads_id" form:"threads_id" json:"threads_id"`
	//CreatedAt time.Time  `gorm:"column:created_at" form:"created_at" json:"created_at"`
	//UpdatedAt time.Time `gorm:"column:updated_at" form:"updated_at" json:"updated_at"`
	//DeletedAt time.Time `gorm:"column:deleted_at" form:"deleted_at" json:"deleted_at"`
	ID        int64 `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

//帐号密码登陆验证
func (m *Posts) TableName() string {
	return "posts"
}
