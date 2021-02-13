package main

import (
	"fmt"
	"imagesCollector/conf_parser"
	"imagesCollector/processing"
	"time"
)

const confPath = "config.yml"

func main() {
	c := getConf()
	measure(func() {
		processing.Process(c)
	})
}

func measure(runner func()) {
	start := time.Now()
	runner()
	duration := time.Since(start)
	fmt.Println(duration)
	fmt.Println(duration.Seconds())
}

func getConf() *conf_parser.Config {
	c, err := conf_parser.Extract(confPath)
	if err != nil {
		panic(err)
	}
	return c
}
