package model

type User struct {
	UserName string `gorm:"primaryKey"`
	Password string `gorm:"not null"`
}
