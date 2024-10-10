package problem1

import (
	"errors"
	"fmt"
)

type Queue struct {
	elements []interface{}
}

func (q *Queue) Enqueue(value interface{}) {
	q.elements = append(q.elements, value)
}

func (q *Queue) Dequeue() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}
	first := q.elements[0]
	q.elements = q.elements[1:]
	return first, nil
}

func (q *Queue) Peek() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}
	return q.elements[0], nil
}

func (q *Queue) Size() int {
	return len(q.elements)
}

func (q *Queue) IsEmpty() bool {
	return len(q.elements) == 0
}

func main() {
	queue := Queue{}

	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)

	peek, _ := queue.Peek()
	fmt.Println("Peek:", peek)
	dequeued, _ := queue.Dequeue()
	fmt.Println("Dequeued:", dequeued)

	fmt.Println("Size:", queue.Size())

	fmt.Println("IsEmpty:", queue.IsEmpty())
}
