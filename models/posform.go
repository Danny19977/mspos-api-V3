package models

import "gorm.io/gorm"

type PosForm struct {
	gorm.Model

	UUID string `gorm:"not null;unique" json:"uuid"`

	Price   int    `gorm:"default:0" json:"price"`
	Comment string `json:"comment"`

	Latitude  float64 `json:"latitude"`  // Latitude of the user
	Longitude float64 `json:"longitude"` // Longitude of the user
	Signature string `json:"signature"`

	PosUUID      string `json:"pos_uuid" gorm:"type:varchar(255);not null"`
	
	CountryUUID  string `json:"country_uuid" gorm:"type:varchar(255);not null"`
	ProvinceUUID string `json:"province_uuid" gorm:"type:varchar(255);not null"`
	AreaUUID     string `json:"area_uuid" gorm:"type:varchar(255);not null"`
	SubAreaUUID  string `json:"sub_area_uuid" gorm:"type:varchar(255);not null"`
	CommuneUUID  string `json:"commune_uuid" gorm:"type:varchar(255);not null"`

	AsmUUID   string `json:"asm_uuid" gorm:"type:varchar(255);not null"`
	SupUUID   string `json:"sup_uuid" gorm:"type:varchar(255);not null"`
	DrUUID    string `json:"dr_uuid" gorm:"type:varchar(255);not null"`
	CycloUUID string `json:"cyclo_uuid" gorm:"type:varchar(255);not null"`

	Pos      Pos      `gorm:"foreignKey:PosUUID;references:UUID"`
	
	Country  Country  `gorm:"foreignKey:CountryUUID;references:UUID"`
	Province Province `gorm:"foreignKey:ProvinceUUID;references:UUID"`
	Area     Area     `gorm:"foreignKey:AreaUUID;references:UUID"`
	SubArea  SubArea  `gorm:"foreignKey:SubAreaUUID;references:UUID"`
	Commune  Commune  `gorm:"foreignKey:CommuneUUID;references:UUID"`

	ASM   Asm   `gorm:"foreignKey:AsmUUID;references:UUID"`
	Sup   Sup   `gorm:"foreignKey:SupUUID;references:UUID"`
	Dr    Dr    `gorm:"foreignKey:DrUUID;references:UUID"`
	Cyclo Cyclo `gorm:"foreignKey:CycloUUID;references:UUID"`
	
	Sync bool `json:"sync"`

	PosFormItems []PosFormItems `gorm:"foreignKey:PosFormUUID;references:UUID"`
	// Sm1      int64 `gorm:"default: 0" json:"sm1"`
	// Br       int64 `gorm:"default: 0" json:"br"`
	// Br1      int64 `gorm:"default: 0" json:"br1"`
	// Tf       int64 `gorm:"default: 0" json:"tf"`
	// Tf1      int64 `gorm:"default: 0" json:"tf1"`
	// Bon      int64 `gorm:"default: 0" json:"bon"`
	// Bon1     int64 `gorm:"default: 0" json:"bon1"`
	// Bonus    int64 `gorm:"default: 0" json:"bonus"`
	// Bonus1   int64 `gorm:"default: 0" json:"bonus1"`
	// Pmg      int64 `gorm:"default: 0" json:"pmg"`
	// Pmg1     int64 `gorm:"default: 0" json:"pmg1"`
	// Pe       int64 `gorm:"default: 0" json:"pe"`
	// Pe1      int64 `gorm:"default: 0" json:"pe1"`
	// Shik     int64 `gorm:"default: 0" json:"shik"`
	// Shik1    int64 `gorm:"default: 0" json:"shik1"`
	// Ab       int64 `gorm:"default: 0" json:"ab"`
	// Ab1      int64 `gorm:"default: 0" json:"ab1"`
	// Caesar   int64 `gorm:"default: 0" json:"caesar"`
	// Caesar1  int64 `gorm:"default: 0" json:"caesar1"`
	// Ck       int64 `gorm:"default: 0" json:"ck"`
	// Ck1      int64 `gorm:"default: 0" json:"ck1"`
	// Sfks     int64 `gorm:"default: 0" json:"sfks"`
	// Sfks1    int64 `gorm:"default: 0" json:"sfks1"`
	// Winston  int64 `gorm:"default: 0" json:"winston"`
	// Winston1 int64 `gorm:"default: 0" json:"winston1"`
	// Sf       int64 `gorm:"default: 0" json:"sf"`
	// Sf1      int64 `gorm:"default: 0" json:"sf1"`
	// Cm       int64 `gorm:"default: 0" json:"cm"`
	// Cm1      int64 `gorm:"default: 0" json:"cm1"`
	// Om       int64 `gorm:"default: 0" json:"om"`
	// Om1      int64 `gorm:"default: 0" json:"om1"`
	// Of       int64 `gorm:"default: 0" json:"of"`
	// Of1      int64 `gorm:"default: 0" json:"of1"`
	// Rmr      int64 `gorm:"default: 0" json:"rmr"`
	// Rmr1     int64 `gorm:"default: 0" json:"rmr1"`
	// Rms      int64 `gorm:"default: 0" json:"rms"`
	// Rms1     int64 `gorm:"default: 0" json:"rms1"`
	// Arf      int64 `gorm:"default: 0" json:"arf"`
	// Arf1     int64 `gorm:"default: 0" json:"arf1"`
	// Ptmn     int64 `gorm:"default: 0" json:"ptmn"`
	// Ptmn1    int64 `gorm:"default: 0" json:"ptmn1"`
	// Monento  int64 `gorm:"default: 0" json:"monento"`
	// Monento1 int64 `gorm:"default: 0" json:"monento1"`
	// Stella   int64 `gorm:"default: 0" json:"stella"`
	// Stella1  int64 `gorm:"default: 0" json:"stella1"`
	// Chikt    int64 `gorm:"default: 0" json:"chikt"`
	// Chikt1   int64 `gorm:"default: 0" json:"chikt1"`
	// Asp      int64 `gorm:"default: 0" json:"asp"`
	// Asp1     int64 `gorm:"default: 0" json:"asp1"`
	// Ld       int64 `gorm:"default: 0" json:"ld"`
	// Ld1      int64 `gorm:"default: 0" json:"ld1"`
	// Lgd      int64 `gorm:"default: 0" json:"lgd"`
	// Lgd1     int64 `gorm:"default: 0" json:"lgd1"`
	// Frm      int64 `gorm:"default: 0" json:"frm"`
	// Frm1     int64 `gorm:"default: 0" json:"frm1"`

}

