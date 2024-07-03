package main

import (
	"sort"
	"math"
)

func minDifference(nums []int) int {
	if len(nums)<=4 {
		return 0
	}

    sort.Ints(nums)
	var minVal int=math.MaxInt

	minVal = min(minVal, nums[len(nums)-4] - nums[0]);
    minVal = min(minVal, nums[len(nums)-1] - nums[3]);
    minVal = min(minVal, nums[len(nums)-2] - nums[2]);
    minVal = min(minVal, nums[len(nums)-3] - nums[1]);

	return minVal

}