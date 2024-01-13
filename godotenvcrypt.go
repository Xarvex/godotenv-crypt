package godotenvcrypt

func Parse(source []byte) (map[string]string, error) {
	var env map[string]string
	if err := eachStatement(source, func(b []byte) error {
		key, value, err := environmentPair(b)
		if err != nil {
			return err
		}

		env[key] = value

		return nil
	}); err != nil {
		return env, err
	}

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
