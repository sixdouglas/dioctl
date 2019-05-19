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
var setPositionCmd = &cobra.Command{
	Use:   "position [latitude] [longitude]",
	Short: "Set the current position on hearth",
	Long: `Set the latitude and longitude on hearth to calculate the Sunrise end Sunset.`,
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

		var latitude, longitude float64
		var err error
		latitude, err = strconv.ParseFloat(args[0], 64)
		if err != nil {
			log.Fatalf("unable to decode latitude, %v", err)
		}
		longitude, err = strconv.ParseFloat(args[1], 64)
		if err != nil {
			log.Fatalf("unable to decode longitude, %v", err)
		}

		fmt.Printf("Current latitude was %f\n", viper.GetFloat64("latitude"))
		fmt.Printf("Current longitude was %f\n", viper.GetFloat64("longitude"))

		viper.Set("latitude", latitude)
		viper.Set("longitude", longitude)

		err = viper.WriteConfig()
		if err != nil {
			log.Fatalf("unable to save changes, %v", err)
		}
		fmt.Printf("Latitude set to: %f\n", latitude)
		fmt.Printf("Longitude set to: %f\n", longitude)
	},
}

func init() {
	setCmd.AddCommand(setPositionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
