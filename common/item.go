// Copyright © 2019 zhijie wang
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
//
package common

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
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
	Status      ItemStatus `gorm:"not null; default:0"`
	ProductCode string     `gorm: "not null"`
	ShippedDate time.Time
}

func (u *ItemStatus) Scan(value interface{}) error {
	*u = ItemStatus(value.(int64))
	return nil
}
func (u ItemStatus) Value() (driver.Value, error) {
	return driver.Value(int64(u)), nil
}

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

// ShipItem mark item as shipped and shipped date as given date
func ShipItem(db *gorm.DB, pCode string, date time.Time) error {
	if (date == time.Time{}) {
		date = time.Now()
	}
	var i *[]Item
	i = findAvailable(db, pCode, 1)
	if len(*i) == 0 {
		return error(fmt.Errorf("Inventory Empty, No Available Item"))
	}
	db.Model(i).Updates(Item{Status: Shipped, ShippedDate: date})
	return db.Error
}

//findAvailable performs query to read available items from the inventory.
// given Gorm does not throw error when performing certain invalid query,
// the result needs to be campared to nil
func findAvailable(db *gorm.DB, pCode string, number int) *[]Item {
	var i []Item
	db.Where(&Item{ProductCode: pCode, Status: Available}).Limit(number).Find(&i)
	return &i
}
func ShipItems(db *gorm.DB, pCode string, number int, date time.Time) error {
	if (date == time.Time{}) {
		date = time.Now()
	}
	i := *findAvailable(db, pCode, number)
	if len(i) != number {
		return error(fmt.Errorf("Not enough item to fulfill the requeste. Requested %d, available %d", number, len(i)))
	}
	db.Model(&i).Updates(Item{Status: Shipped, ShippedDate: date})
	return db.Error
}
