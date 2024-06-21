package main

import (
	"os"
)

func ArquivosPontoFolder() string {
	path, env := os.LookupEnv("ARQUIVOSPONTO_FOLDER")
	if !env {
		path, _ = os.UserHomeDir()
		path += "/arquivosponto/"
	}
	return path
}
