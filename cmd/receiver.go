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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"strconv"
)

// setCmd represents the set command
var setReceiverCmd = &cobra.Command{
	Use:   "receiver [pinId]",
	Short: "Set the receiver pin ID",
	Long: `Set the receiver pin ID.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		var receiverId uint64
		var err error
		receiverId, err = strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			log.Fatalf("unable to decode receiver id, %v", err)
		}
		fmt.Printf("Receiver pin id is %d\n", viper.GetInt("receiverpinid"))
		viper.Set("ReceiverPinId", receiverId)
		err = viper.WriteConfig()
		if err != nil {
			log.Fatalf("unable to save changes, %v", err)
		}
		fmt.Printf("Receiver pin id set to: %d", receiverId)
	},
}

func init() {
	setCmd.AddCommand(setReceiverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
