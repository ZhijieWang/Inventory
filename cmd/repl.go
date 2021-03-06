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
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// replCmd represents the repl command
var replCmd = &cobra.Command{
	Use:   "repl",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("repl called")
		end := false

		var err error
		reader := bufio.NewReader(os.Stdin)
		var a string

		for end != true {
			fmt.Print(">>")
			a, _ = reader.ReadString('\n')
			if a != "\n" {
				end, err = parse(strings.Fields(a))
			}
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println()
		}
		if end {
			return
		}
	},
}

func parse(input []string) (bool, error) {
	if input[0] == "exit" {
		return true, nil
	}
	root := GetRootCmd()
	command, flags, err := root.Find(input)
	if err != nil {
		return false, err
	}
	command.ParseFlags(flags)
	command.Run(command, command.Flags().Args())

	return false, nil

}
func init() {
	rootCmd.AddCommand(replCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// replCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// replCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
