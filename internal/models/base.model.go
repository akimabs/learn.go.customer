package models

import (
	"time"

	uuid "github.com/google/uuid"
)

type BaseModel struct {
	ID        uuid.UUID `gorm:"default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
