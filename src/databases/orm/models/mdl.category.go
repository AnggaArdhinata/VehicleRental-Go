package models

import "time"

type Category struct {
	ID        uint      `gorm:"primaryKey" json:"id,omitemty"`
	Name      string    `gorm:"type: varchar(255); not null" json:"name"`
	Vehicles  []Vehicle `json:"vehicle,omitempty"` //! IMPORTANT : If u will run Automigration, comment this line first !
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

//? To Sync TableName With DB
func (Category) TableName() string {
	return "category"
}

type Categories []Category
