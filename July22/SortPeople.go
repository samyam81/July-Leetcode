package main

import "fmt"

func sortPeople(names []string, heights []int) []string {
	heightSlice := make([]heightS, len(heights), len(heights))
	for i, val := range heights {
		heightSlice[i] = heightS{index: i, height: val}
	}

	result := make([]string, len(names), len(names))
	sorted := qsort(heightSlice)
	fmt.Println(sorted)
	for i, val := range sorted {
		result[i] = names[val.index]
	}
	return result
}

type heightS struct {
	index  int
	height int
}

func qsort(vals []heightS) []heightS {
	if len(vals) <= 1 {
		return vals
	} else {
		var less, more []heightS
		result := make([]heightS, 0, len(vals))
		target := vals[len(vals)/2]
		for _, val := range vals {
			if val.index == target.index {
				continue
			}
			if val.height > target.height {
				less = append(less, val)
			} else {
				more = append(more, val)
			}
		}
		result = append(result, qsort(less)...)
		result = append(result, target)
		result = append(result, qsort(more)...)
		return result
	}
}
