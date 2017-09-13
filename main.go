package main

import (
	"os"

	"fmt"

	"github.com/naoina/toml"
)

func main() {
	f, err := os.Open("example.toml")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	var config map[string]interface{}
	if err := toml.NewDecoder(f).Decode(&config); err != nil {
		panic(err)
	}

	printConfig(config)
}

func printConfig(conf map[string]interface{}) {
	for k, v := range conf {
		fmt.Printf("key: %v\nval: %#v\n\n", k, v)
		switch v.(type) {
		case string:
			fmt.Println("$ $ $ $ $ string $ $ $ $ $")
			continue
		case map[string]interface{}:
			fmt.Println("$ $ $ $ $ map[string]interface{} $ $ $ $ $")
			vm := v.(map[string]interface{})
			printConfig(vm)
		default:
			fmt.Println("$ $ $ $ $ default $ $ $ $ $")

		}
	}
}
