package domain

import "github.com/google/uuid"

type System struct {
	Entity
	CustomerID  uuid.UUID `gorm:"type:char(36);not null"`
	Customer    Customer  `gorm:"foreignKey:CustomerID;references:ID"`
	SystemUsers []SystemUser

	Name string `gorm:"not null:true"`
}
