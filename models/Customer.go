package models

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	ID              string         `json:"id" `
	Name            string         `json:"name" validate:"required" gorm:"not null"`
	Gender          string         `json:"gender" validate:"required"  gorm:"not null;type:varchar(10)"`
	Amount          int            `json:"amount" validate:"required,min=1000000,max=10000000" gorm:"not null"`
	Tenor           int            `json:"tenor" validate:"required" gorm:"not null"`
	BirthDate       time.Time      `json:"birth_date" validate:"required" gorm:"not null"`
	Address         string         `json:"address" validate:"required" gorm:"not null"`
	Nationality     string         `json:"nationality" validate:"required" gorm:"not null;type:varchar(10)"`
	AddressProvince string         `json:"address_province"  gorm:"not null;type:varchar(10)"`
	Ktp             string         `json:"ktp" gorm:"unique;not null" validate:"required,len=16,numeric"`
	Status          string         `json:"status" gorm:"not null;type:varchar(50)"`
	CreatedAt       time.Time      `json:"-"`
	UpdatedAt       time.Time      `json:"-"`
	DeletedAt       gorm.DeletedAt `json:"-" gorm:"index"`
}
