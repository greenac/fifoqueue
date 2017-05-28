package main

import (
	"github.com/greenac/fifoqueue/queue"
	"fmt"
)

type QueueValue struct {
	Value string
}

func main() {
	const size = 10
	text := "Besty Buddy By The Sea In May Is Jay"
	queue := fifoqueue.FifoQueue{}
	for i := 0; i < size; i += 1 {
		v := QueueValue{Value: text[:i]}
		queue.Insert(v)
	}

	fmt.Println(queue.Values())

	for i := 0; i < size; i += 1 {
		value, err := queue.Pop()
		if err == nil {
			println("Popping:", value)
			fmt.Println(queue.Values())
		} else {
			fmt.Println("Failed to pop from queue with error:", err)
		}
	}
}
