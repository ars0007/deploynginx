/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// upgradeCmd represents the upgrade command
var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrades the deployment with given version",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("upgrade called :: ", cmd)
	},
}

func init() {
	rootCmd.AddCommand(upgradeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// upgradeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// upgradeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	upgradeCmd.Flags().Int("version", -1, "version to upgrade the deployment with")
}
