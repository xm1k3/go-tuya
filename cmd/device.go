/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deviceCmd represents the device command
var deviceCmd = &cobra.Command{
	Use:   "device",
	Short: "Get device",
	Long:  `Get device`,
	Run: func(cmd *cobra.Command, args []string) {
		client_id, _ := rootCmd.PersistentFlags().GetString("clientid")
		secret, _ := rootCmd.PersistentFlags().GetString("secret")
		device_id, _ := rootCmd.PersistentFlags().GetString("deviceid")
		GetToken(client_id, secret, device_id)
		device := GetDevice(client_id, secret, device_id)
		fmt.Println(device)
	},
}

func init() {
	rootCmd.AddCommand(deviceCmd)
}
