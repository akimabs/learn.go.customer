package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        string `gorm:"type:string" json:"id" validate:"required,string"`
	Name      string `gorm:"type:string" json:"name" validate:"required,string"`
	Position  string `gorm:"type:string" json:"position" validate:"required,string"`
	Email     string `gorm:"type:string" json:"email" validate:"required,string"`
	Password  string `gorm:"type:string" json:"password" validate:"required,string"`
	Company   uint
	Role      uint
	Device    []Device `gorm:"foreignKey:User;references:ID"`
	CreatedAt string   `gorm:"type:string" json:"createdAt" validate:"required,string"`
	UpdatedAt string   `gorm:"type:string" json:"updatedAt" validate:"required,string"`
}
