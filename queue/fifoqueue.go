package fifoqueue

import (
	"errors"
)

type QueueNode struct {
	Payload interface{}
	Next *QueueNode
	Prev *QueueNode
}

type FifoQueue struct {
	head *QueueNode
	tail *QueueNode
	length uint
}

func (q *FifoQueue)Insert(payload interface{}) {
	var node QueueNode
	node.Payload = payload
	q.length += 1

	if q.head == nil {
		q.head = &node
		q.tail = &node
	} else {
		tempTailPtr := q.tail
		tempTailPtr.Next = &node
		node.Prev = tempTailPtr
		q.tail = &node
	}
}

func (q *FifoQueue)Pop() (interface{}, error) {
	var value int
	if q.tail == nil {
		return value, errors.New("NoElementsInQueue")
	}

	target := q.tail
	if q.tail == q.head {
		q.tail = nil
		q.head = nil
	} else {
		prev := q.tail.Prev
		prev.Next = nil
		target.Prev = nil
		q.tail = prev
	}

	q.length -= 1
	return target.Payload, nil
}

func (q *FifoQueue)Length() uint {
	return q.length
}

func (q *FifoQueue) Values() []interface{} {
	var values []interface{}
	if q.head == nil {
		return values
	}

	node := q.head
	for node != nil {
		values = append(values, node.Payload)
		node = node.Next
	}

	return values
}

func (q *FifoQueue)IsEmpty() bool {
	return q.length == 0
}
