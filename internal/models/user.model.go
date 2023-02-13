package models

type User struct {
	BaseModel
	Name     string `gorm:"type:string" json:"name" validate:"required,string"`
	Position string `gorm:"type:string" json:"position" validate:"required,string"`
	Email    string `gorm:"type:string" json:"email" validate:"required,string"`
	Password string `gorm:"type:string" json:"password" validate:"required,string"`
	// Company  uuid.UUID
	// Role   uint
	// Device []Device `gorm:"foreignKey:User;references:id"`
}
