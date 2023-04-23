package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// Photo represents the model for a Photo
type Photo struct {
	GormModel
	Title    string     `gorm:"not null" json:"title" form:"title" valid:"required~Title is required"`
	Caption  string     `gorm:"not null" json:"caption" form:"caption" valid:"required~Caption is required"`
	PhotoURL string     `gorm:"not null" json:"photo_url" form:"photo_url" valid:"required~Photo URL is required,url~invalid URL format"`
	UserID   uint       `gorm:"index"`
	User     *User      `json:"user,omitempty"`
	Comments []Comments `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comments"`
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)
	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}

func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)
	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}

func (p *Photo) AfterDelete(tx *gorm.DB) (err error) {
	// delete all comments associated with the photo
	if err := tx.Delete(&[]Comments{}, "photo_id = ?", p.ID).Error; err != nil {
		return err
	}
	return nil
}
