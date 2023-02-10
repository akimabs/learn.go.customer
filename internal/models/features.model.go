package models

import (
	"gorm.io/gorm"
)

type Features struct {
	gorm.Model
	Id        string `gorm:"type:string" json:"id" validate:"required,string"`
	Name      string `gorm:"type:string" json:"name" validate:"required,string"`
	Status    string `gorm:"type:enum('ACTIVE', 'INCATIVE')"`
	CreatedAt string `gorm:"type:string" json:"createdAt" validate:"required,string"`
	UpdatedAt string `gorm:"type:string" json:"updatedAt" validate:"required,string"`
}
