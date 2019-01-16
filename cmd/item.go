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

	"github.com/spf13/cobra"
	"github.com/zhijiewang/Inventory/common"
)

// itemCmd represents the item command
var itemCmd = &cobra.Command{
	Use:   "item",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("item called")
	},
}
var itemStatus int
func init() {

	rootCmd.AddCommand(itemCmd)
	itemCmd.AddCommand(addItemCmd)
	itemCmd.AddCommand(listItemCmd)
	addItemCmd.Flags().StringVar(&productCode, "p", "", "serial/sku number of the product")
	addItemCmd.Flags().IntVar(&itemStatus, "s", 0, `status code of the product. 
		\t\t\t\t\t\t\t Code 0 means the item is still availabeli\n
		\t\t\t\t\t\t\t Code 1 means the item is reserved for an order \n
		\t\t\t\t\t\t\t Code 2 means the item is already shipped \n
		\t\t\t\t\t\t\t Code 3 means the item is lost or damaged, could be written off.\n`)
	listItemCmd.Flags().IntVar(&itemStatus, "s", 0, `status code of the product. 
		\t\t\t\t\t\t\t Code 0 means the item is still availabeli\n
		\t\t\t\t\t\t\t Code 1 means the item is reserved for an order \n
		\t\t\t\t\t\t\t Code 2 means the item is already shipped \n
		`)
	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// itemCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// itemCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

var addItemCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief",
	Long:  "A long descrition",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Add item called")
		db = before()
		defer db.Close()
		db.Create(&common.Item{})
	},
}
var listItemCmd = &cobra.Command{
	Use:   "list",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("List item called")
		db = before()
		defer db.Close()
	},
}
