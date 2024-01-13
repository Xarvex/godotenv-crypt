package godotenvcrypt

import (
	"io/fs"
	"os"
)

var fsys fs.FS

func SetFS(filesystem fs.FS) {
	fsys = filesystem
}

func readFile(name string) ([]byte, error) {
	if fsys == nil {
		return os.ReadFile(name)
	} else {
		return fs.ReadFile(fsys, name)
	}
}

func readFiles(names []string, callback func([]byte) error) error {
	for _, name := range names {
		content, err := readFile(name)
		if err != nil {
			return err
		}

		if err := callback(content); err != nil {
			return err
		}
	}

	return nil
}
