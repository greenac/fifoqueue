package fifoqueue

import (
	"errors"
	"github.com/greenac/gologger"
)

var logger = gologger.GoLogger{}

type QueueNode struct {
	Payload interface{}
	Next *QueueNode
	Prev *QueueNode
}

func (n *QueueNode) print() {
	logger.Log("node:", n, "prev:", n.Prev, "next:", n.Next)
}

type FifoQueue struct {
	head *QueueNode
	tail *QueueNode
	length int
}

func (q *FifoQueue)Insert(payload interface{}) *QueueNode {
	node := QueueNode{Payload: payload}
	q.length += 1

	logger.Warn("In fifo queue insert function")
	if q.head == nil {
		q.head = &node
		q.tail = &node
	} else {
		if q.tail == nil {
			logger.Error("The tail is nil!!! holy shit")
			logger.Error("Meanwhile the head is:", q.head)
		}

		tempTailPtr := q.tail
		tempTailPtr.Next = &node
		node.Prev = tempTailPtr
		q.tail = &node
	}

	return &node
}

func (q *FifoQueue)Pop() (*QueueNode, error) {
	if q.tail == nil {
		return nil, errors.New("NoElementsInQueue")
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
	return target, nil
}

func (q *FifoQueue)PopMany(count int) *[]QueueNode {
	var nodes []QueueNode
	if q.IsEmpty() {
		nodes = make([]QueueNode, 0)
		return &nodes
	}

	size := count
	if q.Length() < count {
		size = q.Length()
	}

	nodes = make([]QueueNode, size)
	for i := 0; i < size; i += 1 {
		n, err := q.Pop()
		if err != nil {
			logger.Warn("Error popping many:", err)
		}

		nodes[size - i -1] = *n
	}

	return &nodes
}

func (q *FifoQueue)GetManyPayloads(count int) *[]interface{} {
	nodes := q.PopMany(count)
	payloads := make([]interface{}, len(*nodes))
	for i, n := range *nodes {
		payloads[i] = n.Payload
	}

	return &payloads
}

func (q *FifoQueue) GetPayloads(count int) *[]interface{} {
	var c int
	if count > q.Length() {
		c = q.Length()
	} else {
		c = count
	}

	pls := make([]interface{}, c)
	i := 0
	node := q.head
	for i < c && node != nil {
		pls[i] = node.Payload
		node = node.Next
		i += 1
	}

	return &pls
}

func (q *FifoQueue)Length() int {
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

func (q *FifoQueue) AsSlice() *[]*QueueNode {
	nodes := make([]*QueueNode, 0)
	node := q.head
	for node != nil {
		nodes = append(nodes, node)
		node = node.Next
	}

	return &nodes
}

func (q *FifoQueue) Delete(n *QueueNode) bool {
	node := q.head
	for node != nil {
		if n == node {
			break
		}

		node = node.Next
	}
	
	if node == nil {
		return false
	}

	if node == q.head && node == q.tail {
		q.head = nil
		q.tail = nil
	} else if node == q.head {
		node.Next.Prev = nil
		q.head = node.Next
		node.Next = nil
	} else if node == q.tail {
		node.Prev.Next = nil
		q.tail = node.Prev
		node.Prev = nil
	} else {
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev
	}

	node = nil
	q.length -= 1

	return true
}
