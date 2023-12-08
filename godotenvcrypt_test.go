package godotenvcrypt

import "testing"

func TestLoad(t *testing.T) {
	if err := Load(); err != nil {
		t.Error(err)
	}
}

func TestLoadNonExistent(t *testing.T) {
	err := Load("nonexistent")
	if err == nil {
		t.Error("No error reported for nonexistent file")
	}
}
