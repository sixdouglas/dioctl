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

// offCmd represents the off command
var offCmd = &cobra.Command{
	Use:   "off [id]",
	Short: "Switch off a device",
	Long: `Switch off a device specified by its id.`,
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
			if !key.On {
				log.Fatalf("The device with id %d is already off.", elementId)
			} else {
				idio := gpio.DioObj{}
				irpio := gpio.RpioObj{}
				pin := rpio.Pin(configuration.EmitterPinId)
				idio.SendCommand(irpio, pin, elementId, 0, false)
				key.On = false
				viper.Set("elements", configuration.Elements)
				err = viper.WriteConfig()
				if err != nil {
					log.Fatalf("unable to save changes, %v", err)
				}
				fmt.Printf("OFF called for device: %d", elementId)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(offCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// offCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// offCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
