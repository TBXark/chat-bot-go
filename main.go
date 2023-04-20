package main

import (
	"flag"
	"github.com/TBXark/chat-bot-go/configs"
	"github.com/TBXark/chat-bot-go/pkg/chat"
	"github.com/TBXark/chat-bot-go/pkg/dao"
	"log"
)

var buildVersion = "dev"

func main() {
	cfg := flag.String("config", "config.json", "path to config file")
	ver := flag.Bool("version", false, "show version")
	flag.Parse()
	if *ver {
		log.Printf("Version: %s", buildVersion)
		return
	}
	config, err := configs.NewConfig(*cfg)
	if err != nil {
		log.Fatal(err)
	}
	db := dao.NewDatabase(config)
	app := chat.NewApp(config, dao.NewDao(db))
	app.Run()
}
