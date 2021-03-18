package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// BaseModel is a base type for all entities, containing basic fields
type BaseModel struct {
	ID        uuid.UUID  `gorm:"type:uuid;primary_key" json:"id"`
	CreatedBy uuid.UUID  `gorm:"type:uuid" json:"createdBy"`
	UpdatedBy uuid.UUID  `gorm:"type:uuid" json:"updatedBy"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `sql:"index"`
}

// BeforeCreate - GORM hook
func (base *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	base.ID = uuid.New()

	return nil
}
