package models

import "time"

type Vehicle struct {
	ID          uint      `gorm:"primaryKey" json:"id,omitempty"`
	Name        string    `gorm:"type: varchar(255); not null" json:"name"`
	Location    string    `gorm:"type: varchar(255); not null" json:"location"`
	Description string    `gorm:"type: varchar(255); not null" json:"description"`
	Status      string    `gorm:"type: varchar(255); not null" json:"status"`
	Stock       int       `gorm:"not null" json:"stock"`
	Category_Id int       `gorm:"not null" json:"category_id"`
	Category    Category  `gorm:"foreignKey:Category_Id; references:ID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"-"`
	Price       int       `gorm:"not null" json:"price"`
	Rating      int       `json:"rating"`
	Image       string    `gorm:"type: varchar(255)" json:"image"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (Vehicle) TableName() string {
	return "vehicle"
}

type Vehicles []Vehicle
