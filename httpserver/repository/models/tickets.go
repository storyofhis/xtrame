package models

import "time"

type Tickets struct {
	Id          uint `gorm:"primaryKey;autoIncrement"`
	UserId      uint
	User        Users `gorm:"foreignKey:UserId"`
	CategoryId  uint
	Category    Category `gorm:"foreignKey:CategoryId"`
	Title       string
	Description string
	Price       float32
	Seat        int
	Duration    time.Duration
	BookedAt    time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}
