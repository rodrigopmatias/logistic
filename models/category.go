package models

type Category struct {
	Base
	Title   string `json:"title"`
	OwnerID string `json:"ownerID"`
	Owner   User   `gorm:"references:ID"`
}
