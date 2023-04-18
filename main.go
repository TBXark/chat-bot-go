package main

import (
	"flag"
	"github.com/TBXark/chat-bot-go/configs"
	"github.com/TBXark/chat-bot-go/pkg"
	"log"
)

func main() {
	cfg := flag.String("config", "config.json", "path to config file")
	flag.Parse()
	config, err := configs.NewConfig(*cfg)
	if err != nil {
		log.Fatal(err)
	}
	pkg.NewApp(config).Run()
}
