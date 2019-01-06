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
	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"
)

// addProductCmd represents the addProduct command
var addProductCmd = &cobra.Command{
	Use:   "addProduct",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("addProduct called")
		db, err := gorm.Open("sqlite3", "test.db")
		if err != nil {
			panic("failed to connect database")
		}
		defer db.Close()
		db.Create(&Product{ProductName: productName, Code: code})

	},
}

var productName string
var code string

func init() {
	rootCmd.AddCommand(addProductCmd)
	addProductCmd.Flags().StringvVar(&productName, "p", nil, "Name of the product, Human Friendly version")
	addProductCmd.Flags().StringVar(&code, "c", nil, "quick reference code for the product, machine friendly or easy to type")
	addProductCmd.MarkFlagRequired("p")
	addProductCmd.MarkfFlagRequired("c")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addProductCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addProductCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
