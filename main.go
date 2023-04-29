package main

import (
	"flag"
	"fmt"
	"github.com/TBXark/chat-bot-go/configs"
	"github.com/TBXark/chat-bot-go/pkg/chat"
	"github.com/TBXark/chat-bot-go/pkg/dao"
	"log"
	"strings"
	"sync"
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
	cfgList, err := configs.NewConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	wg := sync.WaitGroup{}
	for _, cfg := range cfgList {
		c := cfg
		wg.Add(1)
		db := dao.NewDatabase(c)
		app := chat.NewApp(c, dao.NewDao(db))
		go func() {
			defer func() {
				if err := recover(); err != nil {
					log.Println(err)
				}
			}()
			log.Printf("Start chat server: %s\n", strings.Split(c.Telegram.Token, ":")[0])
			app.Run()
			wg.Done()
		}()
	}
	wg.Wait()
}
