package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("day09.txt")
	fscanner := bufio.NewScanner(file)

	var heights [100][100]int

	//Read in the heights
	j := 0
	for fscanner.Scan() {
		line := strings.Split(fscanner.Text(), "")

		for i, h := range line {
			val, _ := strconv.Atoi(h)
			heights[j][i] = val
		}
		j++
	}

	//Part 1
	risk := 0
	for j := 0; j < 100; j++ {
		for i := 0; i < 100; i++ {
			up, down, left, right := true, true, true, true
			if j > 0 {
				up = heights[j][i] < heights[j-1][i]
			}
			if j < 99 {
				down = heights[j][i] < heights[j+1][i]
			}
			if i > 0 {
				left = heights[j][i] < heights[j][i-1]
			}
			if i < 99 {
				right = heights[j][i] < heights[j][i+1]
			}

			if up && down && left && right {
				risk += (heights[j][i] + 1)
			}
		}
	}
	fmt.Println("Part 1:", risk)

	//Part 2
	//Setup a DFS, only explore if the neighboring node is higher (i.e. it would flow down)
	var visited [100][100]bool

	for j := 0; j < 100; j++ {
		for i := 0; i < 100; i++ {
}
