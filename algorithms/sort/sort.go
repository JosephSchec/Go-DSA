package sort

import "fmt"

// O(n^2)
func BubbleSort(arr *[]int) {
	fmt.Println("og ", arr)

	for i := 0; i < len(*arr); i++ {

		for j := 0; j < len(*arr)-i-1; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				temp := (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = temp
			}
		}

	}
	fmt.Println("after", arr)
}
