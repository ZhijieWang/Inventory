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

// Available means the item is in inventory, can be sold Reserved means the item
// is marked for an order. If the order is canceld, the item becomes available
// again. If the item is shipped, the item becomes shipped.  Shipped means the
// item is shipped and the order is fullfiled. This marks the item can now be
// consider as part of cost of goods sold.  Lost means the item is either
// damaged or no longer possible for sale. This can be added to either cost of
// goods sold or to expense for write off.
const (
	Available ItemStatus = 1
	Reserved             = 2
	Shipped              = 3
	Lost                 = 4
)

// Item represent each individual goods to be sold, assume discrete and concrete
// item.
type Item struct {
	gorm.Model
	Status      ItemStatus `gorm:"not null;default:0"`
	ProductCode string     `gorm:"not null"`
	UnitCost    float32
	ShippedDate time.Time
}

func (u *ItemStatus) Scan(value interface{}) error {
	*u = ItemStatus(value.(int64))
	return nil
}
func (u ItemStatus) Value() (driver.Value, error) {
	return driver.Value(int64(u)), nil
}

// AddItem function performs the operation adding the Item instance to the
// database. If there is any issue, the operation will throw an error. The
// argument db points to a db connection, the argument pCode points to a
// productCode, should be randomly generated, either UUID or user defined, the
// status argument is the enumeration represent item status.
func (db *Inventory) AddItem(pCode string, cost float32, status ItemStatus) error {
	return db.Create(&Item{Status: status, ProductCode: pCode, UnitCost: cost}).Error
}

// AddItems add many items to that of same product code to the database, the
// serial number is generated automatically. This operation does not check
// whether the item is already in the database or not. This operation implies
// bulk intake of product, thus all item status will set to default Availabel
func (db *Inventory) AddItems(pCode string, cost float32, num int) error {
	for i := 0; i < num; i += 1 {
		db.Create(&Item{Status: Available, ProductCode: pCode, UnitCost: cost})
	}
	return db.Error
}

// ShipItem mark item as shipped and shipped date as given date
func (db *Inventory) ShipItem(pCode string, date time.Time) error {
	if (date == time.Time{}) {
		date = time.Now()
	}
	var i []Item
	db.findAvailable(pCode, 1, &i)
	if (len(i) == 0) || (i[0].Status != Available) {
		return error(fmt.Errorf("Inventory Empty, No Available Item"))
	}
	db.Model(i[0]).Updates(Item{Status: Shipped, ShippedDate: date})
	return db.Error
}

//findAvailable performs query to read available items from the inventory.
//given Gorm does not throw error when performing certain invalid query, the
//result needs to be campared to nil
func (db *Inventory) findAvailable(pCode string, number int, i *[]Item) {
	db.Where(&Item{ProductCode: pCode, Status: Available}).Limit(number).Find(i)
}

// ShipItems ships multiple items, given a ProductCode and the numbers. The
// default date time is Now.
func (db *Inventory) ShipItems(pCode string, number int, date time.Time) error {
	if (date == time.Time{}) {
		date = time.Now()
	}
	var i []Item
	db.findAvailable(pCode, number, &i)
	if len(i) != number {
		return error(fmt.Errorf("Not enough item to fulfill the requeste. Requested %d, available %d", number, len(i)))
	}
	db.Model(&i).Update(Item{Status: Shipped, ShippedDate: date})
	return db.Error
}

// CostOfGoodsSold aggregate the cost of goods shipped or lost(due to return and
// other).  the parameter month indicates the first day of the month of
// interest.

func (db *Inventory) CostOfGoodsSold(month time.Time) int {
	end := month.AddDate(0, 1, 0)
	type Result struct {
		Total int
	}
	var result Result
	db.Model(&Item{}).Select("sum(unit_cost )as total").Where("status = ? ", Shipped).Where("shipped_date BETWEEN ? AND ?", month, end).Scan(&result)
	return int(result.Total)
}
