package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Base type for all entities
type Base struct {
	ID        uuid.UUID  `gorm:"type:uuid;primary_key" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `sql:"index"`
}

// BeforeCreate - GORM hook
func (base *Base) BeforeCreate(tx *gorm.DB) (err error) {
	base.ID = uuid.New()

	return nil
}
