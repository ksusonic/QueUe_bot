package main

import (
	tele "gopkg.in/telebot.v3"
	"log"
	"os"
	"time"
)

func main() {
	if os.Getenv("TOKEN") == "" {
		log.Fatal("Empty token. Set token via \"export TOKEN='<>'\"")
	}

	b, err := tele.NewBot(
		tele.Settings{
			Token:  os.Getenv("TOKEN"),
			Poller: &tele.LongPoller{Timeout: 10 * time.Second}})

	conf := GetConf()
	queue := Queue{}

	allowedGroup := b.Group()
	allowedGroup.Use(GroupsWhitelist(conf.Restrictions.Group))
	if err != nil {
		log.Fatal(err)
		return
	}

	MakeGeneralHandlers(b)
	MakeGroupHandlers(allowedGroup, &queue) // more privileged
	b.Start()
}
