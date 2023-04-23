package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// SocialMedia represents the model for a SocialMedia
type SocialMedia struct {
	GormModel
	Name           string `gorm:"not null" json:"name" form:"name" valid:"required~Name is required"`
	SocialMediaURL string `gorm:"not null; uniqueIndex" json:"social_media_url" form:"social_media_url" valid:"required~Social media URL is required,url~Invalid URL format"`
	UserID         uint   `gorm:"index"`
	User           *User  `json:"user,omitempty"`
}

func (u *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {

	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	//u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}

func (p *SocialMedia) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
