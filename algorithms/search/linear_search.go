package search

import "fmt"

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

func RunSampleLinearSearch() {
	props := SearchProps{
		Haystack: []int{1, 2, 3, 4, 5, 6},
		Needle:   6,
	}
	fmt.Println("found linear search", LinearSearch(props))
}
