package model

import (
	"time"
)

type UsersModel struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	FirstName string `gorm:"column:first_name"`
	LastName  string `gorm:"column:last_name"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (UsersModel) TableName() string {
	return "users"
}
