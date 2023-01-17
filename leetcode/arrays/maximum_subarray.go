package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	maxSubArrSum := maxSubArray(nums)
	fmt.Print(maxSubArrSum)
}

func maxSubArray(nums []int) int {
	var maxSubArrSum = nums[0]
	var subArrSumAtKMinus1 = nums[0]
	for k := 1; k < len(nums); k++ {
		var maxSumAtKth = 0
		if subArrSumAtKMinus1+nums[k] > nums[k] {
			maxSumAtKth = subArrSumAtKMinus1 + nums[k]
		} else {
			maxSumAtKth = nums[k]
		}
		maxSubArrSum = int(math.Max(float64(maxSubArrSum), float64(maxSumAtKth)))
		subArrSumAtKMinus1 = maxSumAtKth
	}
	return maxSubArrSum
}
