package cmd

import (
	"github.com/spf13/cobra"
	"imagesCollector/processing"
)

func init() {
	rootCmd.AddCommand(verboseCmd)
}

var verboseCmd = &cobra.Command{
	Use:     "verbose",
	Short:   "shows files to be copied, but not copy",
	Example: `verbose`,
	Run: func(cmd *cobra.Command, args []string) {
		processing.Process(false)
	},
}
