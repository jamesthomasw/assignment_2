package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID           int     	`json:"id" gorm:"primaryKey"`
	CustomerName string  	`json:"customer_name"`
	OrderedAt    time.Time  `json:"ordered_at"`
	ItemRefer    int		`json:"item_id"`
	Items        Item  		`gorm:"foreignKey:ItemRefer"`
}

