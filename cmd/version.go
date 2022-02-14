package cmd

import (
	"fmt"

	"github.com/muesli/coral"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &coral.Command{
	Use:   "version",
	Short: "Outputs cfm.sh version number",
	Long:  "Outputs cfm.sh version number",
	Run: func(cmd *coral.Command, args []string) {
		fmt.Println("cfm.sh v0.1")
	},
}
