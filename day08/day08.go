package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("day08.txt")
	fscanner := bufio.NewScanner(file)

	signals := make([][]string, 0)
	digits := make([][]string, 0)

	//Read in the signals and digits
	for fscanner.Scan() {
		line := strings.Split(fscanner.Text(), " ")
		this_signal := make([]string, 0)
		this_digit := make([]string, 0)

		for i, val := range line {
			if i < 10 {
				this_signal = append(this_signal, val)
			} else if i > 10 {
				this_digit = append(this_digit, val)
			}

		}
		signals = append(signals, this_signal)
		digits = append(digits, this_digit)
	}

	//1 uses 2 signals
	//7 uses 3 signals
	//4 uses 4 signals
	//8 uses 7 signals
	var counter int
	for _, line := range digits {
		for _, digit := range line {
			if len(digit) == 2 || len(digit) == 3 || len(digit) == 4 || len(digit) == 7 {
				counter += 1
			}
		}
	}
	fmt.Println("Part 1:", counter)

	//Part 2

}
