package common

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Struct ItemStatus

type Item struct {
	gorm.Model
}
