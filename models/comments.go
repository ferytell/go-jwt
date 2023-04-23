package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// Comments represents the model for a Comments
type Comments struct {
	GormModel
	Message string `gorm:"not null" json:"message" form:"message" valid:"required~Message is required"`
	PhotoID uint   `gorm:"index"`
	Photo   *Photo `json:"photo,omitempty"`
	UserID  uint   `gorm:"index"`
	User    *User  `json:"user,omitempty"`
}

func (u *Comments) BeforeCreate(tx *gorm.DB) (err error) {

	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	//u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}

func (p *Comments) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
