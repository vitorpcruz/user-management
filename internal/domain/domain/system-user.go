package domain

import "github.com/google/uuid"

type SystemUser struct {
	Entity
	SystemID uuid.UUID `gorm:"type:char(36);not null"`
	UserID   uuid.UUID `gorm:"type:char(36);not null"`
	System   System    `gorm:"foreignKey:SystemID;references:ID"`
	User     User      `gorm:"foreignKey:UserID;references:ID"`
}
