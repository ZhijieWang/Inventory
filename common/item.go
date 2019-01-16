package common

import (
	"database/sql/driver"
	"time"

	"github.com/jinzhu/gorm"
	// import blank for context
)

// ItemStatus denotes the current status of the individual item. Default to 0, denoting the item is available.
type ItemStatus int

// Available means the item is in inventory, can be sold
// Reserved means the item is marked for an order. If the order is canceld, the item becomes available again. If the item is shipped, the item becomes shipped.
// Shipped means the item is shipped and the order is fullfiled. This marks the item can now be consider as part of cost of goods sold.
// Lost means the item is either damaged or no longer possible for sale. This can be added to either cost of goods sold or to expense for write off.
const (
	Available ItemStatus = 0
	Reserved             = 1
	Shipped              = 2
	Lost                 = 3
)

// Item represent each individual goods to be sold, assume discrete and concrete item.
type Item struct {
	gorm.Model
	Status      ItemStatus
	ProductCode string
	ShippedDate time.Time
}

func (u *ItemStatus) Scan(value interface{}) error { *u = ItemStatus(value.(int)); return nil }
func (u ItemStatus) Value() (driver.Value, error)  { return int(u), nil }

// AddItem function performs the operation adding the Item instance to the database. If there is any issue, the operation will throw an error. The argument db points to a db connection, the argument pCode points to a productCode, should be randomly generated, either UUID or user defined, the status argument is the enumeration represent item status.
func AddItem(db *gorm.DB, pCode string, status ItemStatus) error {
	return db.Create(&Item{Status: Available, ProductCode: pCode}).Error
}

// AddItems add many items to that of same product code to the database, the serial number is generated automatically. This operation does not check whether the item is already in the database or not. This operation implies bulk intake of product, thus all item status will set to default Availabel
func AddItems(db *gorm.DB, pCode string, num int) error {
	for i := 0; i < num; i += 1 {
		db.Create(&Item{Status: Available, ProductCode: pCode})
	}
	return db.Error
}
