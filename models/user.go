package models

import (
	"github.com/asaskevich/govalidator"
	"github.com/ferytell/go-jwt/helpers"
	"gorm.io/gorm"
)

// gorm:"type:int;primary_key
type User struct {
	GormModel
	UserName    string        `gorm:"not null" json:"user_name" form:"user_name" valid:"required~Username is required"`
	Email       string        `gorm:"not null; uniqueIndex" json:"email" form:"email" valid:"required~Your email is required,email~invailid email format"`
	Password    string        `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Age         int           `gorm:"not null" json:"age" form:"age" validate:"required,numeric,min=18"`
	Comments    []Comments    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comments"`
	SocialMedia []SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"social_media"`
	Photo       []Photo       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"photo"`
	Products    []Product     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"products"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {

	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}

func (p *User) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
