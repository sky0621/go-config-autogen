package main

import (
	"flag"
	"fmt"

	gotoml "github.com/pelletier/go-toml"
)

// TODO 機能実現スピード最優先での実装なので要リファクタ
func main() {
	configFilePath := flag.String("config", "_sample/config.toml", "Config File")
	//template := flag.String("template", "_template/config.go", "Template File")
	flag.Parse()

	tree, err := gotoml.LoadFile(*configFilePath)
	if err != nil {
		panic(err)
	}

	fmt.Printf("[tree]%#v\n\n", tree)
	fmt.Printf("[tree.ToMap()]%#v\n\n", tree.ToMap())
	fmt.Printf("[tree.Keys()]%#v\n\n", tree.Keys())
	fmt.Printf("[tree.String()]%#v\n\n", tree.String())
}
