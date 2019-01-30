package common

import "github.com/jinzhu/gorm"

type Inventory struct {
	*gorm.DB
}

func OpenInventory(path string) *Inventory {
	var db *gorm.DB
	var err error
	if path == "" {

		db, err = gorm.Open("sqlite3", "/tmp/test.db")
	} else {
		db, err = gorm.Open("sqlite3", path)
	}
	if err != nil {
		panic(err)
	}

	return &Inventory{db}
}
