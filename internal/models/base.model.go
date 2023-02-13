package models

import (
	"time"

	uuid "github.com/google/uuid"
)

type BaseModel struct {
	ID        uuid.UUID `gorm:"type:uuid" default:"uuid_generate_v4()"`
	CreatedAt time.Time `gorm:"type:string" json:"createdAt" autoUpdateTime:"milli" default:"current_timestamp"`
	UpdatedAt time.Time `gorm:"type:string" json:"updatedAt" autoUpdateTime:"milli" default:"current_timestamp"`
}
