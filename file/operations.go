package file

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type FileRepr struct {
	Path  string
	FInfo os.FileInfo
}

func fileExists(path string) bool {
	_, err := os.Stat(path)

	return err == nil
}

func move(from string) {
	to := generateNewPath(from)

	err := Copy(from, to)

	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	remove(from)
}

func generateNewPath(from string) string {
	oldName := filepath.Base(from)

	return TempFileName("MVD-", "-") + oldName

}

func TempFileName(prefix, suffix string) string {
	randBytes := make([]byte, 4)
	rand.Read(randBytes)
	return filepath.Join("/media/danil/TV/REMOVED_FILES", prefix+hex.EncodeToString(randBytes)+suffix)
}

func Copy(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}

func remove(path string) {
	e := os.Remove(path)
	if e != nil {
		fmt.Printf("can't remove file %s", path)
	}
}

func FindWithExtensions(paths []string, exts map[string]struct{}) ([]FileRepr, error) {
	var files []FileRepr
	for _, path := range paths {
		if isDirExists(path) {
			err := filepath.Walk(path,
				func(path string, info os.FileInfo, err error) error {
					if err != nil {
						return err
					}
					if !info.IsDir() && isAllowedExt(info.Name(), exts) {
						files = append(files, FileRepr{
							Path:  path,
							FInfo: info,
						})
					}

					return nil
				})
			if err != nil {
				return nil, err
			}
		}
	}

	return files, nil
}

func isAllowedExt(path string, exts map[string]struct{}) bool {
	ext := strings.ToLower(filepath.Ext(path))
	if _, ok := exts[ext]; ok {
		return true
	}

	return false
}

func MkDirIfNotExist(path string) error {
	if !isDirExists(path) {
		err := os.MkdirAll(path, 0700)
		if err != nil {
			return err
		}
	}

	return nil
}

func isDirExists(path string) bool {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return true
	}

	return false
}
