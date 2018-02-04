package fifoqueue

import "testing"

func TestFifoQueue_Insert(t *testing.T) {
	const input = 5
	q := FifoQueue{}

	q.Insert(input)
	if q.length != 1 {
		t.Fatalf("Fifo Queue expected inserted %d. Length is %d, but length 1 was expected.", input, q.length)
	}
}

func TestFifoQueue_Pop(t *testing.T) {
	const input = 5
	q := FifoQueue{}

	_, err := q.Pop()
	if err == nil {
		t.Fatal("No error was thrown when popping from empty queue.")
	}

	q.Insert(input)

	v, err := q.Pop()
	if err != nil {
		t.Fatal("Failed to pop from queue. Error was thrown", err)
	}

	if v != input {
		t.Fatalf("Popping from queue should have returned: %d but the value: %d was returned", input, v)
	}

	if q.length != 0 {
		t.Fatalf("Queue length should be zero after popping, but the length is: %d", q.length)
	}
}

func TestFifoQueue_Length(t *testing.T) {
	const input = 5
	q := FifoQueue{}

	q.Insert(input)

	if q.Length() != 1 {
		t.Fatalf("Queue should have a length of 1, but it's length is: %d", q.Length())
	}
}
