/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	_ "encoding/json"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Termux API",
	Long:  `Termux API first apps cli.`,
	Run: func(cmd *cobra.Command, args []string) {
		getRandomBattery()
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)

}

// BatteryStatusResponse is returned by the BatteryStatus function.
type BatteryStatusResponse struct {
	Health      string
	Percentage  int
	Plugged     string
	Status      string
	Temperature float64
}

func getRandomBattery() {
	fmt.Println("Get random dad joke :P")

}

// BatteryStatus mimics the termux-battery-status script/call.
//func BatteryStatus() (*BatteryStatusResponse, error) {
//return batteryStatus(toolExec)
//}
//
//func batteryStatus(execF toolExecFunc) (*BatteryStatusResponse, error) {
//var resp BatteryStatusResponse
//bsrBytes, err := execF(nil, "BatteryStatus")
//if err != nil {
//return nil, err
//}
//err = json.Unmarshal(bsrBytes, &resp)
//if err != nil {
//return nil, err
//}
//return &resp, nil
//}
