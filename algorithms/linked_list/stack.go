package linked_list

import (
	"errors"
	"fmt"
)

/*
 * Pushing and popping from a Stack are constant time operations since it doesn't matter the length
 */

type StackNode[T any] struct {
	Value T
	Prev  *StackNode[T]
}

type Stack[T any] struct {
	Length int
	Head   *StackNode[T]
	Tail   *StackNode[T]
}

func (q *Stack[T]) Push(value T) *Stack[T] {
	node := &StackNode[T]{Value: value}
	q.Length++
	if q.Head == nil {
		q.Head = node
		return q
	}

	node.Prev = q.Head

	q.Head = node

	return q
}

func (q *Stack[T]) Pop() T {
	q.Length = max(0, q.Length-1)
	curHead := q.Head
	if q.Head == nil || q.Length == 0 {
		q.Head = nil
		q.Tail = nil
		return curHead.Value
	}

	q.Head = curHead.Prev
	q.Length--

	return curHead.Value

}

func (q *Stack[T]) PeekStack() (T, error) {
	if q == nil {
		var zero T
		return zero, errors.New("Stack is empty")
	}
	return q.Head.Value, nil
}

func (q *Stack[T]) DisplayStackNodes() {

	if q == nil {
		return
	}
	list := []T{}
	current := q.Head
	for current != nil {
		list = append(list, current.Value)
		current = current.Prev
	}

	fmt.Printf("%v\n", list)

}

func RunSampleStack() {
	myStack := Stack[int]{}
	myStack.Push(2)
	myStack.Push(3)
	myStack.Push(4)
	myStack.DisplayStackNodes()
	fmt.Println("De-stack", myStack.Pop())
	myStack.DisplayStackNodes()

	fmt.Println("De-stack", myStack.Pop())
	myStack.DisplayStackNodes()

}
