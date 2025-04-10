package models

import (
	// "time"

	"gorm.io/gorm"
)

type UserLogs struct {
	gorm.Model
	UUID string `gorm:"type:text;not null;unique" json:"uuid"`

	Name        string `gorm:"type:text;not null" json:"name"`
	UserUUID    string `json:"user_uuid" gorm:"type:varchar(255);not null"`
	User        User   `gorm:"foreignKey:UserUUID;references:UUID"`
	Action      string `gorm:"type:text;not null" json:"action"`
	Description string `gorm:"type:text;not null" json:"description"`
	Signature   string `json:"signature"`
}

// type UserLogPaginate struct {
// 	Id          uint      `json:"id"`
// 	ID        string    `json:"id"`
// 	Name        string    `json:"name"`
// 	Action      string    `json:"action"`
// 	Description string    `json:"description"`
// 	CreatedAt   time.Time `json:"created_at"`
// 	UserUUID    string    `json:"user_UUid"`
// 	Fullname    string    `json:"fullname"`
// 	Title       string    `json:"title"`
// }
