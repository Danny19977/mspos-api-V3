package models

import (
	"gorm.io/gorm"
)

type PosFormItems struct {
	gorm.Model

	UUID string `gorm:"not null;unique" json:"uuid"`

	NumberFarde string `gorm:"not null" json:"number_farde"` // NUMBER Farde
	Counter     string `gorm:"not null" json:"counter"`      // Allows to calculate the Sum of the ND Dashboard

	PosFormUUID string `json:"posform_uuid" gorm:"type:varchar(255);not null"` // Foreign key (belongs to), tag `index` will create index for this column
	BrandUUID   string `json:"brand_uuid" gorm:"type:varchar(255);not null"`   // Foreign key (belongs to), tag `index` will create index for this column
	PosUUID     uint   `json:"pos_uuid" gorm:"type:varchar(255);not null"`   // Foreign key (belongs to), tag `index` will create index for this column

	PosForm PosForm `gorm:"foreignKey:PosFormUUID;references:UUID"` // POS Form of the POS
	Pos     Pos     `gorm:"foreignKey:PosUUID;references:UUID"`     // POS of the POS
	Brand   Brand   `gorm:"foreignKey:BrandUUID;references:UUID"`     // Brand of the POS

}
