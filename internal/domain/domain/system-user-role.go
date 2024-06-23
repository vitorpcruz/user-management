package domain

import (
	"github.com/google/uuid"
)

type SystemUserRole struct {
	Entity
	SystemUserID uuid.UUID  `gorm:"type:char(36);not null"`
	RoleID       uuid.UUID  `gorm:"type:char(36);not null"`
	SystemUser   SystemUser `gorm:"foreignKey:SystemUserID;references:ID"`
	Role         Role       `gorm:"foreignKey:RoleID;references:ID"`
}
