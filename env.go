package envcrypt

import "os"

func OverrideEnv(key string, value string) {
	os.Setenv(key, value)
}

func OverrideAll(env map[string]string) {
	for key, value := range env {
		OverrideEnv(key, value)
	}
}

func SetEnv(key string, value string) {
	if _, set := os.LookupEnv(key); !set {
		OverrideEnv(key, value)
	}
}

func SetAll(env map[string]string) {
	for key, value := range env {
		SetEnv(key, value)
	}
}
