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
	"os"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	// import sqlite for cotntext only
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/zhijiewang/Inventory/common"
	"github.com/zhijiewang/Inventory/helper"
)

var db *gorm.DB
var err error

func before() *gorm.DB {
	os.Remove("/tmp/test.db")
	db, err = gorm.Open("sqlite3", "/tmp/test.db")
	if err != nil {
		panic(fmt.Sprintf("failed to connect database %+v", err))
	}
	db.AutoMigrate(common.Product{}, common.Item{})
	db.LogMode(true)
	return db
}
func TestAddItem(t *testing.T) {
	db = before()
	defer after(db)
	var count int
	db.Model(&common.Item{}).Count(&count)
	if count != 0 {
		t.Errorf("Database Not clean. Message %+v", err)
	}
	err := common.AddItem(db, "A", common.ItemStatus(0))
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

	common.AddItems(db, "A", 10)
	db.Model(&common.Item{}).Count(&count)
	if count != 10 {
		t.Error(fmt.Sprintf("Insertin of 10 items failed %+v", db.Error))
	}
	after(db)
}
func TestShipItemWhenThereIsNoItem(t *testing.T) {
	db = before()
	helper.Assert(t, common.ShipItem(db, "A", time.Time{}) != nil, "Failed to raise Error when shipping item but no item availanle in DB")
	after(db)
}
func TestShipItemSuccess(t *testing.T) {
	db = before()
	defer after(db)
	common.AddItems(db, "A", 1)
	err := common.ShipItem(db, "A", time.Time{})
	if err != nil {
		t.Errorf("Operation failed with %+v", err)
	}
	i := &common.Item{}
	db.First(i)
	if i.Status != common.Shipped {
		t.Errorf("Failed to ship item, operation success but database state not updated")
	}
}
func TestBulkShipItemSuccess(t *testing.T) {
	db = before()
	defer after(db)
	common.AddItems(db, "A", 10)
	err := common.ShipItems(db, "A", 10, time.Now())
	if err != nil {
		t.Errorf("Error during operation %+v", err)
	}
	var a []common.Item
	var i int
	db.Where(&common.Item{Status: common.Shipped}).Find(&a)
	if db.Error != nil {
		t.Errorf("Operation error %+v", err)
	}
	db.Model(&common.Item{}).Where(&common.Item{ProductCode: "A", Status: common.Shipped}).Count(&i)
	if db.Error != nil {
		t.Errorf("Operation error %+v", err)
	}
	if i != 10 {
		t.Errorf("Operation success, but the amount of shipped in db does not match")
	}
}
func after(db *gorm.DB) {
	defer os.Remove("/tmp/test.db")
	db.Close()
}
