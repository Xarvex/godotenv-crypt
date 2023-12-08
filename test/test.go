package main

import (
	"embed"
	"xarvex/godotenvcrypt"
)

//go:embed .env
var dotenv embed.FS

func main() {
	godotenvcrypt.SetFS(dotenv)
	godotenvcrypt.Load()
}
