package schema

import (
	"time"

	"github.com/google/uuid"
)

type BaseEntityResponse struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
