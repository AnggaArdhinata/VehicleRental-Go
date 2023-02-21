package models

import "time"

type Vehicle struct {
	ID          uint      `gorm:"primaryKey" json:"id,omitempty"`
	Name        string    `gorm:"type: varchar(255); not null" json:"name,omitempty"`
	Location    string    `gorm:"type: varchar(255); not null" json:"location,omitempty"`
	Description string    `gorm:"type: varchar(255); not null" json:"description,omitempty"`
	Status      string    `gorm:"type: varchar(255); not null" json:"status,omitempty"`
	Stock       int       `gorm:"not null" json:"stock,omitempty"`
	Category_Id int       `gorm:"not null" json:"category_id,omitempty"`
	Category    Category  `gorm:"foreignKey:Category_Id; references:ID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"-"`
	Price       int       `gorm:"not null" json:"price"`
	Rating      int       `json:"rating,omitempty"`
	Image       string    `gorm:"type: varchar(255)" json:"image,omitempty"`
	ImageUuid   string    `gorm:"type: varchar(255)" json:"-"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (Vehicle) TableName() string {
	return "vehicle"
}

type Vehicles []Vehicle
