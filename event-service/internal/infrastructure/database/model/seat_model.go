package model

type EventModel struct {
	ID    uint `gorm:"primaryKey;autoIncrement"`
	Name  string
	Seats []SeatModel `gorm:"foreignKey:EventID"`
}

type SeatModel struct {
	ID         uint `gorm:"primaryKey"`
	EventID    uint
	SeatNumber string
	Status     string
}

func (EventModel) TableName() string {
	return "events"
}

func (SeatModel) TableName() string {
	return "seats"
}
