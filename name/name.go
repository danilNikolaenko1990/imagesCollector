package name

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"imagesCollector/exifdata"
	"os"
	"path/filepath"
	"strconv"
)

func fileName(fInfo os.FileInfo, exif exifdata.ExifData) string {
	randStr := randStr()
	return fmt.Sprintf("IMG-%s-%s-%s", dateStr(exif), exif.DeviceName(), randStr)
}

func PathName(targetFolder string, exif exifdata.ExifData) string {
	ctime := exif.CreatedAt()

	return filepath.Join(targetFolder, strconv.Itoa(ctime.Year()), ctime.Month().String())
}

func FullFileName(targetFolder string, fInfo os.FileInfo, exif exifdata.ExifData) string {
	pathName := PathName(targetFolder, exif)
	fName := fileName(fInfo, exif)

	return filepath.Join(pathName, fName)
}

func dateStr(exif exifdata.ExifData) string {
	createdAt := exif.CreatedAt()

	return createdAt.Format("01_02_2006_15_04_05")
}

func randStr() string {
	randBytes := make([]byte, 3)
	rand.Read(randBytes)

	return hex.EncodeToString(randBytes)
}
