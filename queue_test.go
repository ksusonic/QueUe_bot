package main

import (
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
