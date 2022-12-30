package models

type Emails struct {
	Id      uint `gorm:"primaryKey;autoIncrement"`
	ToEmail []string
	CcEmail []string
	Subject string
	Message string
}
