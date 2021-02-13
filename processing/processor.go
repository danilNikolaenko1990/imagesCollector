package processing

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"imagesCollector/conf_parser"
	"imagesCollector/dialog"
	"imagesCollector/exif_data"
	"imagesCollector/file"
	"imagesCollector/name"
	"os"
)

const confPath = "config.yml"

func Process(needCopy bool) {
	conf := getConf()
	if dialog.IsUserAgreedWithSettings(conf) {
		files := file.Find(conf.DirsToScan, conf.Extensions())
		total := len(files)
		for id, fileItem := range files {
			fmt.Printf("current: %d, total: %d ", id, total)
			if fileItem.Error == nil {
				performCopy(fileItem, conf, needCopy)
			} else {
				fmt.Printf("fileItem %s, [%s]", fileItem.Path, fileItem.Error.Error())
			}
		}
	}
}

func getConf() *conf_parser.Config {
	c, err := conf_parser.Extract(confPath)
	if err != nil {
		fmt.Println("copy config.yml.dist and place it beside with config.yml name and write correct parameters then")
		os.Exit(1)
	}
	return c
}

func performCopy(fRepr file.FileRepr, conf *conf_parser.Config, needCopy bool) {
	exif, err := exif_data.Extract(fRepr.Path)

	if err != nil {
		fmt.Printf("file [%s] ", err.Error())
		return
	}

	if needCopy {
		pathToCopy := name.PathName(conf.TargetFolderToCopy, exif)
		err = file.MkDirIfNotExist(pathToCopy)
		if err != nil {
			fmt.Printf("failed to create dir [%s]", err.Error())
			return
		}
	}

	fullFileName := name.FullFileName(conf.TargetFolderToCopy, fRepr.FInfo, exif)
	if needCopy {
		err = file.Copy(fRepr.Path, fullFileName)
		if err != nil {
			log.Warnf("failed to copy file [%s]", err.Error())
			return
		}
		fmt.Printf("file [%s] copied to [%s]\n", fRepr.Path, fullFileName)
		return
	}
	fmt.Printf("file [%s] could be copy to [%s]\n", fRepr.Path, fullFileName)
}
