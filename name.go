package main

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
	return fmt.Sprintf("IMG-%s-%s-%s", dateStr(fInfo, exif), exif.DeviceName(), randStr)
}

func pathName(targetFolder string, exif exifdata.ExifData) (string, error) {
	ctime, err := exif.CreatedAt()
	if err != nil {
		return "", err
	}

	return filepath.Join(targetFolder, strconv.Itoa(ctime.Year()), ctime.Month().String()), nil
}

func FullFileName(targetFolder string, fInfo os.FileInfo, exif exifdata.ExifData) (string, error) {
	pathName, err := pathName(targetFolder, exif)
	if err != nil {
		return "", err
	}
	fName := fileName(fInfo, exif)

	return filepath.Join(pathName, fName), err
}

func dateStr(fInfo os.FileInfo, exif exifdata.ExifData) string {
	createdAt, err := exif.CreatedAt()
	if err != nil {
		createdAt = fInfo.ModTime()
	}

	return createdAt.Format("01_02_2006_15_04_05")
}

func randStr() string {
	randBytes := make([]byte, 3)
	rand.Read(randBytes)

	return hex.EncodeToString(randBytes)
}
