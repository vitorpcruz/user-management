package entity

type User struct {
	Entity
	Firstname string   `gorm:"not null:true"`
	Lastname  string   `gorm:"not null:true"`
	Email     string   `gorm:"not null:true"`
	Login     string   `gorm:"not null:true;index"`
	Systems   []System `gorm:"many2many:systems_users"`
}
