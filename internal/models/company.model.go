package models

type Company struct {
	BaseModel
	Name        string `gorm:"type:string" json:"name" validate:"required,string"`
	Description string `gorm:"type:string" json:"description" validate:"required,string"`
	// User        []User `gorm:"foreignKey:Company;references:id"`
}
