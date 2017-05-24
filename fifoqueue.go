package main

import (
	"github.com/greenac/fifoqueue/queue"
	"fmt"
)

func main() {
	queue := fifoqueue.FifoQueue{}
	for i := 0; i < 10; i += 1 {
		queue.Insert(i)
	}

	fmt.Println(queue.Values())

	for i := 0; i < 10; i += 1 {
		value, err := queue.Pop()
		if err == nil {
			println("Popping:", value)
			fmt.Println(queue.Values())
		} else {
			fmt.Println("Failed to pop from queue with error:", err)
		}
	}

	for i := 0; i < 20; i += 1 {
		queue.Insert(i)
	}

	fmt.Println(queue.Values())
}
