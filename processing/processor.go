package processing

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"imagesCollector/conf_parser"
	"imagesCollector/exif_data"
	"imagesCollector/file"
	"imagesCollector/name"
)

func Process(conf *conf_parser.Config) {
	fReprs, err := file.FindWithExtensions(conf.DirsToScan, conf.Extensions())

	if err != nil {
		panic(err.Error())
	}

	for _, fRepr := range fReprs {
		exif, err := exif_data.Extract(fRepr.Path)

		if err != nil {
			log.Warn(err)
			continue
		}

		pathToCopy := name.PathName(conf.TargetFolderToCopy, exif)
		err = file.MkDirIfNotExist(pathToCopy)
		if err != nil {
			log.Warnf("failed to create dir [%s]", err.Error())
			continue
		}
		fullFileName := name.FullFileName(conf.TargetFolderToCopy, fRepr.FInfo, exif)
		err = file.Copy(fRepr.Path, fullFileName)
		if err != nil {
			log.Warnf("failed to copy file [%s]", err.Error())
			continue
		}

		fmt.Printf("file [%s] copied to [%s]\n", fRepr.Path, fullFileName)
	}
}
