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
	"testing"

	"github.com/zhijiewang/Inventory/cmd"
)

func TestReplSearch(t *testing.T) {
	db = before()
	defer after(db)
	if len(*db.ListProduct()) != 0 {
		t.Errorf("Database Not clean, there should be no product in the testing db at the begining.")
	}
	root := cmd.GetRootCmd()
	a, b, _ := root.Find([]string{"product", "add", "--p", "name", "--c", "1000"})
	a.ParseFlags(b)
	a.Run(a, a.Flags().Args())
	if len(*db.ListProduct()) != 1 {
		t.Errorf("Add Command Failed. Either the command failed to insert into db or the parser failed to call AddProductCommand")
	}
}
