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
	"github.com/sixdouglas/dioctl/gpio"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stianeikeland/go-rpio"
	"log"
)

// readCmd represents the on command
var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Decode sender id.",
	Long: `Decode sender id.`,
	Run: func(cmd *cobra.Command, args []string) {
		var code uint64
		var err error

		var configuration config.Configuration
		err = viper.Unmarshal(&configuration)
		if err != nil {
			log.Fatalf("unable to decode into struct, %v", err)
		}

		idio := gpio.Dio{}
		pin := rpio.Pin(configuration.EmitterPinId)
		code, err = idio.ReadCode(pin)
		if err != nil {
			log.Fatalf("error reading code from pin %d, %v", configuration.ReceiverPinId, err)
		}
		fmt.Printf("Sender is %d.", code)
	},
}

func init() {
	rootCmd.AddCommand(readCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// readCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// readCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
