package main

import "flag"

// TODO 機能実現スピード最優先での実装なので要リファクタ
func main() {
	cfg := flag.String("config", "../_sample/config.toml", "Config File")
	template := flag.String("template", "../../template/sqs.md", "Template File")

}
