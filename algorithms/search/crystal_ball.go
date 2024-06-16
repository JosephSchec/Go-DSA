package search

import (
	"fmt"
	"math"
)

/*
*
Goal: we have 2 crystal balls and we want to find the max floor we can drop from without breaking;
This is a sorted array of falses until true and then contiunouse true
If we were to cut in half (binary search) then at most we would need to walk from the beginning to half point ;
But if we continuously jump by sqrt then at most we would walk from the last square root to the next square root
*
*/
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

func RunSampleCBSearch() {
	fmt.Println("found crystal", TwoCrystalBallsSearch([]bool{false, false, false, false, false, true, true, true}))
}
