package config

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var (
	YandexApiKey   string
	YandexFolderId string
)

func init() {
	initConfig()
}

func initConfig() {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	viper.AddConfigPath(home)
	viper.SetConfigName(".translater")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintln(os.Stderr, "Can't read config:", err)
		os.Exit(1)
	}

	YandexApiKey = viper.GetString("yandexapikey")
	if YandexApiKey == "" {
		errorParam("yandexapikey")
	}
	YandexFolderId = viper.GetString("yandexfolderid")
	if YandexApiKey == "" {
		errorParam("yandexfolderid")
	}
}

func errorParam(paramName string) {
	fmt.Fprintf(os.Stderr, "the %v parameter is not filled in\n", paramName)
	os.Exit(1)
}
