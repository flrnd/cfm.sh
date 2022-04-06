package cmd

import (
	"fmt"

	"github.com/flrnd/gorng"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Pass a space spearated list of words",
	Long:  "Pass a list of space separated words to choose from",
	Run: func(cmd *cobra.Command, args []string) {
		chooseOne(args)
	},
}

func chooseOne(args []string) {
	fmt.Println("cfm says, you should choose: ", args[gorng.Rng.Int()%len(args)])
}
