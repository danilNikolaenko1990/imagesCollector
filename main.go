package main

import (
	"imagesCollector/conf_parser"
	"imagesCollector/processing"
)

const confPath = "config.yml"

func main() {
	c := getConf()
	processing.Process(c)
}

func getConf() *conf_parser.Config {
	c, err := conf_parser.Extract(confPath)
	if err != nil {
		panic(err)
	}
	return c
}
