// queue.go
// Generates a FIFO queue structure

package solver

import "lem-in/types"

// Queue (FIFO) structure containing a list of rooms and a size.
type Queue struct {
	items []types.Room
	size  int
}

// GenQueue initializes a new queue of a given length
func GenQueue(length int) *Queue {
	return &Queue{
		items: make([]types.Room, length),
		size:  length,
	}
}

// GetSize returns the current size of the queue.
func (q *Queue) GetSize() int {
	return q.size
}

// Enqueue adds an item to the end of the queue.
func (q *Queue) Enqueue(room types.Room) {
	q.size++
	// Create a new queue with the new size and copy the elements over.
	newq := GenQueue(q.size)
	copy(newq.items, q.items)
	q.items = newq.items
	// Insert the room into the last spot of the queue.
	q.items[q.size-1] = room
}

// Dequeue returns an item from the first spot in the queue, deletes
// it from the queue, and reorganizes the remaining items without disrupting
// the queue's order.
func (q *Queue) Dequeue() types.Room {
	// Return a nil room if the queue is empty.
	if q.size == 0 {
		return types.Room{}
	}
	room := q.items[0]
	// Slice off the first spot in the queue.
	q.items = q.items[1:]
	// Decrement the size of the queue.
	q.size--
	return room
}
