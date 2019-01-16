package common

import (
	"github.com/jinzhu/gorm"
)

// Product is the interface defines one type of product or sku
type Product struct {
	gorm.Model
	Code     string `gorm:"unique;not null"`
	NickName string `gorm:"unique;not null"`
}
