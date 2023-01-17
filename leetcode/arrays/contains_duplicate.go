package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 1}
	result := containsDuplicate(nums)
	fmt.Print(result)
}

func containsDuplicate(nums []int) bool {

	var set = make(map[int]int)
	//set := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		_, contains := set[nums[i]]
		if contains {
			return true
		}

		if !contains {
			set[nums[i]] = nums[i]
		}
	}

	return false
}
