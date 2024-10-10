package problem4

import "fmt"

type Element[T any] struct {
	value T
	next  *Element[T]
}

type LinkedList[T any] struct {
	head *Element[T]
	size int
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (ll *LinkedList[T]) Add(element *Element[T]) {
	if ll.head == nil {
		ll.head = element
	} else {
		current := ll.head
		for current.next != nil {
			current = current.next
		}
		current.next = element
	}
	ll.size++
}

func (ll *LinkedList[T]) Insert(element *Element[T], position int) error {
	if position < 0 || position > ll.size {
		return fmt.Errorf("index out of range")
	}
	if position == 0 {
		element.next = ll.head
		ll.head = element
	} else {
		current := ll.head
		for i := 0; i < position-1; i++ {
			current = current.next
		}
		element.next = current.next
		current.next = element
	}
	ll.size++
	return nil
}

func (ll *LinkedList[T]) Delete(element *Element[T]) error {
	if ll.head == nil {
		return fmt.Errorf("list is empty")
	}
	if ll.head.value == element.value {
		ll.head = ll.head.next
		ll.size--
		return nil
	}
	current := ll.head
	for current.next != nil && current.next.value != element.value {
		current = current.next
	}
	if current.next == nil {
		return fmt.Errorf("element not found")
	}
	current.next = current.next.next
	ll.size--
	return nil
}

func (ll *LinkedList[T]) Find(value T) (*Element[T], error) {
	current := ll.head
	for current != nil {
		if current.value == value {
			return current, nil
		}
		current = current.next
	}
	return nil, fmt.Errorf("element not found")
}

func (ll *LinkedList[T]) List() []*Element[T] {
	result := []*Element[T]{}
	current := ll.head
	for current != nil {
		result = append(result, current)
		current = current.next
	}
	return result
}

func (ll *LinkedList[T]) Size() int {
	return ll.size
}

func (ll *LinkedList[T]) IsEmpty() bool {
	return ll.size == 0
}
