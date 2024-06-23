package domain

type User struct {
	Entity
	Firstname   string `gorm:"not null:true"`
	Lastname    string `gorm:"not null:true"`
	Email       string `gorm:"not null:true"`
	Login       string `gorm:"not null:true;index"`
	SystemUsers []SystemUser
}
