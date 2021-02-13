package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:  "imageCollector",
	Long: "image-collector takes all the photos from folders you picked and copies to the target folder. Also it groups pictures by years and months based on EXIF data from file",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
