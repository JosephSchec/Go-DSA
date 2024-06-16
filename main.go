package main

import (
	"github.com/josephschec/go-dsa/algorithms/linked_list"
	// "github.com/josephschec/go-dsa/algorithms/search"
	// "github.com/josephschec/go-dsa/algorithms/sort"
)

func main() {
	// props := search.SearchProps{
	// 	Haystack: []int{1, 2, 3, 4, 5, 6},
	// 	Needle:   6,
	// }
	// fmt.Println("found linear search", search.LinearSearch(props))
	// fmt.Println("found binary search", search.BinarySearch(props))
	// fmt.Println("found crystal", search.TwoCrystalBallsSearch([]bool{false, false, false, false, true, true, true}))
	// myArr := []int{1, 5, 252, 34, 4}
	// sort.BubbleSort(&myArr)

	linked_list.RunSampleQueue()

	linked_list.RunSampleStack()

}
