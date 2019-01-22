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

func ListProuct(db *gorm.DB) *[]Product {
	res := []Product{}
	db.Find(&res)
	return &res
}

func AddProduct(db *gorm.DB, pName string, code string) error {
	db.Create(&Product{NickName: pName, Code: code})
	return db.Error

}
