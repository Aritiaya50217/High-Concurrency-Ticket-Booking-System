package model

type SeatModel struct {
	ID         uint `gorm:"primaryKey"`
	EventID    uint
	SeatNumber string
	Status     string
	Version    int
}

func (SeatModel) TableName() string {
	return "seats"
}
