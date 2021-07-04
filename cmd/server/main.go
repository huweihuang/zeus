package main

import (
	"os"

	"github.com/huweihuang/gin-api-frame/cmd/server/app"
)

const (
	bashCompleteFile = "/etc/bash_completion.d/apiserver.bash_complete"
)

func main() {
	command := app.NewServerCommand()
	command.GenBashCompletionFile(bashCompleteFile)
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
