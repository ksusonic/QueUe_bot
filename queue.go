package main

import (
	"errors"
	"gopkg.in/telebot.v3"
)

type QueueMember struct {
	Username  string   `json:"username"`
	Usernames []string `json:"usernames"`
	Message   string   `json:"message"`
}

type Queue struct {
	Chat    telebot.ChatID `json:"chat_id"`
	Members []QueueMember  `json:"members"`
}

func (q *Queue) Len() int {
	return len(q.Members)
}

func (q *Queue) Swap(i, j int) {
	q.Members[i], q.Members[j] = q.Members[j], q.Members[i]
}

func (q *Queue) Push(x QueueMember) {
	q.Members = append(q.Members, x)
}
func (q *Queue) Pop() (QueueMember, error) {
	if len(q.Members) <= 1 {
		return QueueMember{}, errors.New("empty queue")
	}
	oldMembers := q.Members
	n := len(oldMembers)
	x := oldMembers[n-1]
	q.Members = oldMembers[0 : n-1]
	return x, nil
}
