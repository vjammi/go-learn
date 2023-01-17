package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Print(result)
}

func twoSum(nums []int, target int) []int {
	//mapp := make(map[int]int)
	//var mapp map[int]int = make(map[int]int)
	var mapp = make(map[int]int)

	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		_, containsKey := mapp[complement]
		if containsKey {
			return []int{i, mapp[complement]}
		}
		mapp[nums[i]] = i
	}
	return []int{}
}
