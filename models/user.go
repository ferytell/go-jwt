package models

import (
	"fmt"

	"github.com/ferytell/go-jwt/helpers"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var validate *validator.Validate = validator.New()

// gorm:"type:int;primary_key
type User struct {
	GormModel
	UserName    string        `gorm:"not null; uniqueIndex" json:"user_name" form:"user_name" valid:"required~Username is required"`
	Email       string        `gorm:"not null; uniqueIndex" json:"email" form:"email" valid:"required~Your email is required,email~invailid email format"`
	Password    string        `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Age         int           `gorm:"not null" json:"age" form:"age" validate:"required,numeric,min=18"`
	Comments    []Comments    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comments"`
	SocialMedia []SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"social_media"`
	Photo       []Photo       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"photo"`
	Products    []Product     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"products"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if err := validate.Struct(u); err != nil {
		fmt.Println("Validation error:", err)
		return err
	}

	u.Password = helpers.HashPass(u.Password)
	return nil
}

func (p *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if err := validate.Struct(p); err != nil {
		fmt.Println("Validation error:", err)
		return err
	}

	return nil
}
