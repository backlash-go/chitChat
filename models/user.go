package models

import "time"


//用户表
type User struct {
	ID        int64     `gorm:"column:id" form:"id" json:"id"`
	Name      string     `gorm:"column:name" form:"name" json:"name"`
	Password  string     `gorm:"column:password" form:"password" json:"password"`
	Email     string     `gorm:"column:email", form:"email",json:"email"`
	CreatedAt *time.Time `gorm:"column:created_at" form:"created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at" form:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" form:"deleted_at" json:"deleted_at"`
}

//帐号密码登陆验证
func (m *User) TableName() string {
	return "user"

}
