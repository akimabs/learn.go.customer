package models

import (
	"gorm.io/gorm"
)

type Roles struct {
	gorm.Model
	Id          string  `gorm:"type:string" json:"id" validate:"required,string"`
	User        []Users `gorm:"foreignKey:Role"`
	Name        string  `gorm:"type:string" json:"name" validate:"required,string"`
	Description string  `gorm:"type:string" json:"description" validate:"required,string"`
	CreatedAt   string  `gorm:"type:string" json:"createdAt" validate:"required,string"`
	UpdatedAt   string  `gorm:"type:string" json:"updatedAt" validate:"required,string"`
}
