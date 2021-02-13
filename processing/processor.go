package processing

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"imagesCollector/conf_parser"
	"imagesCollector/exif_data"
	"imagesCollector/file"
	"imagesCollector/name"
	"runtime"
	"sync"
)

func Process(conf *conf_parser.Config) {
	files := file.Find(conf.DirsToScan, conf.Extensions())
	workers := runtime.NumCPU()
	processAsyncAndWait(workers, func() {
		for fileItem := range files {
			if fileItem.Error == nil {
				performCopy(fileItem, conf)
			} else {
				fmt.Printf("fileItem %s, [%s]", fileItem.Path, fileItem.Error.Error())
			}
		}
	})
}

func processAsyncAndWait(workers int, performer func()) {
	wg := &sync.WaitGroup{}
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			performer()
		}()
	}
	wg.Wait()
}

func performCopy(fRepr file.FileRepr, conf *conf_parser.Config) {
	exif, err := exif_data.Extract(fRepr.Path)

	if err != nil {
		log.Warn(err)
		return
	}

	pathToCopy := name.PathName(conf.TargetFolderToCopy, exif)
	err = file.MkDirIfNotExist(pathToCopy)
	if err != nil {
		log.Warnf("failed to create dir [%s]", err.Error())
		return
	}
	fullFileName := name.FullFileName(conf.TargetFolderToCopy, fRepr.FInfo, exif)
	err = file.Copy(fRepr.Path, fullFileName)
	if err != nil {
		log.Warnf("failed to copy file [%s]", err.Error())
		return
	}

	fmt.Printf("file [%s] copied to [%s]\n", fRepr.Path, fullFileName)
}
