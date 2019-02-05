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
// limitations under the License.
package common_test

import (
	"fmt"
	"testing"
	"time"

	// import sqlite for cotntext only
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/zhijiewang/Inventory/common"
	"github.com/zhijiewang/Inventory/helper"
)

func TestAddItem(t *testing.T) {
	db = before()
	defer after(db)
	var count int
	db.Model(&common.Item{}).Count(&count)
	if count != 0 {
		t.Errorf("Database Not clean. Message %+v", err)
	}
	err := db.AddItem("A", int64(0), common.ItemStatus(0))
	if err != nil {
		t.Errorf("Insertion error %+v", err)
	}
	db.Model(&common.Item{}).Count(&count)
	if count != 1 {
		t.Error("Insertion resulted in multiple record")
	}
}
func TestAddItems(t *testing.T) {
	db = before()
	var count int
	db.Model(&common.Item{}).Count(&count)
	if count != 0 {
		t.Errorf("Database not clean. message %+v", db.Error)
	}

	db.AddItems("A", int64(5), 10)
	db.Model(&common.Item{}).Count(&count)
	if count != 10 {
		t.Error(fmt.Sprintf("Insertin of 10 items failed %+v", db.Error))
	}
	after(db)
}
func TestShipItemWhenThereIsNoItem(t *testing.T) {
	db = before()
	helper.Assert(t, db.ShipItem("A", time.Time{}) != nil, "Failed to raise Error when shipping item but no item availanle in DB")
	after(db)
}
func TestShipItemSuccess(t *testing.T) {
	db = before()
	defer after(db)
	db.AddItems("A", int64(10), 1)
	db.AddItems("A", int64(10), 1)
	err := db.ShipItem("A", time.Time{})
	if err != nil {
		t.Errorf("Operation failed with %+v", err)
	}
	i := &common.Item{}
	db.First(i)
	if i.Status != common.Shipped {
		t.Error("Failed to ship item, operation success but database state not updated")
	}
	var iList []common.Item
	db.Model(&common.Item{}).Where("Status =?", common.Shipped).Find(&iList)
	if len(iList) != 1 {
		t.Errorf("Failed to ship the correct amount of the item. Expected 1, acutally %d", len(iList))
	}
	err = db.ShipItem("A", time.Time{})
	db.Model(&common.Item{}).Where("Status =?", common.Shipped).Find(&iList)
	if len(iList) != 2 {
		t.Errorf("Failed to ship the correcnt amount fo the item. Expected 2 to be shipped after 2 operations, actually %d", len(iList))
	}
}
func TestBulkShipItemSuccess(t *testing.T) {
	db = before()
	defer after(db)
	db.AddItems("A", 5, 10)
	err := db.ShipItems("A", 10, time.Now())
	var i int
	if err != nil {
		t.Errorf("Error during operation %+v", err)
	}
	db.Model(&common.Item{}).Where(&common.Item{ProductCode: "A", Status: common.Shipped}).Count(&i)
	if db.Error != nil {
		t.Errorf("Operation error %+v", err)
	}
	if i != 10 {
		t.Errorf("Operation success, but the amount of shipped in db does not match. Expected 10, actual %d", i)
	}
}
func TestCostOfGoodsCalculation(t *testing.T) {
	db = before()
	defer after(db)
	err := db.AddItems("A", 5, 10)

	if err != nil {
		t.Errorf("Error during operation %+v", err)
	}
	err = db.ShipItems("A", 10, time.Now())
	var r []common.Item

	db.Model(&common.Item{}).Find(&r)
	//t.Logf("%+v \n", r)
	result := db.CostOfGoodsSold(time.Now().AddDate(0, 0, -2))
	if result != 50 {
		t.Errorf("Expected Cost of Goods Sold for 10 units of $5 is $50, actually is %v", result)
	}
}
