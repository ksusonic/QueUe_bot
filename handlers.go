package main

import (
	tele "gopkg.in/telebot.v3"
	"strconv"
)

func MakeGeneralHandlers(b *tele.Bot) {
	b.Handle("/ping", func(c tele.Context) error {
		return c.Send("pong 🏓")
	})
	b.Handle("/pwd", func(c tele.Context) error {
		prefix := "Chat id: "
		id := strconv.FormatInt(c.Chat().ID, 10)
		return c.Send(prefix+id, &tele.SendOptions{
			Entities: tele.Entities{tele.MessageEntity{Type: tele.EntityCode, Offset: len(prefix), Length: len(id)}}})
	})
}

func MakeGroupHandlers(g *tele.Group, q *Queue) {
	g.Handle("/push", func(c tele.Context) error {
		payload := c.Message().Payload
		if payload == "" {
			username := c.Sender().Username
			q.Push(QueueMember{Usernames: []string{username}})
			currentLen := strconv.Itoa(q.Len())
			return c.Send("@" + username + " " + currentLen + "й в очереди.")
		}
		return c.Send("Payload for push in dev")
	})
	g.Handle("/pop", func(c tele.Context) error {
		member, err := q.Pop(c.Sender().Username)
		if err != nil {
			return c.Send("Очередь пуста.")
		}
		return c.Send("@" + member.Usernames[0] + member.Message + " вышел из очереди.") // TODO for users
	})
	g.Handle("/queue", func(c tele.Context) error {
		if q.Len() == 0 {
			return c.Send("Очередь пуста 🍻")
		}
		var message string
		for member := range q.Members {
			message += strconv.Itoa(member+1) + ": @" + q.Members[member].Usernames[0] // TODO for users
			if q.Members[member].Message != "" {
				message += " c " + q.Members[member].Message
			}
			message += "\n"
		}
		return c.Send(message, tele.Silent)
	})
}
