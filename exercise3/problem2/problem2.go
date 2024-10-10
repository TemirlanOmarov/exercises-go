package problem2

import (
	"errors"
	"fmt"
)

type Stack struct {
	elements []interface{}
}

func (s *Stack) Push(value interface{}) {
	s.elements = append(s.elements, value)
}

func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	top := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return top, nil
}

func (s *Stack) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	return s.elements[len(s.elements)-1], nil
}

func (s *Stack) Size() int {
	return len(s.elements)
}

func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}

func main() {
	stack := Stack{}

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	peek, _ := stack.Peek()
	fmt.Println("Peek:", peek)

	popped, _ := stack.Pop()
	fmt.Println("Popped:", popped)

	fmt.Println("Size:", stack.Size())

	fmt.Println("IsEmpty:", stack.IsEmpty())
}
