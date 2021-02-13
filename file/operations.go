package file

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

type FileRepr struct {
	Path  string
	FInfo os.FileInfo
	Error error
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

func Find(paths []string, exts map[string]struct{}) []FileRepr {
	var files []FileRepr
	for _, path := range paths {
		if IsDirExists(path) {
			err := filepath.Walk(path,
				func(path string, info os.FileInfo, err error) error {
					if err != nil {
						files = append(files, FileRepr{
							Error: err,
						})
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
				files = append(files, FileRepr{
					Error: err,
				})
			}
		}
	}

	return files
}

func isAllowedExt(path string, exts map[string]struct{}) bool {
	ext := strings.ToLower(filepath.Ext(path))
	if _, ok := exts[ext]; ok {
		return true
	}

	return false
}

func MkDirIfNotExist(path string) error {
	if !IsDirExists(path) {
		err := os.MkdirAll(path, 0700)
		if err != nil {
			return err
		}
	}

	return nil
}

func IsDirExists(path string) bool {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return true
	}

	return false
}
