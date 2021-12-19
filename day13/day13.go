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
	//Keep track of where the dots are
	dots := make(map[point]bool)
	//Read in the folds
	folds_dir := make([]string, 0)
	folds_num := make([]int, 0)

	max_x := 0
	max_y := 0
	//Read in the dots
	for fscanner.Scan() {
		line := strings.Split(fscanner.Text(), ",")
		x := line[0]
		y := line[1]
		xx, _ := strconv.Atoi(x)
		yy, _ := strconv.Atoi(y)
		dots[point{xx, yy}] = true

		if xx > max_x {
			max_x = xx
		}
		if yy > max_y {
			max_y = yy
		}
	}

	fmt.Println(max_x, max_y)
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
	// fmt.Println(folds_dir, folds_num)

	//Make the folds
	for i := 0; i < 1; i++ { // i < len(folds_dir); i++ {
		dir := folds_dir[i]
		num := folds_num[i]
		fmt.Println(dir, num)
		visible_dots := 0
		if dir == "x" {
			fold_x(&dots, num, &visible_dots)
		} else {
			fold_y(&dots, num, &visible_dots)
		}
		fmt.Println(visible_dots)
	}
}

func fold_x(dots *map[point]bool, num int, visible_dots *int) { //map[point]bool {
	// fmt.Println(dots)
	for p := range *dots {
		//If it's to the left of the fold, always increase the counter
		if p.x < num {
			*visible_dots += 1
		} else {
			//If it's to the right of the fold

			//Calculate the mirrored x value (y value will be the same)
			diff := p.x - num

			//Only increase the counter if there is NO mirrored point
			if !(*dots)[point{p.x - 2*diff, p.y}] {
				*visible_dots += 1
			}
		}
	}

	return
}

func fold_y(dots *map[point]bool, num int, visible_dots *int) { //map[point]bool {
	return
}
