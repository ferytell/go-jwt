package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comments struct {
	GormModel
	Message string `gorm:"not null" json:"message" form:"message" valid:"required~Your full name is required"`
	PhotoID uint
	Photo   *Photo
	UserID  uint
	User    *User
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
