package dialog

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/manifoldco/promptui"
	"imagesCollector/conf_parser"
	"strings"
)

func IsUserAgreedWithSettings(config *conf_parser.Config) bool {
	fmt.Println("this is your config based on config.yml file")
	fmt.Println("check this, please")
	fmt.Println("folders to be scanned:")
	for _, folder := range config.DirsToScan {
		fmt.Printf(" - %s\n", folder)
	}
	fmt.Println(strings.Repeat("-", 10))
	fmt.Println("target folder to copy:")
	fmt.Printf(" - %s\n", config.TargetFolderToCopy)
	fmt.Println(strings.Repeat("-", 10))
	fmt.Println("file extensions to be copied:")
	for _, folder := range config.Exts {
		fmt.Printf(" - %s\n", folder)
	}
	return confirmConfig()
}

func confirmConfig() bool {
	prompt := promptui.Select{
		Label: "Do you agree with settings? [Yes/No]",
		Items: []string{"No", "Yes"},
	}
	_, result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}
	return result == "Yes"
}
