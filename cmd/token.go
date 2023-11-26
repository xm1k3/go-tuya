/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// tokenCmd represents the token command
var tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Get token",
	Long:  `Get token`,
	Run: func(cmd *cobra.Command, args []string) {
		client_id, _ := rootCmd.PersistentFlags().GetString("clientid")
		secret, _ := rootCmd.PersistentFlags().GetString("secret")
		device_id, _ := rootCmd.PersistentFlags().GetString("deviceid")
		token := GetToken(client_id, secret, device_id)
		fmt.Println(token)
	},
}

func init() {
	rootCmd.AddCommand(tokenCmd)
}
