package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("day01.txt")
	fscanner := bufio.NewScanner(file)
	nums := make([]int, 0)

	for fscanner.Scan() {
		i, _ := strconv.Atoi(fscanner.Text())
		nums = append(nums, i)
	}

	counter := 0
	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			counter++
		}
	}
	fmt.Println(counter)

	//Part 2
	counter2 := 0
	for i := 3; i < len(nums); i++ {
		if nums[i-3] < nums[i] {
			counter2++
		}
	}
	fmt.Println(counter2)
}
