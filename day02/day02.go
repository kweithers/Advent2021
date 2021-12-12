package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("day02.txt")
	fscanner := bufio.NewScanner(file)
	dirs := make([]string, 0)
	nums := make([]int, 0)

	for fscanner.Scan() {
		line := strings.Split(fscanner.Text(), " ")
		i, _ := strconv.Atoi(line[1])
		dirs = append(dirs, line[0])
		nums = append(nums, i)
	}

	// Part 1
	var x, y int

	for i := 0; i < len(dirs); i++ {
		switch dirs[i] {
		case "forward":
			x += nums[i]
		case "up":
			y -= nums[i]
		case "down":
			y += nums[i]
		}
	}
	fmt.Println("Part 1: ", x*y)

	// Part 2
	x, y = 0, 0
	var aim int

	for i := 0; i < len(dirs); i++ {
		switch dirs[i] {
		case "forward":
			x += nums[i]
			y += nums[i] * aim
		case "up":
			aim -= nums[i]
		case "down":
			aim += nums[i]
		}
	}
	fmt.Println("Part 2: ", x*y)
}
