package models

type Status string

const (
	ACTIVE   Status = "ACTIVE"
	INCATIVE Status = "INCATIVE"
)

type Feature struct {
	BaseModel
	Name   string `gorm:"type:string" json:"name" validate:"required,string"`
	Status Status
}
