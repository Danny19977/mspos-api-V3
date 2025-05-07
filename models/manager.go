package models

import "gorm.io/gorm"

type Manager struct {
	gorm.Model
	UUID string `gorm:"not null;unique" json:"uuid"`

	Title       string  `gorm:"not null" json:"title"` // Example Head of Sales, Support, Manager, etc
	CountryUUID string    `json:"country_uuid" gorm:"type:varchar(255);not null"`
	Country     Country `gorm:"foreignKey:CountryUUID;references:UUID"`
	UserUUID    string    `json:"user_uuid"` // Corrected field name
	User        User    `gorm:"foreignKey:UserUUID"`
	Signature   string  `json:"signature"`
}

