package conf_extrac

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExtract(t *testing.T) {
	path := "_files/config.yml"
	result, err := Extract(path)
	assert.Nil(t, err)
	assert.Equal(t, "/home/test_user/one_folder", result.DirsToScan[0])
	assert.Equal(t, "/media/test_user/DRIVE_FOR_TV/second_folder", result.DirsToScan[1])
	assert.Equal(t, map[string]struct{}{".jpg": {}, ".png": {}}, result.Extensions())
}
