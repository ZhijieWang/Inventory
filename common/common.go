package common

import "github.com/jinzhu/gorm"

type Inventory struct {
	*gorm.DB
}
