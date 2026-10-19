package domain

import (
	"time"
	"uuid"
)

type DefaultAttrs struct {
	ID        uuid.UUID
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
