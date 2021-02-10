package exifdata

import (
	"fmt"
	"testing"
)

func TestExtractExif(t *testing.T) {
	//path := "/home/danil/qumo_flash/DCIM/Camera/P_20150122_202959.jpg"
	path := "/media/danil/TV/фотки/марк утренник/prazdnik_oseni/304.jpg"
	//path := "/home/danil/DCIM/.thumbnails/(2073)(3)1422960337371.png"

	data, err := Extract(path)
	if err != nil {
		panic(err)
	}
	//
	//createdAt, err:= data.CreatedAt()
	////fmt.Printf("%v\n", err)
	//fmt.Printf("%v\n", createdAt)

	deviceName := data.CreatedAt()
	//fmt.Printf("%v\n", err)
	fmt.Printf("%v\n", deviceName)

}
