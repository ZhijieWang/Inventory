package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/sqlite"
)

type Product struct {
	gorm.Model
	Code     string
	NickName string
}
