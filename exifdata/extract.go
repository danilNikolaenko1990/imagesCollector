package exifdata

import (
	"fmt"
	"github.com/xiam/exif"
	"os"
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
	fileMTime time.Time
	filePath  string
	tags      map[string]string
}

func NewExifData(tags map[string]string, filepath string) ExifData {
	return ExifData{
		tags:     tags,
		filePath: filepath,
	}
}

//todo remove err from signature
func (ed *ExifData) CreatedAt() (time.Time, error) {
	if date, ok := ed.tags[tagDateAndTimeOriginal]; ok {
		t, err := time.Parse(dateLayout, date)
		if err != nil {
			return fileMTime(ed.filePath), nil
		}
		return t, nil
	}

	return fileMTime(ed.filePath), nil
}

func fileMTime(path string) time.Time {
	info, err := os.Stat(path)
	mTime := time.Now()
	if err == nil {
		mTime = info.ModTime()
	}

	return mTime
}

func (ed *ExifData) DeviceName() string {
	if len(ed.extractManufact()) < 1 && len(ed.extractModel()) < 1 {
		return defaultDeviceName
	}
	devName := fmt.Sprintf("%s_%s", ed.extractManufact(), ed.extractModel())
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
