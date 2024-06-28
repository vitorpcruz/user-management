package entity

import "github.com/google/uuid"

type System struct {
	Entity
	CustomerID uuid.UUID `gorm:"type:char(36);not null"`
	Customer   Customer  `gorm:"foreignKey:CustomerID;references:ID"`
	Name       string    `gorm:"not null:true"`
	Users      []User    `gorm:"many2many:systems_users"`
}
