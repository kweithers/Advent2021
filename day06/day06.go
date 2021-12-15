package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("day06.txt")
	fscanner := bufio.NewScanner(file)
	fish := make([]int, 9)

	//Read in the fish
	for fscanner.Scan() {
		line := strings.Split(fscanner.Text(), ",")

		for _, val := range line {
			f, _ := strconv.Atoi(val)
			fish[f]++
		}
	}

	//Simulate behavior
	for i := 1; i <= 256; i++ {
		new_fish := make([]int, 9)

		for i, count := range fish {
			if i == 0 {
				new_fish[6] += count
				new_fish[8] += count
			} else {
				new_fish[i-1] += count
			}
		}
		fish = new_fish

		if i == 80 {
			fmt.Println("Part 1:", count_fish(fish))
		}
	}
	fmt.Println("Part 2:", count_fish(fish))
}

//Function to sum the slice of fish counter ints
func count_fish(fish []int) int {
	count := 0
	for _, c := range fish {
		count += c
	}
	return count
}
