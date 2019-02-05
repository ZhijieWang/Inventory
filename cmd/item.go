// Copyright © 2019 Zhijie (Bill) Wang <wangzhijiebill@gmail.com>
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

package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"github.com/zhijiewang/Inventory/common"
)

// itemCmd represents the item command
var itemCmd = &cobra.Command{
	Use:   "item",
	Short: "This is the command interface for managing items in inventory",
	Long:  `A command to manage items in inventory. There are 3 sub commands: Add, Ship, List. See detailed information for each subcommand.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("item called")
	},
}
var itemStatus int
var shipDate string

func init() {

	itemCmd.PersistentFlags().StringVarP(&productCode, "product", "p", "", "serial/sku number of the product")
	addItemCmd.Flags().IntVar(&itemStatus, "s", 0, `status code of the product. 
		Code 0 means the item is still available
		Code 1 means the item is reserved for an order 
		Code 2 means the item is already shipped 
		Code 3 means the item is lost or damaged, could be written off.`)

	listItemCmd.Flags().IntVar(&itemStatus, "s", 2, `status code of the product. 
		Code 0 means the item is still available
		Code 1 means the item is reserved for an order
		Code 2 means the item is already shipped
		`)
	addItemCmd.MarkFlagRequired("p")
	shipItemCmd.Flags().StringVar(&shipDate, "d", "", "date of the item shipped. Month-Date-Year format")
	rootCmd.AddCommand(itemCmd)
	itemCmd.AddCommand(addItemCmd)
	itemCmd.AddCommand(listItemCmd)
	itemCmd.AddCommand(shipItemCmd)

}

var addItemCmd = &cobra.Command{
	Use:   "add",
	Short: "This is the command to add item to the inventory.",
	Long:  "This command is responsible for for adding inventory to the system. You can add single item to the inventory with specific status for inventory reconciliation purpose. Or you can add multiple items of the same product code to the inventory, automatically marked as available for stock up.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Add item called")
		db = common.OpenInventory("")
		if len(args) > 0 {
			i, err := strconv.Atoi(args[0])
			if err != nil {
				err = db.AddItem(productCode, int64(1), common.ItemStatus(itemStatus))
				if err != nil {
					panic(err)
				}
			} else {
				db.AddItems(productCode, int64(1), i)
			}
		}
	},
}
var listItemCmd = &cobra.Command{
	Use:   "list",
	Short: "This is the command to list items currently in inventory.",
	Long:  "This command will list all inventory if no product code was specified. If a product code is specified, this command will list all items related to the specific product code.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("List item called")
		db = common.OpenInventory("")
		res := *db.ListInventory(productCode)
		for _, r := range res {

			fmt.Printf("%+v\n", r)
		}
	},
}
var shipItemCmd = &cobra.Command{
	Use:   "ship",
	Short: "This item marks items as shiped status.",
	Long:  "This item marks item as shipped status, for a given product code. The count number option is optional. If the count is not provided, it will default to 1.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Ship item called")
		db = common.OpenInventory("")
		defer db.Close()
		d, _ := time.Parse("01-02-2006", shipDate)
		db.ShipItem(productCode, d)
	},
}
