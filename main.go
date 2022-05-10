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

	queue := Queue{}

	if err != nil {
		log.Fatal(err)
		return
	}

	MakeGeneralHandlers(b)
	MakeGroupHandlers(b, &queue)
	b.Start()
}
