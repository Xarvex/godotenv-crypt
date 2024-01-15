package envcrypt

import (
	"embed"
	"testing"
)

//go:embed .env
var embeddedDotenv embed.FS

func TestEmbeddedSetFS(t *testing.T) {
	SetFS(embeddedDotenv)
	if fsys != embeddedDotenv {
		t.Error("fsys not set")
	}
}

func TestEmbeddedSetFSNil(t *testing.T) {
	TestEmbeddedSetFS(t) // do not rely on run order
	SetFS(nil)
	if fsys != nil {
		t.Error("fsys not unset")
	}
}

func TestEmbeddedLoad(t *testing.T) {
	SetFS(embeddedDotenv)
	TestLoad(t)
}

func TestEmbeddedLoadNonExistent(t *testing.T) {
	SetFS(embeddedDotenv)
	TestLoadNonExistent(t)
}
