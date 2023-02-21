package models

import "time"

type Reservation struct {
	Id            uint      `gorm:"primaryKey" json:"id,omitempty"`
	User_Id       uint      `gorm:"not null" json:"user_id,omitempty"`
	User          User      `gorm:"foreignKey:User_Id; references:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"user_data,omitempty"`
	VehicleID     int       `gorm:"not null" json:"vehicle_id"`
	Vehicle       Vehicle   `gorm:"foreignKey:VehicleID; references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"vechicle_data,omitempty"`
	Qty           int       `gorm:"not null" json:"qty,omitempty"`
	Start_Date    string    `gorm:"type: varchar(255); not null" json:"start_date,omitempty"`
	Return_Date   string    `gorm:"type: varchar(255); not null" json:"return_date,omitempty"`
	Total_Payment int       `json:"total_payment,omitempty"`
	Isbooked      bool      `gorm:"default: false" json:"isbooked,omitempty"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (Reservation) TableName() string {
	return "reservation"
}

type Reservations []Reservation
