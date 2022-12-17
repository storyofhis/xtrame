package models

import "time"

type Category struct {
	Id        uint `gorm:"primaryKey;autoIncrement"`
	UserId    uint
	User      Users `gorm:"foreignKey:UserId"`
	Type      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
