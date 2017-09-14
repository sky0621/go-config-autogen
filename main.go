package main

import (
	"os"

	"fmt"

	"time"

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
		fmt.Printf("key: %v\nval: %#v\n", k, v)
		switch v.(type) {
		case string:
			fmt.Printf("[value.type = string]\n\n")
			continue
		case int:
			fmt.Printf("[value.type = int]\n\n")
			continue
		case time.Time:
			fmt.Printf("[value.type = time.Time]\n\n")
			continue
		case []string:
			fmt.Printf("[value.type = []string]\n\n")
			continue
		case map[string]interface{}:
			fmt.Printf("[value.type = map[string]interface{}\n\n")
			vm := v.(map[string]interface{})
			printConfig(vm)
		default:
			fmt.Println(v)

		}
	}
}
