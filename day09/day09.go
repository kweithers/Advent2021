package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	//Setup a DFS
	var visited [100][100]bool
	basins := make([]int, 0)
	for j := 0; j < 100; j++ {
		for i := 0; i < 100; i++ {
			if !visited[j][i] && heights[j][i] < 9 {
				current_basin := 0
				visit(&visited, &heights, &current_basin, j, i)
				basins = append(basins, current_basin)
			}
		}
	}
	sort.Ints(basins)
	fmt.Println("Part 2:", basins[len(basins)-3]*basins[len(basins)-2]*basins[len(basins)-1])
}

func visit(visited *[100][100]bool, heights *[100][100]int, current_basin *int, j int, i int) {
	if visited[j][i] {
		return
	}
	visited[j][i] = true
	(*current_basin)++
	if j > 0 { //Go up
		if heights[j-1][i] < 9 && !visited[j-1][i] {
			visit(visited, heights, current_basin, j-1, i)
		}
	}
	if j < 99 { //Go down
		if heights[j+1][i] < 9 && !visited[j+1][i] {
			visit(visited, heights, current_basin, j+1, i)
		}
	}
	if i > 0 { //Go left
		if heights[j][i-1] < 9 && !visited[j][i-1] {
			visit(visited, heights, current_basin, j, i-1)
		}
	}
	if i < 99 { //Go right
		if heights[j][i+1] < 9 && !visited[j][i+1] {
			visit(visited, heights, current_basin, j, i+1)
		}
	}
}
