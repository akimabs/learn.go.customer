package models

import (
	"gorm.io/gorm"
)

type Status string

const (
	ACTIVE   Status = "ACTIVE"
	INCATIVE Status = "INCATIVE"
)

type Feature struct {
	gorm.Model
	ID        string `gorm:"type:string" json:"id" validate:"required,string"`
	Name      string `gorm:"type:string" json:"name" validate:"required,string"`
	Status    Status
	CreatedAt string `gorm:"type:string" json:"createdAt" validate:"required,string"`
	UpdatedAt string `gorm:"type:string" json:"updatedAt" validate:"required,string"`
}
