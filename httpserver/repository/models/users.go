package models

import "time"

type Role string

const (
	Admin    Role = "admin"
	Customer Role = "user"
)

type Users struct {
	Id        uint `gorm:"primaryKey;autoIncrement"`
	FullName  string
	NickName  string
	UserName  string
	Email     string
	Password  string
	Role      Role `gorm:"type:role;default:'user'"`
	Age       uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
