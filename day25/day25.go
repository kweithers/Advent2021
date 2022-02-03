package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type point struct {
	y, x int
}

func main() {
	file, _ := os.Open("day25.txt")
	fscanner := bufio.NewScanner(file)

	sea_floor := make(map[point]int, 0)
	m := make(map[string]int, 0)
	m["."] = 0
	m[">"] = 1
	m["v"] = 2

	//Read in the sea floor
	depth := 0
	width := 0
	for fscanner.Scan() {
		line := fscanner.Text()
		parts := strings.Split(line, "")

		for i := 0; i < len(parts); i++ {
			if m[parts[i]] > 0 {
				sea_floor[point{depth, i}] = m[parts[i]]
			}
		}
		depth++
		width = len(parts)
	}

	//Take steps until no cucumber has changed
	steps := 0
	changed := true
	for changed {
		//Record if any cucumber has moved
		changed = false
		//Record the result after the first herd moves
		intermediate_sea_floor := make(map[point]int, 0)
		//Record the result after the second herd moves
		final_sea_floor := make(map[point]int, 0)

		//Move the east-facing herd
		for pt, cucumber := range sea_floor {
			//If the spot to the right is open, move there
			if cucumber == 1 && sea_floor[point{y: pt.y, x: (pt.x + 1) % width}] == 0 {
				intermediate_sea_floor[point{y: pt.y, x: (pt.x + 1) % width}] = 1
				changed = true
			} else {
				intermediate_sea_floor[pt] = cucumber
			}
		}

		// Move the south-facing herd
		for pt, cucumber := range intermediate_sea_floor {
			if cucumber == 2 && intermediate_sea_floor[point{y: (pt.y + 1) % depth, x: pt.x}] == 0 {
				//If the spot below is open, move there
				final_sea_floor[point{y: (pt.y + 1) % depth, x: pt.x}] = 2
				changed = true
			} else {
				final_sea_floor[pt] = cucumber
			}
		}
		steps++
		//Set up the sea floor for the next step
		sea_floor = final_sea_floor
	}
	fmt.Println(steps)
}
