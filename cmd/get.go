// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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
	"github.com/sixdouglas/dioctl/config"
	"github.com/spf13/viper"
	"log"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a list of all saved devices",
	Long: `List all the devices memorized devices.`,
	Run: func(cmd *cobra.Command, args []string) {

		var configuration config.Configuration
		err := viper.Unmarshal(&configuration)
		if err != nil {
			log.Fatalf("unable to decode into struct, %v", err)
		}
		fmt.Println("Here is a list of all elements currently defined and their status")
		fmt.Println("")
		fmt.Println("  +----------+-------+")
		fmt.Println("  |   ID     | STATE |")
		fmt.Println("  +----------+-------+")
		for i := 0 ; i < len(configuration.Elements); i++ {
			key := configuration.Elements[i]
			fmt.Printf("  | %d | %-5t |\n", key.Id, key.On)
		}
		fmt.Println("  +----------+-------+")
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
