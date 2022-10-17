package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("target not found in array")

func main() {
	//var arr = []int{20, 32, 54, 78, 82, 94, 97}
	arr := []int{20, 32, 54, 78, 82, 94, 97}

	index1, err := binarySearchRecursive(arr, 82, 0, 6)
	fmt.Println(index1, err)

	index2, err := binarySearchIterative(arr, 82)
	fmt.Println(index2, err)
}

// Binary search for target within a sorted array by repeatedly dividing the array in half and comparing the midpoint with the target.
// This function uses recursive call to itself.
// if a target is found
//
//	the index of the target is returned.
//
// else
//
//	returns -1 and ErrNotFound.
func binarySearchRecursive(arr []int, target int, low int, high int) (int, error) {

	if high < low || len(arr) == 0 {
		return -1, ErrNotFound
	}

	mid := int(low + (high-low)/2)

	if arr[mid] > target {
		return binarySearchRecursive(arr, target, low, mid-1)
	} else if arr[mid] < target {
		return binarySearchRecursive(arr, target, mid+1, high)
	} else {
		return mid, nil
	}
}

func binarySearchIterative(arr []int, target int) (int, error) {
	low := 0
	high := len(arr)

	if low > high {
		return -1, nil
	}

	for low <= high {
		mid := int(low + (high-low)/2)

		if target < arr[mid] {
			high = mid - 1
		} else if target > arr[mid] {
			low = mid + 1
		} else { //arr[mid] == target
			return mid, nil
		}
	}
	return -1, nil
}

//         num     33
//         nums   [44, 55, 66, 77, 88, 99, 1, 11, 22, 33  33]
//            i    0   1   2   3   4   5   6  7   8   9   10
//                 ^                                      ^
//          mid                        ^
//                                                ^
//                                                    ^
//                                                 33 found
//        0+(11-0)/2   = 0+11/2   = 5
//        6+(10-6)/2   = 6+4/2    = 8
//        9+(10-9)/2   = 9+1/2    = 9
func searchInARotatedArray() {

}