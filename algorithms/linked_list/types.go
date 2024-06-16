package linked_list

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type Queue[T any] struct {
	Length int
	Head   *Node[T]
	Tail   *Node[T]
}

type StackNode[T any] struct {
	Value T
	Prev  *StackNode[T]
}

type Stack[T any] struct {
	Length int
	Head   *StackNode[T]
	Tail   *StackNode[T]
}
