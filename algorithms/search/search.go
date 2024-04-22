package search

import "math"

type SearchProps struct {
	Haystack []int
	Needle   int
}

func LinearSearch(props SearchProps) bool {

	for i := 0; i < len(props.Haystack); i++ {
		if props.Haystack[i] == props.Needle {
			return true
		}

	}
	return false
}

// Only works on sorted arr
func BinarySearch(props SearchProps) bool {
	haystack := props.Haystack
	needle := props.Needle
	low := 0
	high := len(haystack)

	for low < high {
		midpoint := int(math.Floor(float64(low + (high-low)/2)))
		val := haystack[midpoint]
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

/**
Goal: we have 2 crystal balls and we want to find the max floor we can drop from without breaking;
This is a sorted array of falses until true and then contiunouse true
If we were to cut in half (binary search) then at most we would need to walk from the beginning to half point ;
But if we continuously jump by sqrt then at most we would walk from the last square root to the next square root
**/
func TwoCrystalBallsSearch(floors []bool) int {
	jumpAmount := int(math.Floor(math.Sqrt(float64(len(floors)))))
	i := jumpAmount
	for ; i < len(floors); i += jumpAmount {
		if floors[i] {
			break
		}
	}

	i -= jumpAmount

	for lastSafePoint := jumpAmount; lastSafePoint < len(floors); lastSafePoint++ {

		if floors[i] {
			return i
		}
		i++
	}
	return -1
}
