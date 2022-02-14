package cmd

import (
	"fmt"

	"github.com/flrnd/gorng"
	"github.com/muesli/coral"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &coral.Command{
	Use:   "list",
	Short: "Pass a space spearated list of words",
	Long:  "Pass a list of space separated words to choose from",
	Run: func(cmd *coral.Command, args []string) {
		chooseOne(args)
	},
}

func chooseOne(args []string) {
	fmt.Println("cfm says, you should choose: ", args[gorng.Rng.Int()%len(args)])
}
