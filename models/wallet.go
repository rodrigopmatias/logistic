package models

type Wallet struct {
	Base
	Name    string `json:"name" gorm:"not null"`
	OwnerID string `json:"ownerID"`
	Owner   User   `gorm:"references:ID"`
}
