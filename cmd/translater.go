package main

import (
	"github.com/andreysidor4uk/cli-translater/internal/cmd"
	"github.com/andreysidor4uk/cli-translater/internal/config"
	"github.com/andreysidor4uk/cli-translater/internal/translaters/yandex"
)

func main() {
	tranlater := yandex.New(config.YandexApiKey, config.YandexFolderId)
	cmd.Execute(tranlater)
}
