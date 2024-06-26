package cmd

import (
	"deploynginx/handlers"
	"fmt"

	"github.com/spf13/cobra"
)

// scaleCmd represents the scale command
var scaleCmd = &cobra.Command{
	Use:   "scale",
	Short: "Scales the number of pods in the deployment",
	Run: func(cmd *cobra.Command, args []string) {
		err := handlers.ScaleDeployment(cmd)
		if err != nil {
			panic(err)
		}

		fmt.Println("Successfully scaled")
	},
}

func init() {
	rootCmd.AddCommand(scaleCmd)
	scaleCmd.Flags().Int("replicas", 1, "Number of replicas to scale the deplyment")
}
