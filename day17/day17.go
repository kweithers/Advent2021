package main

import (
	"fmt"
)

//Target area
const x_min int = 230
const x_max int = 283
const y_min int = -107
const y_max int = -57

func main() {
	current_max_height := 0
	path_counter := 0

	for i := 0; i <= 300; i++ {
		for j := -200; j < 200; j++ {
			simulate_path(i, j, &current_max_height, &path_counter)
		}
	}
	fmt.Println("Part 1:", current_max_height)
	fmt.Println("Part 2:", path_counter)
}

func simulate_path(x_vel int, y_vel int, current_max_height *int, path_counter *int) {
	x_pos := 0
	y_pos := 0
	max_height := 0

	//While we are to the left of and above the bottom right corner of the target area
	for x_pos < x_max && y_pos > y_min {
		x_pos += x_vel
		y_pos += y_vel

		if y_pos > max_height {
			max_height = y_pos
		}

		//If we are in the target area
		if x_pos >= x_min && x_pos <= x_max && y_pos >= y_min && y_pos <= y_max {
			(*path_counter)++
			if max_height > *current_max_height {
				*current_max_height = max_height
			}
			break
		}

		//Drag x velocity towards zero
		if x_vel > 0 {
			x_vel -= 1
		} else if x_vel < 0 {
			x_vel += 1
		}

		//Gravity decreases y velocity
		y_vel -= 1
	}
}
