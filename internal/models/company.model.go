package models

type Company struct {
	ID          string `gorm:"type:char(36)" json:"id"`
	CreatedAt   string `gorm:"type:string" json:"createdAt" autoUpdateTime:"milli"`
	UpdatedAt   string `gorm:"type:string" json:"updatedAt" autoUpdateTime:"milli"`
	Name        string `gorm:"type:string" json:"name" validate:"required,string"`
	Description string `gorm:"type:string" json:"description" validate:"required,string"`
}
