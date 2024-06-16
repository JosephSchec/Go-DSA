package queue

import (
	"errors"
	"fmt"
)

/*
 * Pushing and popping from a queue are constant time operations since it doesn't mattere the length
 */
type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type Queue[T any] struct {
	Length int
	Head   *Node[T]
	Tail   *Node[T]
}

func (q *Queue[T]) Enqueue(value T) *Queue[T] {
	node := &Node[T]{Value: value}
	q.Length++
	if q.Tail != nil {
		q.Tail.Next = node
	}
	q.Tail = node
	if q.Head == nil {
		q.Head = node
	}

	return q
}

func (q *Queue[T]) Dequeue() T {

	if q.Head == nil {
		q.Tail = nil
	}
	prevHead := q.Head
	q.Head = prevHead.Next
	q.Length--

	return prevHead.Value

}

func (q *Queue[T]) Peek() (T, error) {
	if q == nil {
		var zero T
		return zero, errors.New("queue is empty")
	}
	return q.Head.Value, nil
}

func (q *Queue[T]) DisplayAllNodes() {

	if q == nil {
		return
	}
	list := []T{}
	current := q.Head
	for current != nil {
		list = append(list, current.Value)
		current = current.Next
	}

	fmt.Printf("%v\n", list)

}
