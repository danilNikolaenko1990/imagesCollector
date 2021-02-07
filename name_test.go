package main

import (
	"fmt"
	"imagesCollector/exifdata"
	"os"
	"testing"
)

func TestFileName(t *testing.T) {
	path := "/home/danil/DCIM/Camera/P_20141207_181719.jpg"
	data, _ := exifdata.Extract(path)

	fileinfo, _ := os.Stat(path)
	res := fileName(fileinfo, data)
	fmt.Println(res)
}

func TestPathName(t *testing.T) {
	path := "/home/danil/DCIM/Camera/P_20141207_181719.jpg"
	targetDir := "/media/danil/TV/STRUCTURED_IMAGES/"
	data, _ := exifdata.Extract(path)

	res, err := pathName(targetDir, data)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

func TestFullFileName(t *testing.T) {
	path := "/home/danil/DCIM/Camera/P_20141207_181719.jpg"
	targetDir := "/media/danil/TV/STRUCTURED_IMAGES/"
	exifData, _ := exifdata.Extract(path)
	fileinfo, _ := os.Stat(path)

	res, err := FullFileName(targetDir, fileinfo, exifData)
	if err != nil {
		panic(err)
	}

	print(res)
}



