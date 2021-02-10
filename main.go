package main

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"imagesCollector/conf_extrac"
	"imagesCollector/exifdata"
	"imagesCollector/file"
)

const targetFolderToCopy = "/media/danil/TV/STRUCTURED_IMAGES"
const confPath = "config.yml"

func main() {
	c := getConf()
	fReprs, err := file.FindWithExtensions(c.DirsToScan, c.Extensions())

	if err != nil {
		panic(err.Error())
	}

	for _, fRepr := range fReprs {
		exif, err := exifdata.Extract(fRepr.Path)

		if err != nil {
			log.Warn(err)
			continue
		}

		pathToCopy := PathName(targetFolderToCopy, exif)
		err = file.MkDirIfNotExist(pathToCopy)
		if err != nil {
			log.Warnf("failed to create dir [%s]", err.Error())
			continue
		}
		fullFileName := FullFileName(targetFolderToCopy, fRepr.FInfo, exif)
		err = file.Copy(fRepr.Path, fullFileName)
		if err != nil {
			log.Warnf("failed to copy file [%s]", err.Error())
			continue
		}

		fmt.Printf("file [%s] copied to [%s]\n", fRepr.Path, fullFileName)
	}
}

func getConf() *conf_extrac.Data {
	c, err := conf_extrac.Extract(confPath)
	if err != nil {
		panic(err)
	}
	return c
}
