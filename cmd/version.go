package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Outputs cfm.sh version number",
	Long:  "Outputs cfm.sh version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cfm.sh v0.1")
	},
}
