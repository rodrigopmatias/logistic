package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID string `json:"id" gorm:"type:UUID;primaryKey"`
}

func (base *Base) BeforeCreate(tx *gorm.DB) error {
	base.ID = uuid.NewV4().String()
	return nil
}
