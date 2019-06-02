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
	"time"
)

// analyseCmd represents the on command
var analyseCmd = &cobra.Command{
	Use:   "analyse [timeout as Duration]",
	Short: "Listen on the air to help finding the signals timings.",
	Long: `Listen on the air to help finding the signals timings. 
The timout should follow the Duration parsing function: https://golang.org/pkg/time/#ParseDuration`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var timeout time.Duration
		var err error

		timeout, err = time.ParseDuration(args[0])
		if err != nil {
			log.Fatalf("unable to decode the timeout, %v", err)
		}

		var configuration config.Configuration
		err = viper.Unmarshal(&configuration)
		if err != nil {
			log.Fatalf("unable to decode into struct, %v", err)
		}

		idio := gpio.Dio{}
		pin := rpio.Pin(configuration.EmitterPinId)
		err = idio.Analyse(pin, timeout)
		if err != nil {
			log.Fatalf("error reading code from pin %d, %v", configuration.ReceiverPinId, err)
		}
		fmt.Println("Analyse finished.")
	},
}

func init() {
	rootCmd.AddCommand(analyseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// analyseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// analyseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
