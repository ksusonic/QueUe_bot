package main

import (
	tele "gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/middleware"
	"log"
	"os"
	"time"
)

func main() {
	b, err := tele.NewBot(
		tele.Settings{
			Token:  os.Getenv("TOKEN"),
			Poller: &tele.LongPoller{Timeout: 10 * time.Second}})
	b.Use(middleware.Logger())

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
