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

	fmt.Println(config)
}
