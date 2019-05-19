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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"strconv"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [id] [state]",
	Short: "Add a device giving its id and current state",
	Long: `Add a device giving its id and its current state.`,
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

		var elementId uint64
		var elementState bool
		var err error
		elementId, err = strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			log.Fatalf("unable to decode device id, %v", err)
		}
		elementState, err = strconv.ParseBool(args[1])
		if err != nil {
			log.Fatalf("unable to decode device state, %v", err)
		}

		var configuration config.Configuration
		err = viper.Unmarshal(&configuration)
		if err != nil {
			log.Fatalf("unable to decode into struct, %v", err)
		}
		var elementPos = -1
		for i := 0 ; i < len(configuration.Elements); i++ {
			key := configuration.Elements[i]
			if key.Id == elementId {
				elementPos = i
				break
			}
		}

		if elementPos == -1 {
			newElement := config.Element{elementId, elementState}
			configuration.Elements = append(configuration.Elements, newElement)
			viper.Set("elements", configuration.Elements)
			err = viper.WriteConfig()
			if err != nil {
				log.Fatalf("unable to save changes, %v", err)
			}
			fmt.Printf("device id: %d has been added", elementId)
		} else {
			fmt.Printf("device with id: %d already exists in the current configuration", elementId)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
