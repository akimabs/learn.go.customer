package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uuid.UUID `gorm:"type:char(36)" json:"id"`
	CreatedAt time.Time `gorm:"type:string" json:"createdAt" autoUpdateTime:"milli"`
	UpdatedAt time.Time `gorm:"type:string" json:"updatedAt" autoUpdateTime:"milli"`
}

func (base *BaseModel) BeforeCreate(tx *gorm.DB) error {
	uuid := uuid.NewV4().String()
	tx.Statement.SetColumn("id", uuid)

	now := time.Now()
	if base.CreatedAt.IsZero() {
		tx.Statement.SetColumn("createdAt", now)
	}
	if base.UpdatedAt.IsZero() {
		tx.Statement.SetColumn("updatedAt", now)
	}
	return nil
}
