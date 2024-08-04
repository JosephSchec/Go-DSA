package main

import (
	"github.com/josephschec/go-dsa/algorithms/sort"
)

// "github.com/josephschec/go-dsa/algorithms/sort"

func main() {

	myArr := []int{1, 5, 252, 34, 4}
	// sort.BubbleSort(&myArr)
	sort.QuickSort(&myArr)

	/** SEARCH START **/

	// search.RunSampleLinearSearch()
	// search.RunSampleBinarySearch()
	// search.RunSampleCBSearch()

	/** SEARCH END **/
	/** LINKED LIST START **/
	// linked_list.RunSampleQueue()

	// linked_list.RunSampleStack()
	/** LINKED LIST END **/

	/**RECURSION*/
	// recursion.RunSampleMaze()
}
