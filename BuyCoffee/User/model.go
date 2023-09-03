package User

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	UserName           string
	PassWord           string
	EmailAddress       string
	FirstName          string
	LastName           string
	Description        string
	ProfileAvatarID    string
	BackGroundAvatarID string
}
