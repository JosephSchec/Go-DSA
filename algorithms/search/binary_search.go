package search

import (
	"fmt"
)

// Only works on sorted arr
func BinarySearch(props SearchProps) bool {
	haystack := props.Haystack
	needle := props.Needle
	low := 0
	high := len(haystack)
	for low < high {
		midpoint := int(float64(low + (high-low)/2))

		val := haystack[midpoint]
		fmt.Println("low:", low, "high:", high, "mid", midpoint, needle, val)
		if needle == val {
			return true
		} else if val < needle {
			low = midpoint + 1
		} else {
			high = midpoint
		}
	}
	return false
}

func RunSampleBinarySearch() {
	props := SearchProps{
		Haystack: []int{1, 2, 3, 4, 5, 6},
		Needle:   6,
	}
	fmt.Println("found binary search", BinarySearch(props))
}
