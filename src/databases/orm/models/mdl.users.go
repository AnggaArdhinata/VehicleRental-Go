package models

import "time"

type User struct {
	Id        uint      `gorm:"primaryKey" json:"id,omitempty" valid:"-"`
	Name      string    `gorm:"type: varchar(255); not null" json:"name" valid:"-"`
	Email     string    `gorm:"type: varchar(255); not null" json:"email" valid:"email"`
	Password  string    `gorm:"type: varchar(255); not null" json:"password" valid:"-"`
	Birthday  string    `gorm:"type: date; not null" json:"birthday" valid:"-"`
	Gender    string    `gorm:"type: varchar(255); not null" json:"gender" valid:"-"`
	Address   string    `gorm:"type: varchar(255); not null" json:"address" valid:"-"`
	Phone     string    `gorm:"type: varchar(255)" json:"phone" valid:"-"`
	Image     string    `gorm:"type: varchar(255)" json:"image" valid:"-"`
	ImgId     string    `gorm:"type: varchar(255)" json:"-" valid:"-"`
	Role      string    `gorm:"type: varchar(15); default: user" json:"role,omitempty" valid:"-"`
	CreatedAt time.Time `json:"created_at" valid:"-"`
	UpdatedAt time.Time `json:"updated_at" valid:"-"`
}

func (User) TableName() string {
	return "users"
}

type Users []User
