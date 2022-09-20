package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		input := []int{1, 5, 10, 14, 16, 20}
		ch1 := make(chan int)
		ch2 := make(chan int)
		ch3 := make(chan int)
		ch4 := make(chan int)

		go pollDatasource1(input, ch1)

		go pollDatasource2(input, ch2)

		go func() {
			var sum int
			for i := 0; i < len(input); i++ {
				sum += input[i]
			}
			time.Sleep(time.Millisecond * 25)
			random := getRandValInBet(0, 1000)
			time.Sleep(time.Duration(rand.Int31n(random)) * time.Millisecond)
			fmt.Println("returning after ", random, " milliseconds")

			ch3 <- sum
		}()

		go func() {
			var sum int
			for i := 0; i < len(input); i++ {
				sum += input[i]
			}
			random := getRandValInBet(0, 1000)
			time.Sleep(time.Duration(rand.Int31n(random)) * time.Millisecond)
			fmt.Println("returning after ", random, " milliseconds")

			ch4 <- sum
		}()

		select {
		case resultFromDatasource1 := <-ch1:
			fmt.Println("Accepted response from ch1 ", resultFromDatasource1)
		case resultFromDatasource2 := <-ch2:
			fmt.Println("Accepted response from ch2 ", resultFromDatasource2)
		case resultFromDatasource3 := <-ch3:
			fmt.Println("Accepted response from ch3 ", resultFromDatasource3)
		case resultFromDatasource4 := <-ch4:
			fmt.Println("Accepted response from ch4 ", resultFromDatasource4)
		}
	}
}

func pollDatasource1(arr []int, ch1 chan int) {
	var sum int
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	random := getRandValInBet(0, 1000)
	time.Sleep(time.Duration(rand.Int31n(random)) * time.Millisecond)
	fmt.Println("returning after ", random, " milliseconds")

	ch1 <- sum
	close(ch1)
}

func pollDatasource2(arr []int, ch2 chan int) {
	var sum int
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	random := getRandValInBet(0, 1000)
	time.Sleep(time.Duration(rand.Int31n(random)) * time.Millisecond)
	fmt.Println("returning after ", random, " milliseconds")

	ch2 <- sum
	close(ch2)
}

func getRandValInBet(min int32, max int32) int32 {
	return rand.Int31n(max-min) + min
}
