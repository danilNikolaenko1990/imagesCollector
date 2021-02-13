package file

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

//todo make tests based on real files
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
