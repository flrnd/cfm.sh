package cmd

import (
	"log"

	"github.com/muesli/coral"
)

var rootCmd = &coral.Command{Use: "cfm.sh"}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Panic(err)
	}
}
