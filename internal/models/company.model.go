package models

import (
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	Id          string `gorm:"type:string" json:"id" validate:"required,string"`
	Name        string `gorm:"type:string" json:"name" validate:"required,string"`
	Description string `gorm:"type:string" json:"description" validate:"required,string"`
	User        []User `gorm:"foreignKey:Company"`
	CreatedAt   string `gorm:"type:string" json:"createdAt" validate:"required,string"`
	UpdatedAt   string `gorm:"type:string" json:"updatedAt" validate:"required,string"`
}
