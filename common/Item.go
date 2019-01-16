package common

import (
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
	SerialID    string
	ProductCode string
}
