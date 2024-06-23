package entity

import (
	"time"

	"github.com/google/uuid"
)

type Entity struct {
	ID         uuid.UUID `gorm:"primarykey;type:char(36);not null"`
	CreatedBy  uuid.UUID `gorm:"not null"`
	CreatedAt  time.Time `gorm:"not null"`
	ModifiedBy *uuid.UUID
	ModifiedAt *time.Time
	IsActive   bool `gorm:"default:true;not null"`
}
