package main

import (
	tele "gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/middleware"
)

func RestrictGroups(v middleware.RestrictConfig) tele.MiddlewareFunc {
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		if v.In == nil {
			v.In = next
		}
		if v.Out == nil {
			v.Out = next
		}
		return func(c tele.Context) error {
			for _, chat := range v.Chats {
				if chat == c.Chat().ID {
					return v.In(c)
				}
			}
			return v.Out(c)
		}
	}
}

func GroupsWhitelist(groupIds ...int64) tele.MiddlewareFunc {
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return RestrictGroups(middleware.RestrictConfig{
			Chats: groupIds,
			In:    next,
			Out:   func(c tele.Context) error { return nil },
		})(next)
	}
}
