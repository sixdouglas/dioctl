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
	"github.com/spf13/viper"
	"github.com/stianeikeland/go-rpio"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

// onCmd represents the on command
var onCmd = &cobra.Command{
	Use:   "on [id]",
	Short: "Switch on a device",
	Long: `Switch off the specified device.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var elementId uint64
		var err error

		elementId, err = strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			log.Fatalf("unable to decode device id, %v", err)
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
			log.Fatalf("No device with id %d defined. Are you sure you have added it?", elementId)
		} else {
			key := configuration.Elements[elementPos]
			if key.On {
				log.Fatalf("The device with id %d is already on.", elementId)
			} else {
				idio := gpio.Dio{}
				pin := rpio.Pin(configuration.EmitterPinId)
				idio.SendCommand(pin, elementId, 0, true)
				key.On = true
				viper.Set("elements", configuration.Elements)
				err = viper.WriteConfig()
				if err != nil {
					log.Fatalf("unable to save changes, %v", err)
				}
				fmt.Printf("ON called for device: %d", elementId)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(onCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// onCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// onCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
