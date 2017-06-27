package main

import (
	"flag"

	"github.com/spf13/viper"
)

// TODO 機能実現スピード最優先での実装なので要リファクタ
func main() {
	configFilePath := flag.String("config", "_sample/config.toml", "Config File")
	template := flag.String("template", "_template/config.go", "Template File")
	flag.Parse()

	viper.SetConfigFile(*configFilePath)
	viper.ReadInConfig()
}
