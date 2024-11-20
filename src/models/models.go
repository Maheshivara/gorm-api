package models

import (
	"time"

	"github.com/google/uuid"
)

type Food struct {
	ID        uuid.UUID  `gorm:"primaryKey; type:uuid; default:gen_random_uuid()" json:"id" example:"4e51bdb9-c75a-4198-8f2d-9695e3ffaa83"`
	Name      string     `json:"name" example:"Macarrão à Milanesa"`
	Price     float32    `json:"price" example:"15.47"`
	CreatedAt time.Time  `gorm:"default:now()" json:"createdAt" example:"2024-11-20T02:18:56.744307Z"`
	UpdatedAt time.Time  `gorm:"default:now()" json:"updatedAt" example:"2024-11-20T02:18:56.744307Z"`
	DeletedAt *time.Time `json:"-"`
}
