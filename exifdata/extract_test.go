package exifdata

import (
	"fmt"
	"testing"
)

func TestExtractExif(t *testing.T) {
	//path := "/home/danil/qumo_flash/DCIM/Camera/P_20150122_202959.jpg"
	path := "/home/danil/DCIM2/DCIM/101NIKON/DSCN2685.JPG"
	//path := "/home/danil/DCIM/.thumbnails/(2073)(3)1422960337371.png"

	data, err := Extract(path)
	if err != nil {
		panic(err)
	}
	//
	//createdAt, err:= data.CreatedAt()
	////fmt.Printf("%v\n", err)
	//fmt.Printf("%v\n", createdAt)

	deviceName :=  data.DeviceName()
	//fmt.Printf("%v\n", err)
	fmt.Printf("%s\n", deviceName)

}
