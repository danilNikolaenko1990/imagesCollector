package exifdata

import (
	"fmt"
	"github.com/xiam/exif"
	"strings"
	"time"
)

const (
	tagDateAndTimeOriginal = "Date and Time (Original)"
	dateLayout             = "2006:01:02 15:04:05"
	defaultDeviceName      = "unknown-device"
	tagModel               = "Model"
	tagManufacturer        = "Manufacturer"
)

type ExifData struct {
	filePath string
	tags     map[string]string
}

func NewExifData(tags map[string]string, filepath string) ExifData {
	return ExifData{
		tags:     tags,
		filePath: filepath,
	}
}

func (ed *ExifData) CreatedAt() (time.Time, error) {
	if date, ok := ed.tags[tagDateAndTimeOriginal]; ok {
		t, err := time.Parse(dateLayout, date)
		if err != nil {
			return time.Time{}, fmt.Errorf("file %s cannot parse date", ed.filePath)
		}
		return t, nil
	}

	return time.Time{}, fmt.Errorf("no date tag %s", ed.filePath)
}

func (ed *ExifData) DeviceName() string {
	devName := fmt.Sprintf("%s_%s", ed.extractManufact(), ed.extractModel())
	if len(devName) < 1 {
		return defaultDeviceName
	}
	return strings.Replace(devName, " ", "_", -1)
}

func (ed *ExifData) extractManufact() string {
	m := ""
	if m, ok := ed.tags[tagManufacturer]; ok {
		return m
	}

	return m
}

func (ed *ExifData) extractModel() string {
	m := ""
	if m, ok := ed.tags[tagModel]; ok {
		return m
	}

	return m
}

func Extract(path string) (ExifData, error) {
	data, err := exif.Read(path)
	if err != nil {
		return ExifData{}, fmt.Errorf("%s - %s", path, err.Error())
	}

	return NewExifData(data.Tags, path), err
}
