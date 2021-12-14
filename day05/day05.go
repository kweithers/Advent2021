package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

func main() {
	file, _ := os.Open("day05.txt")
	fscanner := bufio.NewScanner(file)
	counts := make(map[point]int, 1000000) //Count how many vents appear at each point

	//Read in the vents
	for fscanner.Scan() {
		line := strings.Split(fscanner.Text(), " -> ")

		point1 := strings.Split(line[0], (","))
		point2 := strings.Split(line[1], (","))
		x1, _ := strconv.Atoi(point1[0])
		y1, _ := strconv.Atoi(point1[1])
		x2, _ := strconv.Atoi(point2[0])
		y2, _ := strconv.Atoi(point2[1])

		if x1 == x2 { //ONLY horizontal and vertical lines for now
			if y1 > y2 {
				y1, y2 = y2, y1
			}
			for i := y1; i <= y2; i++ {
				counts[point{x1, i}] += 1
			}

		} else if y1 == y2 {
			if x1 > x2 {
				x1, x2 = x2, x1
			}
			for i := x1; i <= x2; i++ {
				counts[point{i, y1}] += 1
			}
		} else { //Diagonals
			if x1 > x2 {
				x1, x2 = x2, x1
				y1, y2 = y2, y1
			}
			//Now the first point will always be on the left.

			if y1 < y2 { //Positive slope line
				for i := 0; i <= (x2 - x1); i++ {
					counts[point{x1 + i, y1 + i}] += 1
				}
			} else { //Negative slope line
				for i := 0; i <= (x2 - x1); i++ {
					counts[point{x1 + i, y1 - i}] += 1
				}
			}
		}

	}

	//Count how many are greater than two
	counter := 0
	for _, v := range counts {
		if v >= 2 {
			counter++
		}
	}
	fmt.Println(counter)
}
