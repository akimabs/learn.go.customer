package models

import (
	"time"
)

type Role struct {
	BaseModel
	// User        []User    `gorm:"foreignKey:Role;references:id"`
	Name        string    `gorm:"type:string" json:"name" validate:"required,string"`
	Description string    `gorm:"type:string" json:"description" validate:"required,string"`
	CreatedAt   time.Time `gorm:"type:string" json:"createdAt autoCreateTime"`
	UpdatedAt   time.Time `gorm:"type:string" json:"updatedAt autoCreateTime"`
}
