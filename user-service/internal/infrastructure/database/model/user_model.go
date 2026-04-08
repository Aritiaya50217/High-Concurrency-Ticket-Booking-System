package model

import "time"

type UsersModel struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
