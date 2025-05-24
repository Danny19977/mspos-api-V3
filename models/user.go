package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	UUID string `gorm:"type:text;not null;unique" json:"uuid"` // Explicitly set type:text

	Fullname        string `gorm:"not null;default:''" json:"fullname"`
	Email           string `json:"email" gorm:"unique;default:''"`
	Title           string `json:"title" gorm:"default:''"`
	Phone           string `json:"phone" gorm:"not null;unique;default:''"` // Added unique constraint
	Password        string `json:"password" validate:"required" gorm:"default:''"`
	PasswordConfirm string `json:"password_confirm" gorm:"-"`

	Role       string `json:"role" gorm:"default:''"`
	Permission string `json:"permission" gorm:"default:''"`
	Image      string `json:"image" gorm:"default:''"`
	Status     bool   `json:"status" gorm:"default:true"`
	Signature  string `json:"signature" gorm:"default:''"`

	CountryUUID  string `json:"country_uuid" gorm:"type:varchar(255);not null;default:''"`
	ProvinceUUID string `json:"province_uuid" gorm:"type:varchar(255);not null;default:''"`
	AreaUUID     string `json:"area_uuid" gorm:"type:varchar(255);not null;default:''"`
	SubAreaUUID  string `json:"subarea_uuid" gorm:"type:varchar(255);not null;default:''"`
	CommuneUUID  string `json:"commune_uuid" gorm:"type:varchar(255);not null;default:''"`

	Country  Country  `gorm:"foreignKey:CountryUUID;references:UUID"`
	Province Province `gorm:"foreignKey:ProvinceUUID;references:UUID"`
	Area     Area     `gorm:"foreignKey:AreaUUID;references:UUID"`
	SubArea  SubArea  `gorm:"foreignKey:SubAreaUUID;references:UUID"`
	Commune  Commune  `gorm:"foreignKey:CommuneUUID;references:UUID"`

	AsmUUID   string `json:"asm_uuid" gorm:"type:varchar(255);not null;default:''"`
	Asm       Asm    `gorm:"foreignKey:AsmUUID;references:UUID"`
	SupUUID   string `json:"sup_uuid" gorm:"type:varchar(255);not null;default:''"`
	Sup       Sup    `gorm:"foreignKey:SupUUID;references:UUID"`
	DrUUID    string `json:"dr_uuid" gorm:"type:varchar(255);not null;default:''"`
	Dr        Dr     `gorm:"foreignKey:DrUUID;references:UUID"`
	CycloUUID string `json:"cyclo_uuid" gorm:"type:varchar(255);not null;default:''"`
	Cyclo     Cyclo  `gorm:"foreignKey:CycloUUID;references:UUID"`

	RoutePlan []RoutePlan `gorm:"foreignKey:UserUUID;references:UUID"`
	Pos       []Pos       `gorm:"foreignKey:UserUUID;references:UUID"`
	PosForms  []PosForm   `gorm:"foreignKey:UserUUID;references:UUID"`
	// Managers  []Manager   `gorm:"foreignKey:UserUUID;references:UUID"`
	// UserLogs  []UserLogs  `gorm:"foreignKey:UserUUID;references:UUID"`

}

type UserResponse struct {
	ID           uint
	UUID         string `json:"uuid"`
	Fullname     string `json:"fullname"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Title        string `json:"title"`
	Role         string `json:"role"`
	CountryUUID  string `json:"country_uuid" gorm:"type:varchar(255);not null"`
	Country      Country
	ProvinceUUID string `json:"province_uuid" gorm:"type:varchar(255);not null"`
	Province     Province
	AreaUUID     string `json:"area_uuid" gorm:"type:varchar(255);not null"`
	Area         Area
	SubAreaUUID  string `json:"subarea_uuid" gorm:"type:varchar(255);not null"`
	SubArea      SubArea
	CommuneUUID  string `json:"commune_uuid" gorm:"type:varchar(255);not null"`
	Commune      Commune
	Permission   string `json:"permission"`
	Status       bool   `json:"status"`
	Signature    string `json:"signature"`
	CreatedAt    time.Time
	UpdatedAt    time.Time

	Asm   Asm   `gorm:"foreignKey:UserUUID;references:UUID"`
	Sup   Sup   `gorm:"foreignKey:UserUUID;references:UUID"`
	Dr    Dr    `gorm:"foreignKey:UserUUID;references:UUID"`
	Cyclo Cyclo `gorm:"foreignKey:UserUUID;references:UUID"`
}

type Login struct {
	// Email    string `json:"email" validate:"required,email"`
	// Phone    string `json:"phone" validate:"required"`
	Identifier string `json:"identifier" validate:"required"`
	Password   string `json:"password" validate:"required"`
}

func (u *User) SetPassword(p string) {
	hp, _ := bcrypt.GenerateFromPassword([]byte(p), 14)
	u.Password = string(hp)
}

func (u *User) ComparePassword(p string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(p))
	return err
}
