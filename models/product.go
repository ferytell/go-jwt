package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// Product represents the model for a product
type Product struct {
	GormModel
	Title       string `json:"title" form:"title" valid:"required~Title of your product is required"`
	Description string `json:"description" form:"description" valid:"required~Description of your product is required"`
	UserID      uint   `gorm:"index"`
	User        *User  `json:"user,omitempty"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (p *Product) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
