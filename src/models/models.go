package models

import (
	"time"

	"github.com/google/uuid"
)

type Food struct {
	ID        uuid.UUID  `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Name      string     `json:"name"`
	Price     float32    `json:"price"`
	CreatedAt time.Time  `gorm:"default:now()" json:"createdAt"`
	UpdatedAt time.Time  `gorm:"default:now()" json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}
