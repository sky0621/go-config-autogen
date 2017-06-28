package main

import (
	"flag"

	"fmt"

	"github.com/BurntSushi/toml"
)

// TODO 機能実現スピード最優先での実装なので要リファクタ
func main() {
	configFilePath := flag.String("config", "_sample/config.toml", "Config File")
	//template := flag.String("template", "_template/config.go", "Template File")
	flag.Parse()

	v := make(map[string]interface{}, 1)
	metadata, err := toml.Decode(*configFilePath, v)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", v)
	fmt.Printf("%#v\n", metadata)
	//viper.SetConfigFile(*configFilePath)
	//viper.ReadInConfig()
}
