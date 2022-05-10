package main

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

var (
	q1 = QueueMember{
		Usernames: []string{"testUser1"},
		Message:   "TestMessage1",
	}
	q2 = QueueMember{
		Usernames: []string{"testUser2"},
		Message:   "TestMessage2",
	}
	q3 = QueueMember{
		Usernames: []string{"testUser2", "testUser3"},
		Message:   "TestMessage3",
	}
)

func AssertQueue(t *testing.T, q Queue, members ...QueueMember) {
	for i := range members {
		assert.Equal(t, q.Members[i], members[i])
	}
}

func TestQueue_Push_Len(t *testing.T) {
	q := Queue{}
	q.Push(q1)
	if q.Len() != 1 {
		t.Fatal("Expected 1, got ", q.Len())
	}
	q.Push(q2)
	if q.Len() != 2 {
		t.Fatal("Expected 2, got ", q.Len())
	}
}

func TestQueue_Swap(t *testing.T) {
	q := Queue{}
	q.Push(q1)
	q.Push(q2)
	b, mess := q.Swap(1, 2)
	if b == false || mess != nil {
		t.Fatal("Expected OK, got message ", mess)
	}
	if q.Members[0].Usernames[0] != q2.Usernames[0] {
		t.Fatal("Expected to swap 1 2, got incorrect", q.DebugString())
	}

	b, mess = q.Swap(2, 3)
	if b != false || mess == nil {
		t.Fatal("Expected NOT OK, got success. Queue: ", q.DebugString())
	}
}

func TestQueue_GetQueuePos(t *testing.T) {
	q := Queue{}
	q.Push(q1)
	q.Push(q2)
	q.Push(q3)
	if q.GetQueuePos("testUser2") != 2 {
		t.Fatal("Expected 2, got: ", q.GetQueuePos("testUser2"))
	}
}

func TestQueue_Pop(t *testing.T) {
	q := Queue{Title: "Test Pop"}
	q.Push(q1)
	q.Push(q2)
	q.Push(q3)
	user, err := q.Pop("testUser2")
	if len(user.Usernames) == 0 || err != nil || q.Len() != 2 {
		t.Fatal("Expected to pop, got: ", err)
	}
	AssertQueue(t, q, q1, q3)

	user, err = q.Pop("testUser1")
	if len(user.Usernames) == 0 || err != nil || q.Len() != 1 {
		t.Fatal("Expected to pop, got: ", err)
	}
	AssertQueue(t, q, q3)

	user, err = q.Pop("testUser1")
	if len(user.Usernames) != 0 || err == nil {
		t.Fatal("Expected NOT to pop, got: ", err)
	}
	assert.NotEqual(t, user, q1)
	AssertQueue(t, q, q3)
}
