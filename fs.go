package godotenvcrypt

import (
	"io/fs"
	"os"
)

var fsys fs.FS

func readFile(name string) ([]byte, error) {
	if fsys == nil {
		return os.ReadFile(name)
	} else {
		return fs.ReadFile(fsys, name)
	}
}

func readFiles(names []string, callback func([]byte)) error {
	for _, name := range names {
		content, err := readFile(name)
		if err != nil {
			return err
		}

		callback(content)
	}

	return nil
}
