package godotenvcrypt

import (
	"fmt"
	"io/fs"
)

func SetFS(filesystem fs.FS) {
	fsys = filesystem
}

func Parse(src []byte) (map[string]string, error) {
	var environment map[string]string
	eachStatement(src, func(b []byte) {
		var key, value string
		key, value, _ = environmentPair(b)
		fmt.Println(key, value)
	})

	return environment, nil
}

func Load(filenames ...string) error {
	if len(filenames) == 0 {
		filenames = []string{".env"}
	}
	if err := readFiles(filenames, func(b []byte) {
		Parse(b)
	}); err != nil {
		return err
	}

	return nil
}
