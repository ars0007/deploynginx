package cmd

import (
	"deploynginx/handlers"
	"fmt"

	"github.com/spf13/cobra"
)

// upgradeCmd represents the upgrade command
var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrades the deployment with given version",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		err := handlers.UpdateVersion(cmd)
		if err != nil {
			panic(err)
		}

		fmt.Println("Version updated successfully")
	},
}

func init() {
	rootCmd.AddCommand(upgradeCmd)
	upgradeCmd.Flags().String("version", "", "version to upgrade the deployment with")
}
