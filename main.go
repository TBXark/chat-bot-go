package main

import (
	"flag"
	"github.com/TBXark/chat-bot-go/configs"
	"github.com/TBXark/chat-bot-go/pkg/chat"
	"github.com/TBXark/chat-bot-go/pkg/dao"
	"log"
)

func main() {
	cfg := flag.String("config", "config.json", "path to config file")
	flag.Parse()
	config, err := configs.NewConfig(*cfg)
	if err != nil {
		log.Fatal(err)
	}
	db := dao.NewDatabase(config)
	app := chat.NewApp(config, dao.NewDao(db))
	app.Run()
}
