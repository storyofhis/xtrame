package models

type Tickets struct {
	Id   uint `gorm:"primaryKey;autoIncrement"`
	Name string
}
