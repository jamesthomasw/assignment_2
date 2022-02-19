package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	ID          int    `json:"id" gorm:"primaryKey"`
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}
