package main

import (
	"github.com/asaskevich/govalidator"
	tele "gopkg.in/telebot.v3"
	"strconv"
	"strings"
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

func MakeGroupHandlers(b *tele.Bot, q *Queue) {
	b.Handle("/push", func(c tele.Context) error {
		payload := c.Message().Payload
		if len(payload) == 0 {
			username := c.Sender().Username
			q.Push(QueueMember{Usernames: []string{username}})
			currentLen := strconv.Itoa(q.Len())
			return c.Send("@" + username + " " + currentLen + " в очереди.")
		} else {
			splitted := strings.Split(payload, " ")
			var message string
			var usernames []string
			var addMyself = true
			for _, w := range splitted {
				if w[0] == '@' && len(w) > 1 {
					usernames = append(usernames, w[1:])
				} else if w == "#nome" {
					addMyself = false
				} else {
					message += w
				}
			}
			if addMyself {
				usernames = append(usernames, c.Sender().Username)
			}
			member := QueueMember{
				Usernames: usernames,
				Message:   message,
			}
			q.Push(member)
			return c.Send(member.UsernamesString() + " " + strconv.Itoa(q.Len()) + "й в очереди" +
				func(s string) string {
					if len(s) > 0 {
						return " с темой: " + s
					} else {
						return ""
					}
				}(message))
		}
	})
	b.Handle("/pop", func(c tele.Context) error {
		if q.Empty() {
			return c.Send("Очередь пуста.")
		}
		member, err := q.Pop(c.Sender().Username)
		if err != nil {
			return c.Send("@" + c.Sender().Username + " не стоял в очереди -_-")
		}
		return c.Send("@" + member.Usernames[0] + member.Message + " вышел из очереди.") // TODO for users
	})
	b.Handle("/queue", func(c tele.Context) error {
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
	b.Handle("/swap", func(c tele.Context) error {
		if len(c.Message().Payload) == 0 {
			return c.Send("Чтобы поменяться местами, введи номера в очереди, например: /swap 1 2")
		}
		if q.Len() == 0 {
			return c.Send("Очередь пуста 🍻")
		}
		splitted := strings.Split(c.Message().Payload, " ")
		if len(splitted) >= 2 && govalidator.IsInt(splitted[0]) && govalidator.IsInt(splitted[1]) {
			l, _ := strconv.ParseInt(splitted[0], 10, 32)
			r, _ := strconv.ParseInt(splitted[1], 10, 32)
			err := q.Swap(int(l), int(r))
			if err == nil {
				return c.Send(q.Members[r-1].UsernamesString() + " и " + q.Members[l-1].UsernamesString() + " поменялись местами")
			}
		}
		return c.Send("Не могу этого сделать :/")
	})
}
