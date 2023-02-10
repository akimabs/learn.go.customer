package models

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID          string `gorm:"type:string" json:"id" validate:"required,string"`
	User        []User `gorm:"foreignKey:Role;references:ID"`
	Name        string `gorm:"type:string" json:"name" validate:"required,string"`
	Description string `gorm:"type:string" json:"description" validate:"required,string"`
	CreatedAt   string `gorm:"type:string" json:"createdAt" validate:"required,string"`
	UpdatedAt   string `gorm:"type:string" json:"updatedAt" validate:"required,string"`
}
