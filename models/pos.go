package models

import "gorm.io/gorm"

type Pos struct {
	gorm.Model
	UUID string `gorm:"not null;unique" json:"uuid"`

	Name      string `gorm:"not null;unique" json:"name"` // Celui qui vend
	Shop      string `json:"shop"`                        // Nom du shop
	Postype   string `json:"postype"`                     // Type de POS
	Gerant    string `json:"gerant"`                      // name of the onwer of the pos
	Avenue    string `json:"avenue"`
	Quartier  string `json:"quartier"`
	Reference string `json:"reference"`
	Telephone string `json:"telephone"`
	Image     string `json:"image"`

	CountryUUID  string   `json:"country_uuid" gorm:"type:varchar(255);not null"`
	Country      Country  `gorm:"foreignKey:CountryUUID;references:UUID"`
	ProvinceUUID string   `json:"province_uuid" gorm:"type:varchar(255);not null"`
	Province     Province `gorm:"foreignKey:ProvinceUUID;references:UUID"`
	AreaUUID     string   `json:"area_uuid" gorm:"type:varchar(255);not null"`
	Area         Area     `gorm:"foreignKey:AreaUUID;references:UUID"`
	SubAreaUUID  string   `json:"subarea_uuid" gorm:"type:varchar(255);not null"`
	SubArea      SubArea  `gorm:"foreignKey:SubAreaUUID;references:UUID"`
	CommuneUUID  string   `json:"commune_uuid" gorm:"type:varchar(255);not null"`
	Commune      Commune  `gorm:"foreignKey:CommuneUUID;references:UUID"`

	UserUUID string `json:"user_uuid" gorm:"type:varchar(255);not null"`
	User     User   `gorm:"foreignKey:UserUUID;references:UUID"`

	AsmUUID   string `json:"asm_uuid" gorm:"type:varchar(255);not null"`
	Asm       Asm    `gorm:"foreignKey:AsmUUID;references:UUID"`
	SupUUID   string `json:"sup_uuid" gorm:"type:varchar(255);not null"`
	Sup       Sup    `gorm:"foreignKey:SupUUID;references:UUID"`
	DrUUID    string `json:"dr_uuid" gorm:"type:varchar(255);not null"`
	Dr        Dr     `gorm:"foreignKey:DrUUID;references:UUID"`
	CycloUUID string `json:"cyclo_uuid" gorm:"type:varchar(255);not null"`
	Cyclo     Cyclo  `gorm:"foreignKey:CycloUUID;references:UUID"`

	Status    bool   `json:"status"`
	Signature string `json:"signature"`

	// PosFormItems  []PosFormItems `gorm:"foreignKey:PosUUID;references:UUID"`
	PosEquipments []PosEquipment `gorm:"foreignKey:PosUUID;references:UUID"`
	PosForms      []PosForm      `gorm:"foreignKey:PosUUID;references:UUID"`
}

type PosPaginate struct {
	Id                 string `json:"id"`
	UUID               string `json:"uuid"`
	Name               string `json:"name"`    // Celui qui vend
	Shop               string `json:"shop"`    // Nom du shop
	Manager            string `json:"manager"` // name of the onwer of the pos
	Commune            string `json:"commune"`
	Avenue             string `json:"avenue"`
	Quartier           string `json:"quartier"`
	Reference          string `json:"reference"`
	Telephone          string `json:"telephone"`
	Eparasol           bool   `json:"eparasol"`
	Etable             bool   `json:"etable"`
	Ekiosk             bool   `json:"ekiosk"`
	InputGroupSelector string `json:"inputgroupselector"`
	Cparasol           bool   `json:"cparasol"`
	Ctable             bool   `json:"ctable"`
	Ckiosk             bool   `json:"ckiosk"`
	Province           string `json:"province"`
	Area               string `json:"area"`
	Dr                 string `json:"dr"`
	Status             bool   `json:"status"`
	Signature          string `json:"signature"`
}
