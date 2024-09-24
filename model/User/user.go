package User

import "gorm.io/gorm"

// User struct
type Users struct {
	gorm.Model
	Login string `gorm:"uniqueIndex;not null;size:200" json:"login"`
	Senha string `gorm:"not null;size:200" json:"password"`
}
