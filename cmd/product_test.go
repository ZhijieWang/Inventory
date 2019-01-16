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
package cmd_test

import (
	"os"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/zhijiewang/Inventory/common"
)

var db *gorm.DB
var err error

func before() *gorm.DB {
	os.Remove("/tmp/test.db")
	db, err = gorm.Open("sqlite3", "/tmp/test.db")
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(common.Product{}, common.Item{})
	return db
}
func TestProductUniqueness(t *testing.T) {
	db = before()
	var count int
	db.Model(&common.Product{}).Count(&count)
	if count != 0 {
		t.Log(count)
		t.Errorf("Database Not clean. Message: %+v", err)

	}
	db.Create(&common.Product{NickName: "productName", Code: "code"})
	if len(db.GetErrors()) != 0 {
		for e := range db.GetErrors() {
			t.Errorf("DB Error %+v", e)
		}
	}
	db.Model(&common.Product{}).Count(&count)
	if count != 1 {
		t.Errorf("Failed to insert query 1. Message: %+v", err)
	}
	db.Create(&common.Product{NickName: "productName", Code: "code"})
	if count != 1 {
		t.Errorf("Failed to to comply with uniqueness. Message: %+v", err)
	}
	defer db.Close()
}
func after(db *gorm.DB) {
	// if err != nil {
	// 	panic("Failed to close database")
	// }
	defer os.Remove("/tmp/test.db")
	db.Close()
}
