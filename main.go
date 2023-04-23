package main

import (
	"flag"
	"fmt"
	"github.com/TBXark/chat-bot-go/configs"
	"github.com/TBXark/chat-bot-go/pkg/chat"
	"github.com/TBXark/chat-bot-go/pkg/dao"
	"log"
)

var (
	config  string
	version bool
)

func init() {
	flag.StringVar(&config, "config", "config.json", "path to config file")
	flag.BoolVar(&version, "version", false, "show version")
	flag.Parse()
}

func main() {
	if version {
		fmt.Printf("Version: %s\nBuildTime: %s\n", configs.Version, configs.BuildTime)
		return
	}
	cfg, err := configs.NewConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	db := dao.NewDatabase(cfg)
	app := chat.NewApp(cfg, dao.NewDao(db))
	app.Run()
}
