package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type square struct {
	x      int //Left Column is 0
	y      int //Top Row is 0
	called bool
}

type bingo_tracker struct {
	counts   [10]int //0-4 are columns, 5-9 are rows
	max      int
	finished bool
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	file, _ := os.Open("number_order.txt")
	fscanner := bufio.NewScanner(file)
	nums := make([]int, 0)

	//Read in the number order
	for fscanner.Scan() {
		line := strings.Split(fscanner.Text(), ",")

		for _, num := range line {
			nn, _ := strconv.Atoi(num)
			nums = append(nums, nn)
		}
	}

	//Read in Bingo Boards
	file, _ = os.Open("bingo_boards.txt")
	fscanner = bufio.NewScanner(file)
	lines := make([][]int, 0)

	for fscanner.Scan() {
		line := strings.Split(fscanner.Text(), " ")
		row := make([]int, 0)

		for _, num := range line {
			if num == "" {
				continue
			}
			nn, _ := strconv.Atoi(num)
			row = append(row, nn)
		}
		if len(row) == 5 {
			lines = append(lines, row)
		}
	}

	// Make the boards
	var boards [100]map[int]*square
	for i := 0; i < 100; i++ {
		boards[i] = make(map[int]*square, 25)
	}

	for i, line := range lines {
		for col, n := range line {
			boards[i/5][n] = &square{x: col, y: i % 5, called: false}
		}
	}
	//Call the numbers
	var trackers [100]bingo_tracker

	done := 0
	for _, n := range nums {
		//Mark all the boards
		for i, b := range boards {
			if _, ok := b[n]; ok {
				b[n].called = true
				trackers[i].counts[b[n].x] += 1   //Column
				trackers[i].counts[b[n].y+5] += 1 //Row
				trackers[i].max = max(trackers[i].max, max(trackers[i].counts[b[n].x], trackers[i].counts[b[n].y+5]))

				if trackers[i].max == 5 && trackers[i].finished == false {
					trackers[i].finished = true
					done += 1
					if done == 1 || done == 100 {
						sum := 0
						for k, v := range b {
							if v.called == false {
								sum += k
							}
						}
						if done == 1 {
							fmt.Println("Part 1:", n*sum)
						} else {
							fmt.Println("Part 2:", n*sum)
							return
						}
					}
				}
			}
		}
	}
}
