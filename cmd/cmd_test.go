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
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/zhijiewang/Inventory/common"
)

var db *common.Inventory
var err error
var conn *gorm.DB

func before() *common.Inventory {
	os.Remove("/tmp/test.db")
	conn, err = gorm.Open("sqlite3", "/tmp/test.db")
	if err != nil {
		panic(fmt.Sprintf("failed to connect database %+v", err))
	}
	db = &common.Inventory{conn}
	db.AutoMigrate(common.Product{}, common.Item{})
	db.LogMode(true)
	return db
}
func after(db *common.Inventory) {
	defer os.Remove("/tmp/test.db")
	db.Close()
}
