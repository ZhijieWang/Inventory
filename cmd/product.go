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
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zhijiewang/Inventory/common"
)

// productCmd represents the addProduct command
var productCmd = &cobra.Command{
	Use:   "product",
	Short: "A brief description of your command",
	Long:  `A longer description that spans multiple lines and likely contains examples.`,
}
var addProductCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long:  `A longer description `,
	Run: func(cmd *cobra.Command, args []string) {
		db = common.OpenInventory("")
		defer db.Close()
		fmt.Println(productName, productCode)
		db.AddProduct(productName, productCode)
	},
}

var listProductCmd = &cobra.Command{
	Use:   "list",
	Short: "A breif description of your command",
	Long:  `Longer`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("List product called")
		db = common.OpenInventory("")
		defer db.Close()
		res := db.ListProduct()
		for _, i := range *res {
			fmt.Printf("%+v\n", i)
		}
	},
}

var productCode string
var productName string

func init() {
	rootCmd.AddCommand(productCmd)
	addProductCmd.Flags().StringVar(&productName, "p", "", "Name of the product, Human Friendly version")
	addProductCmd.Flags().StringVar(&productCode, "c", "", "quick reference code for the product, machine friendly or easy to type")
	addProductCmd.MarkFlagRequired("p")
	addProductCmd.MarkFlagRequired("c")
	productCmd.AddCommand(listProductCmd)
	productCmd.AddCommand(addProductCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// productCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// productCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
