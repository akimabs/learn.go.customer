package models

type Device struct {
	BaseModel
	User             uint
	DeviceId         string `gorm:"type:string" json:"deviceId" validate:"required,string"`
	DeviceName       string `gorm:"type:string" json:"deviceName" validate:"required,string"`
	DeviceType       string `gorm:"type:string" json:"deviceType" validate:"required,string"`
	TokenNotfication string `gorm:"type:string" json:"tokenNotification" validate:"required,string"`
	IsTrusted        bool   `gorm:"type:string" json:"isTrusted" validate:"required,string"`
}
