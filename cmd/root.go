package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "cfm.sh"}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Panic(err)
	}
}
