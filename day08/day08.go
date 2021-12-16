package main

import (
	"bufio"
	"fmt"
	"math"
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

	//1 uses 2 signals, 7 uses 3 signals, 4 uses 4 signals, 8 uses 7 signals
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
	var sum int
	for i, line := range signals {
		var one string
		var four string
		for _, signal := range line {
			if len(signal) == 2 {
				one = signal
			} else if len(signal) == 4 {
				four = signal
			}
		}
		var result int

		for n, digit := range digits[i] {
			if len(digit) == 2 {
				result += 1 * int(math.Pow10(4-n-1))
			} else if len(digit) == 3 {
				result += 7 * int(math.Pow10(4-n-1))
			} else if len(digit) == 4 {
				result += 4 * int(math.Pow10(4-n-1))
			} else if len(digit) == 7 {
				result += 8 * int(math.Pow10(4-n-1))
			} else if len(digit) == 5 {
				//If the number of signals present is not enough to determine the digit,
				//compare with the other digits we already do know (1,4)
				overlap_with_1 := 0
				overlap_with_4 := 0
				for _, d := range digit {
					if strings.Contains(one, string(d)) {
						overlap_with_1++
					}
					if strings.Contains(four, string(d)) {
						overlap_with_4++
					}
				}

				if overlap_with_1 == 2 {
					result += 3 * int(math.Pow10(4-n-1))
				} else if overlap_with_4 == 2 {
					result += 2 * int(math.Pow10(4-n-1))
				} else {
					result += 5 * int(math.Pow10(4-n-1))

				}
			} else if len(digit) == 6 {
				//If the number of signals present is not enough to determine the digit,
				//compare with some other digits we already do know (1,4)
				overlap_with_1 := 0
				overlap_with_4 := 0
				for _, d := range digit {
					if strings.Contains(one, string(d)) {
						overlap_with_1++
					}
					if strings.Contains(four, string(d)) {
						overlap_with_4++
					}
				}
				if overlap_with_1 == 1 {
					result += 6 * int(math.Pow10(4-n-1))
				} else if overlap_with_4 == 4 {
					result += 9 * int(math.Pow10(4-n-1))
				} else {
					result += 0 * int(math.Pow10(4-n-1))
				}
			}
		}
		sum += result
	}
	fmt.Println("Part 2:", sum)
}
