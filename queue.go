package main

// In-memory queue //

import (
	"errors"
	"fmt"
	"gopkg.in/telebot.v3"
	"strconv"
	"strings"
)

type QueueMember struct {
	Usernames []string `json:"usernames"`
	Message   string   `json:"message"`
}

type Queue struct {
	Id      uint32         `json:"queue_id"`
	Title   string         `json:"title"`
	Chat    telebot.ChatID `json:"chat_id"`
	Members []QueueMember  `json:"members"`
}

func (q *Queue) Len() int {
	return len(q.Members)
}

func (q *Queue) DebugString() string {
	var sb strings.Builder

	if len(q.Title) > 0 {
		sb.WriteString("Title: " + q.Title + "\n\n")
	} else {
		sb.WriteString("Untitled queue\n")
	}

	for qi, m := range q.Members {
		sb.WriteString("Members N" + strconv.Itoa(qi+1) + "\n")
		for i, user := range m.Usernames {
			sb.WriteString("  " + strconv.Itoa(i+1) + " @" + user + "\n")
		}
		if len(m.Message) != 0 {
			sb.WriteString("with message: " + m.Message + "\n\n")
		}
	}

	return sb.String()
}

func (q *Queue) Swap(ipos, jpos int) (bool, error) {
	i, j := ipos-1, jpos-1
	if i < 0 || j < 0 {
		return false, errors.New("one of indexes <= 0")
	}
	if q.Len() <= i || q.Len() <= j {
		return false, errors.New("one of indexes is out of range")
	}

	q.Members[i], q.Members[j] = q.Members[j], q.Members[i]
	return true, nil
}

func (q *Queue) Push(x QueueMember) {
	q.Members = append(q.Members, x)
}

func (q *Queue) GetQueuePos(luser string) int {
	// Returns position of user(s) in queue if exists. Else -1
	for i, member := range q.Members {
		for _, ruser := range member.Usernames {
			if luser == ruser {
				return i + 1
			}
		}
	}
	return -1
}

func (q *Queue) Pop(username string) (QueueMember, error) {
	pos := q.GetQueuePos(username)
	if pos != -1 {
		posIndex := pos - 1
		member := q.Members[posIndex]
		q.Members = append(q.Members[:posIndex], q.Members[posIndex+1:]...)
		return member, nil
	} else {
		return QueueMember{}, fmt.Errorf("no such member in queue")
	}
}
