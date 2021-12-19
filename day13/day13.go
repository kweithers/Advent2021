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
	//Read in the dots
	file, _ := os.Open("dots.txt")
	fscanner := bufio.NewScanner(file)
	//Keep track of the dots
	dots := make(map[point]bool)
	//Keep track of the folds
	folds_dir := make([]string, 0)
	folds_num := make([]int, 0)

	//Read in the dots
	for fscanner.Scan() {
		line := strings.Split(fscanner.Text(), ",")
		x := line[0]
		y := line[1]
		xx, _ := strconv.Atoi(x)
		yy, _ := strconv.Atoi(y)
		dots[point{xx, yy}] = true
	}

	//Read in the folds
	file, _ = os.Open("folds.txt")
	fscanner = bufio.NewScanner(file)
	for fscanner.Scan() {
		line := strings.Split(fscanner.Text(), " ")
		the_fold := strings.Split(line[2], "=")
		folds_dir = append(folds_dir, the_fold[0])
		num, _ := strconv.Atoi(the_fold[1])
		folds_num = append(folds_num, num)
	}

	//Successively make each fold
	for i := 0; i < len(folds_dir); i++ {
		dir := folds_dir[i]
		num := folds_num[i]
		visible_dots := 0

		if dir == "x" {
			dots = fold_x(&dots, num, &visible_dots)
		} else {
			dots = fold_y(&dots, num, &visible_dots)
		}

		if i == 0 {
			fmt.Println("Part 1:", visible_dots)
		}
	}

	fmt.Println("Part 2:")
	//Write the result to a file
	f, _ := os.Create("result.txt")
	defer f.Close()
	w := bufio.NewWriter(f)
	//Also print it to the console!
	for y := 0; y < 8; y++ {
		line := ""
		for x := 0; x < 45; x++ {
			if dots[point{x, y}] {
				line += "#"
			} else {
				line += "."
			}
		}
		fmt.Println(line)
		w.WriteString(line)
		w.WriteString("\n")
		w.Flush()
	}
}

func fold_x(dots *map[point]bool, num int, visible_dots *int) map[point]bool {
	new_dots := make(map[point]bool, 0)
	for p := range *dots {
		//If it's to the left of the fold, always increase the counter
		if p.x < num {
			*visible_dots += 1
			//Set it in the next map
			new_dots[p] = true
		} else {
			//If it's to the right of the fold

			//Calculate the mirrored x value (y value will be the same)
			diff := p.x - num

			//Only increase the counter if there is NO mirrored point
			if !(*dots)[point{p.x - 2*diff, p.y}] {
				*visible_dots += 1
			}
			//Set it in the next map
			new_dots[point{p.x - 2*diff, p.y}] = true
		}
	}
	return new_dots
}

func fold_y(dots *map[point]bool, num int, visible_dots *int) map[point]bool {
	new_dots := make(map[point]bool, 0)
	for p := range *dots {
		//If it's above the fold, always increase the counter
		if p.y < num {
			*visible_dots += 1
			//Set it in the next map
			new_dots[p] = true
		} else {
			//If it's below the fold

			//Calculate the mirrored y value (x value will be the same)
			diff := p.y - num

			//Only increase the counter if there is NO mirrored point
			if !(*dots)[point{p.x, p.y - 2*diff}] {
				*visible_dots += 1
			}
			//Set it in the next map
			new_dots[point{p.x, p.y - 2*diff}] = true
		}
	}
	return new_dots
}
