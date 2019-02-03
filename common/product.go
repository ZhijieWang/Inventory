package common

import (
	"github.com/jinzhu/gorm"
)

// Product is the interface defines one type of product or sku
type Product struct {
	gorm.Model
	Code     string `gorm:"unique;not null"`
	NickName string `gorm:"unique;not null"`
	Items    []Item `gorm:"foreignkey:ProductCode"`
}

func (db *Inventory) ListProduct() *[]Product {
	res := []Product{}
	db.Find(&res)
	return &res
}

func (db *Inventory) AddProduct(pName string, code string) error {
	db.Create(&Product{NickName: pName, Code: code})
	return db.Error

}
func (db *Inventory) ListInventory(pCode string) *[]Item {
	p := Product{Code: pCode}
	var i []Item
	db.Find(&p)
	db.Model(&p).Related(&i, "items")
	return &i
}
