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
	"testing"

	"github.com/zhijiewang/Inventory/common"
)

func TestProductUniqueness(t *testing.T) {
	db = before()
	defer after(db)
	var count int
	count = len(*db.ListProduct())
	if count != 0 {
		t.Log(count)
		t.Errorf("Database Not clean. Message: %+v", err)

	}
	db.AddProduct("productName", "code")
	if len(db.GetErrors()) != 0 {
		for e := range db.GetErrors() {
			t.Errorf("DB Error %+v", e)
		}
	}
	count = len(*db.ListProduct())
	if count != 1 {
		t.Errorf("Failed to insert query 1. Message: %+v", err)
	}
	db.AddProduct("productName", "code")
	count = len(*db.ListProduct())
	if count != 1 {
		t.Errorf("Failed to to comply with uniqueness. Message: %+v", err)
	}
	after(db)
}
func TestProductList(t *testing.T) {
	db = before()
	db.AddProduct("Product A", "00000")
	db.AddProduct("Product B", "00001")
	res := db.ListProduct()
	var match bool = true
	match = match && ((*res)[0].Code == "00000")
	match = match && ((*res)[0].NickName == "Product A")

	match = match && ((*res)[1].Code == "00001")
	match = match && ((*res)[1].NickName == "Product B")
	if !match {
		t.Errorf("Expected Product List is [Product A 00000, Product B 00001], actual product list is %+v", res)
	}
}

func TestProductInventoryAssociation(t *testing.T) {
	db = before()
	var a, b common.Product
	var h, j, k common.Item
	a = common.Product{
		NickName: "A",
		Code:     "10",
	}
	b = common.Product{
		NickName: "B",
		Code:     "11",
	}
	db.Save(&a)
	db.Save(&b)

	h = common.Item{
		Status:      common.Available,
		ProductCode: "10",
		UnitCost:    10.0,
	}
	j = common.Item{
		Status:      common.Available,
		ProductCode: "10",
		UnitCost:    10.0,
	}
	k = common.Item{
		Status:      common.Available,
		ProductCode: "10",
		UnitCost:    10.,
	}
	db.Save(&h)
	db.Save(&j)
	db.Save(&k)
	var p []common.Product

	db.Find(&p)
	db.Preload("Items").Find(&p)
	if len(p[0].Items) != 3 {
		t.Errorf("Expecting the association rule to be able to find all associated items.\n Expecting 3 items: \n %+v,\n %+v,\n %+v,\n  Found %d, which are %+v", h, j, k, len(p[0].Items), p[0].Items)

	}
	var i []common.Item
	var s = common.Product{Code: "10"}
	db.Find(&s, common.Product{Code: "10"})
	db.Model(&s).Related(&i, "items")
	if len(i) != 3 {
		t.Errorf("Failed")
	}
}
