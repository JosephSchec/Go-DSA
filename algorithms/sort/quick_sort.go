package sort

import (
	"fmt"
)

func getPartition(arr []int, low int, high int) (int, []int) {

	pivot := arr[high]
	idx := low - 1
	/*
	* [6,4,9,7,8,5]
	* in this scenario, pivot is 5 , we check at 6 are you les than 5 no ok , 4 are you less than 5, yes
	* so we swap you to the first index so after one iteration our array looks like this
	* [4,6,9,7,8,5]
	 */
	for i := low; i < high; i++ {

		if arr[i] <= pivot {
			idx++
			temp := (arr)[i]
			(arr)[i] = arr[idx]
			arr[idx] = temp
		}
	}
	/*
	 * Then before we return we need to tak our pivot ant place it correctly in the array
	 * so our array after 1 iteration should look like [4,5,6,9,7,8]
	 */
	idx++
	arr[high] = arr[idx]
	arr[idx] = pivot
	return idx, arr

}

func performQs(arr []int, low int, high int) {
	if low >= high {
		return
	}
	pivotIdx, _ := getPartition(arr, low, high)
	// sort left side of pivot (anything less than the pivot)
	performQs(arr, low, pivotIdx-1)

	// sort right side of pivot (anything more than the pivot still needs to be sorted)
	performQs(arr, pivotIdx+1, high)

}

// O(n Log(n)) - > O(n^2)
func QuickSort(arr *[]int) {

	fmt.Println("og ", arr)

	performQs(*arr, 0, len(*arr)-1)

	fmt.Println("after", arr)
}
