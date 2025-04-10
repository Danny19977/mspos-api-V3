package models

import (
	"gorm.io/gorm"
)

type RutePlanItem struct {
	gorm.Model
	UUID string `gorm:"not null;unique" json:"uuid"`

	RoutePlanID uint      `json:"routeplan_id"`
	RoutePlan   RoutePlan `gorm:"foreignKey:RoutePlanID"`

	PosUUID uint `json:"pos_uuid" gorm:"type:varchar(255);not null"`
	Pos     Pos  `gorm:"foreignKey:PosUUID;references:UUID"`

	Status bool `json:"status"`
}
