package main

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"imagesCollector/exifdata"
	"imagesCollector/file"
)

const targetFolderToCopy = "/media/danil/TV/STRUCTURED_IMAGES"

func main() {
	dirNames := []string{
		"/home/danil/Dropbox",
		"/home/danil/Dropbox (Old)",
		"/home/danil/dropbox_old2",
		"/home/danil/dell_latitude",
		"/home/danil/DCIM",
		"/home/danil/DCIM2",
		"/home/danil/Documents",
		"/home/danil/Pictures",
		"/home/danil/qumo_flash",
		"/home/danil/книги",
		"/home/danil/книги марку",
		"/media/danil/TV/фотки",
		"/media/danil/TV/xiaomi_danil",
		"/media/danil/TV/kvitancii",
	}

	fReprs, err := file.FindWithExtensions(dirNames, map[string]struct{}{".jpg": {}})

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
