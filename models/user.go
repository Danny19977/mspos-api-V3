package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	UUID string `gorm:"type:text;not null;unique" json:"uuid"` // Explicitly set type:text

	Fullname        string `gorm:"not null" json:"fullname"`
	Email           string `json:"email" gorm:"unique"`
	Title           string `json:"title"`
	Phone           string `json:"phone" gorm:"not null;unique"` // Added unique constraint
	Password        string `json:"password" validate:"required"`
	PasswordConfirm string `json:"password_confirm" gorm:"-"`

	Role       string `json:"role"`
	Permission string `json:"permission"`
	Image      string `json:"image"`
	Status     bool   `json:"status"`
	Signature  string `json:"signature"`

	CountryUUID  string `json:"country_uuid" gorm:"type:varchar(255);not null"`
	ProvinceUUID string `json:"province_uuid" gorm:"type:varchar(255);not null"`
	AreaUUID     string `json:"area_uuid" gorm:"type:varchar(255);not null"`
	SubAreaUUID  string `json:"subarea_uuid" gorm:"type:varchar(255);not null"`
	CommuneUUID  string `json:"commune_uuid" gorm:"type:varchar(255);not null"`

	Country  Country  `gorm:"foreignKey:CountryUUID;references:UUID"`
	Province Province `gorm:"foreignKey:ProvinceUUID;references:UUID"`
	Area     Area     `gorm:"foreignKey:AreaUUID;references:UUID"`
	SubArea  SubArea  `gorm:"foreignKey:SubAreaUUID;references:UUID"`
	Commune  Commune  `gorm:"foreignKey:CommuneUUID;references:UUID"`

	Asm   Asm   `gorm:"foreignKey:UserUUID;references:UUID"`
	Sup   Sup   `gorm:"foreignKey:UserUUID;references:UUID"`
	Dr    Dr    `gorm:"foreignKey:UserUUID;references:UUID"`
	Cyclo Cyclo `gorm:"foreignKey:UserUUID;references:UUID"`

	// Asms   []Asm   `gorm:"foreignKey:UserUUID"`
	// Sups   []Sup   `gorm:"foreignKey:UserUUID"`
	// Drs    []Dr    `gorm:"foreignKey:UserUUID"`
	// Cyclos []Cyclo `gorm:"foreignKey:UserUUID"`

	RoutePlan []RoutePlan `gorm:"foreignKey:UserUUID;references:UUID"`
	Managers  []Manager   `gorm:"foreignKey:UserUUID;references:UUID"`
	UserLogs  []UserLogs  `gorm:"foreignKey:UserUUID;references:UUID"`

	CountryName  string `json:"country_name" gorm:"-"`
	ProvinceName string `json:"province_name" gorm:"-"`
	AreaName     string `json:"area_name" gorm:"-"`
	SubAreaName  string `json:"subarea_name" gorm:"-"`
	CommuneName  string `json:"commune_name" gorm:"-"` 
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

	// CountryName  string `json:"country_name" gorm:"-"`
	// ProvinceName string `json:"province_name" gorm:"-"`
	// AreaName     string `json:"area_name" gorm:"-"`
	// SubAreaName  string `json:"subarea_name" gorm:"-"`
	// CommuneName  string `json:"commune_name" gorm:"-"` 

	Asm          Asm   
	Sup          Sup   
	Dr           Dr    
	Cyclo        Cyclo 
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
