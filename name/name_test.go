package name

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

	res := PathName(targetDir, data)
	fmt.Println(res)
}

func TestFullFileName(t *testing.T) {
	path := "/home/danil/DCIM/Camera/P_20141207_181719.jpg"
	targetDir := "/media/danil/TV/STRUCTURED_IMAGES/"
	exifData, _ := exifdata.Extract(path)
	fileinfo, _ := os.Stat(path)

	res := FullFileName(targetDir, fileinfo, exifData)

	print(res)
}
