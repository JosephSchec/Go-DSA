package main

import (
	"fmt"

	"github.com/josephschec/go-dsa/algorithms/queue"
	"github.com/josephschec/go-dsa/algorithms/search"
	"github.com/josephschec/go-dsa/algorithms/sort"
)

func main() {
	props := search.SearchProps{
		Haystack: []int{1, 2, 3, 4, 5},
		Needle:   6,
	}
	fmt.Println("found linear search", search.LinearSearch(props))
	fmt.Println("found binary search", search.BinarySearch(props))
	fmt.Println("found crystal", search.TwoCrystalBallsSearch([]bool{false, false, false, false, true, true, true}))
	myArr := []int{1, 5, 252, 34, 4}
	sort.BubbleSort(&myArr)
	myQueue := &queue.Queue[int]{}

	myQueue.Enqueue(1)
	myQueue.Enqueue(2)
	myQueue.Enqueue(3)
	myQueue.DisplayAllNodes()

	myQueue.Dequeue()
	myQueue.Enqueue(4)
	myQueue.DisplayAllNodes()

	// myQueue.Dequeue()
	peekValue, err := myQueue.Peek()

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Peek:", peekValue)
	}

}
