package main

import (
	"embed"
	"xarvex/envcrypt"
)

//go:embed .env
var dotenv embed.FS

func main() {
	envcrypt.SetFS(dotenv)
	envcrypt.Load()
}
