package models

import (
	"gorm.io/gorm"
)

type Device struct {
	gorm.Model
	Id               string `gorm:"type:string" json:"id" validate:"required,string"`
	User             User   `gorm:"foreignKey:Device"`
	DeviceId         string `gorm:"type:string" json:"deviceId" validate:"required,string"`
	DeviceName       string `gorm:"type:string" json:"deviceName" validate:"required,string"`
	DeviceType       string `gorm:"type:string" json:"deviceType" validate:"required,string"`
	TokenNotfication string `gorm:"type:string" json:"tokenNotification" validate:"required,string"`
	IsTrusted        bool   `gorm:"type:string" json:"isTrusted" validate:"required,string"`
	CreatedAt        string `gorm:"type:string" json:"createdAt" validate:"required,string"`
	UpdatedAt        string `gorm:"type:string" json:"updatedAt" validate:"required,string"`
}
