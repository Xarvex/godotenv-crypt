package godotenvcrypt

import "io/fs"

func SetFS(filesystem fs.FS) {
	fsys = filesystem
}

func Parse(source []byte) (map[string]string, error) {
	var env map[string]string
	eachStatement(source, func(b []byte) {
		key, value, _ := environmentPair(b)
		env[key] = value
	})

	return env, nil
}

func Read(filenames ...string) (map[string]string, error) {
	var env map[string]string
	if err := readFiles(filenames, func(b []byte) error {
		fileEnv, err := Parse(b)
		if err != nil {
			return err
		}

		for key, value := range fileEnv {
			env[key] = value
		}

		return nil
	}); err != nil {
		return env, err
	}

	return env, nil
}

func Load(filenames ...string) error {
	env, err := Read(filenames...)
	if err != nil {
		return err
	}

	SetAll(env)

	return nil
}
