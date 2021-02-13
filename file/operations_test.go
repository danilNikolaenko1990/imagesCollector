package file

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWalkFile(t *testing.T) {
	dirNames := []string{"/home/danil/qumo_flash", "/home/danil/DCIM/"}

	exts := map[string]struct{}{".jpg": {}}
	result, _ := Find(dirNames, exts)
	fmt.Printf("%v\n", result)

}

func TestIsAllowedExt(t *testing.T) {
	exts := map[string]struct{}{".jpg": {}}
	path := "/home/danil/DCIM/.thumbnails/(2073)(3)1422960337371.jpg"
	result := isAllowedExt(path, exts)
	assert.True(t, result)

	path = "/home/danil/DCIM/.thumbnails/(2073)(3)1422960337371.png"
	result = isAllowedExt(path, exts)
	assert.False(t, result)
}

func TestMkDirIfNotExist(t *testing.T) {
	dirname := "/media/danil/TV/STRUCTURED_IMAGES/Jun/jul/lll"

	err := MkDirIfNotExist(dirname)
	fmt.Sprintf("%v", err)
}

func TestIsDirExists(t *testing.T) {
	dirname := "/home/danil"
	fmt.Print(IsDirExists(dirname))
}
