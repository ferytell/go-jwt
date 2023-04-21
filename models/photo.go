package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	GormModel
	Title    string `gorm:"not null" json:"title" form:"title" valid:"required~Title is required"`
	Caption  string `gorm:"not null" json:"caption" form:"caption" valid:"required~Caption is required"`
	PhotoURL string `gorm:"not null" json:"photo_url" form:"photo_url" valid:"required~Photo URL is required,url~invalid URL format"`
	UserID   uint
	User     *User
}

func (u *Photo) BeforeCreate(tx *gorm.DB) (err error) {

	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
