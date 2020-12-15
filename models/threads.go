package models

import "time"

//帖子表
type Threads struct {
	ID        uint64     `gorm:"column:id"form:"id" json:"id"`
	Topic     string     `gorm:type:text","column:topic", form:"topic",json:"topic"`
	UserId    uint64     `gorm:"column:user_id" form:"user_id" json:"user_id"`
	CreatedAt *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at" form:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" form:"deleted_at" json:"deleted_at"`
}

//帐号密码登陆验证
func (m *Threads) TableName() string {
	return "threads"
}
