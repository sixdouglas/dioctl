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
var setEmitterCmd = &cobra.Command{
	Use:   "emitter [pinId]",
	Short: "Set the emitter pin ID",
	Long: `Set the emitter pin ID.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		var emitterId uint64
		var err error
		emitterId, err = strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			log.Fatalf("unable to decode emitter id, %v", err)
		}
		fmt.Printf("Emitter pin id is %d\n", viper.GetInt("emitterpinid"))
		viper.Set("emitterPinId", emitterId)
		err = viper.WriteConfig()
		if err != nil {
			log.Fatalf("unable to save changes, %v", err)
		}
		fmt.Printf("Emitter pin id set to: %d", emitterId)
	},
}

func init() {
	setCmd.AddCommand(setEmitterCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
