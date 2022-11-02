package main

import "fmt"

func isPalindrome(s string) bool {

	i := 0
	j := len(s) - 1

	for i <= j {
		if s[i] != s[j] {
			return false
		}

		i++
		j--
	}

	return true
}

func main() {

	//"kayak"
	//"hello"
	//"RACEACAR"
	//"A"
	//"ABCDABCD"
	//"DCBAABCD"
	//"ABCBA"

	palindrome1 := isPalindrome("kayak")
	fmt.Println(palindrome1)
	palindrome2 := isPalindrome("hello")
	fmt.Println(palindrome2)

}
