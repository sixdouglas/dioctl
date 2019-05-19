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
	"github.com/sixdouglas/suncalc"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"time"
)

// onCmd represents the on command
var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "Get the times",
	Long:  `Get all the times.`,
	Run: func(cmd *cobra.Command, args []string) {
		now := time.Now()

		if !viper.IsSet("latitude") || !viper.IsSet("longitude") {
			if !viper.IsSet("latitude") {
				fmt.Println("Latitude has not been set. Consider setting it using command `dioct set position`")
			}

			if !viper.IsSet("longitude") {
				fmt.Println("Longitude has not been set. Consider setting it using command `dioct set position`")
			}
		} else {
			latitude := viper.GetFloat64("latitude")
			longitude := viper.GetFloat64("longitude")
			fmt.Printf("Current latitude is %f\n", "latitude")
			fmt.Printf("Current longitude is %f\n", "longitude")

			times := suncalc.GetTimes(now, latitude, longitude)

			for i := 0; i < len(times); i++ {
				oneTime := times[i]

				fmt.Printf("%-13s %d-%02d-%02d %02d:%02d:%02d\n", string(oneTime.MorningName),
					oneTime.Time.Year(), oneTime.Time.Month(), oneTime.Time.Day(),
					oneTime.Time.Hour(), oneTime.Time.Minute(), oneTime.Time.Second())
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(timeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// onCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// onCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
