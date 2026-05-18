package model

type EventModel struct {
	ID    uint `gorm:"primaryKey;autoIncrement"`
	Name  string
	Seats []SeatModel `gorm:"foreignKey:EventID"`
}

func (EventModel) TableName() string {
	return "events"
}
