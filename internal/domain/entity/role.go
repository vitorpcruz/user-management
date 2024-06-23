package entity

import "github.com/google/uuid"

type Role struct {
	Entity
	ID              uuid.UUID `gorm:"type:char(36);not null"`
	Code            string
	Description     string
	SystemUserRoles []SystemUserRole
}
