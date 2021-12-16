package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	file, _ := os.Open("day07.txt")
	fscanner := bufio.NewScanner(file)
	locations := make([]int, 0)

	var max int
	//Read in the positions
	for fscanner.Scan() {
		line := strings.Split(fscanner.Text(), ",")

		for _, val := range line {
			f, _ := strconv.Atoi(val)
			locations = append(locations, f)
			if max < f {
				max = f
			}
		}
	}

	//Part 1
	min_gas := math.MaxInt
	for x := 0; x <= max; x++ {
		this_gas := 0

		for _, c := range locations {
			this_gas += Abs(x - c)
		}

		if this_gas < min_gas {
			min_gas = this_gas
		} else {
			fmt.Println("Part 1:", min_gas)
			break
		}
	}

	//Part 2

	//Make the cost array
	//E.g. Moving 10 spaces - can lookup cost with costs[10]
	costs := make([]int, 0)
	costs = append(costs, 0)
	for i := 1; i <= max; i++ {
		costs = append(costs, costs[i-1]+i)
	}

	min_gas = math.MaxInt
	for x := 0; x <= max; x++ {
		this_gas := 0

		for _, c := range locations {
			this_gas += costs[Abs(x-c)]
		}

		if this_gas < min_gas {
			min_gas = this_gas
		} else {
			fmt.Println("Part 2:", min_gas)
			break
		}
	}

}
