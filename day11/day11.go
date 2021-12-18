package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct {
	y int
	x int
}

func main() {
	file, _ := os.Open("day11.txt")
	fscanner := bufio.NewScanner(file)
	var octopus [10][10]int

	//Read in the octopus
	j := 0
	for fscanner.Scan() {
		line := strings.Split(fscanner.Text(), "")

		for i, h := range line {
			val, _ := strconv.Atoi(h)
			octopus[j][i] = val
		}
		j++
	}

	// Simulate behavior
	var flash_count int
	i := 1
	for {
		//Track if each octopus has flashed this step
		var flashed [10][10]bool
		//Count the number of flashes this step
		var step_flash_count int

		//Increase every energy level by 1
		for y := 0; y < 10; y++ {
			for x := 0; x < 10; x++ {
				octopus[y][x] += 1
			}
		}

		//Check if each octopus will flash (this will recursively check its neighbors)
		for y := 0; y < 10; y++ {
			for x := 0; x < 10; x++ {
				check_flash(&flashed, &octopus, y, x, &flash_count, &step_flash_count)
			}
		}

		//Set energy level to 0 for all octopus that flashed this step
		for y := 0; y < 10; y++ {
			for x := 0; x < 10; x++ {
				if flashed[y][x] {
					octopus[y][x] = 0
				}
			}
		}
		if i == 100 {
			fmt.Println("Part 1:", flash_count)
		}
		if step_flash_count == 100 {
			fmt.Println("Part 2:", i)
			break
		}
		i++
	}
}

func check_flash(flashed *[10][10]bool, octopus *[10][10]int, y int, x int, flash_count *int, step_flash_count *int) {
	if octopus[y][x] <= 9 || flashed[y][x] {
		return
	}
	//Now, we flash
	flashed[y][x] = true
	*flash_count += 1
	*step_flash_count += 1

	//Increase count for neighbors, then check flash for them
	neighbors := generate_neighbors(y, x)
	for _, n := range neighbors {
		octopus[n.y][n.x] += 1
		check_flash(flashed, octopus, n.y, n.x, flash_count, step_flash_count)
	}
}

func generate_neighbors(y int, x int) []point {
	dirs := []int{-1, 0, 1}
	points := make([]point, 0)
	for _, j := range dirs {
		for _, i := range dirs {
			if i == 0 && j == 0 {
				continue
			}
			if y+j >= 0 && y+j <= 9 && x+i >= 0 && x+i <= 9 { //If the point is in bounds, add it to the slice
				points = append(points, point{y + j, x + i})
			}
		}
	}
	return points
}
