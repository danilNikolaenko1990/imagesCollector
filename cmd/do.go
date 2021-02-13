package cmd

import (
	"github.com/spf13/cobra"
	"imagesCollector/processing"
)

func init() {
	rootCmd.AddCommand(doCmd)
}

var doCmd = &cobra.Command{
	Use:     "do",
	Short:   "copies files into the target folder",
	Example: `do`,
	Run: func(cmd *cobra.Command, args []string) {
		processing.Process(true)
	},
}
