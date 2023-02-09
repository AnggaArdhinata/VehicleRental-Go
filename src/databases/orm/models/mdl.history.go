package models

import "time"

type History struct {
	Id             uint        `gorm:"primaryKey" json:"id,omitempty"`
	Reservation_Id int         `gorm:"not null" json:"reservation_id"`
	Reservation    Reservation `gorm:"foreignKey:Reservation_Id; references:Id; constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"-"`
	Status         string      `gorm:"type: varchar(255)" json:"status"`
	CreatedAt      time.Time   `json:"created_at"`
	UpdatedAt      time.Time   `json:"updated_at"`
}

func (History) TableName() string {
	return "history"
}

type Historis []History
