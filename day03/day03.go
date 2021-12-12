package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("day03.txt")
	fscanner := bufio.NewScanner(file)
	nums := make([]string, 0)
	for fscanner.Scan() {
		nums = append(nums, fscanner.Text())
	}

	bits := make([]int, len(nums[0]))

	for _, num := range nums {
		s := string(num)
		for i, bit := range s {
			n, _ := strconv.Atoi(string(bit))
			bits[i] += n
		}
	}

	// Part 1
	var gamma, epsilon int

	for i, n := range bits {
		if n > 500 {
			gamma += int(math.Pow(2, float64(len(bits)-i-1)))
		} else {
			epsilon += int(math.Pow(2, float64(len(bits)-i-1)))
		}
	}
	fmt.Println("Part 1: ", gamma*epsilon)

	//Part 2

	//Oxygen Rating
	excluded := make([]int, len(nums))
	sum_excluded := 0
	position := 0
	for sum_excluded < 999 {
		bit := find_most_common_bit(&nums, position, &excluded)
		exclude_nums_and_increase_excluded(bit, &nums, position, &excluded, &sum_excluded)
		position += 1
	}
	oxygen := convert_bit_string_to_decimal_int(find_element(&nums, &excluded))

	//CO2 Rating
	excluded = make([]int, len(nums))
	sum_excluded = 0
	position = 0
	for sum_excluded < 999 {
		bit := find_least_common_bit(&nums, position, &excluded)
		exclude_nums_and_increase_excluded(bit, &nums, position, &excluded, &sum_excluded)
		position += 1
	}
	co2 := convert_bit_string_to_decimal_int(find_element(&nums, &excluded))

	fmt.Println("Part 2: ", oxygen*co2)
}

func find_most_common_bit(nums *[]string, position int, excluded *[]int) int {
	var line_counter float64
	var bit_counter float64
	for i, n := range *nums {
		if (*excluded)[i] == 0 {
			line_counter += 1
			nn, _ := strconv.Atoi(string(n[position]))
			bit_counter += float64(nn)
		}
	}
	if bit_counter*2 == line_counter {
		return 1
	}
	return int(math.Round(float64(bit_counter / line_counter)))
}

func find_least_common_bit(nums *[]string, position int, excluded *[]int) int {
	var line_counter float64
	var bit_counter float64
	for i, n := range *nums {
		if (*excluded)[i] == 0 {
			line_counter += 1
			nn, _ := strconv.Atoi(string(n[position]))
			bit_counter += float64(nn)
		}
	}
	if bit_counter*2 == line_counter {
		return 0
	}
	if int(math.Round(float64(bit_counter/line_counter))) == 0 {
		return 1
	} else {
		return 0
	}
}

func exclude_nums_and_increase_excluded(most_common_bit int, nums *[]string, position int, excluded *[]int, sum_excluded *int) {
	for i, n := range *nums {
		if (*excluded)[i] == 0 {
			this_bit, _ := strconv.Atoi(string(n[position]))
			if this_bit != most_common_bit {
				(*excluded)[i] = 1
				(*sum_excluded) += 1
			}
		}
	}
}

func find_element(nums *[]string, excluded *[]int) string {
	for i, n := range *nums {
		if (*excluded)[i] == 0 {
			return n
		}
	}
	return "ERROR"
}

func convert_bit_string_to_decimal_int(input string) int {
	var result int
	for i, n := range input {
		this_bit, _ := strconv.Atoi(string(n))
		if this_bit == 1 {
			result += int(math.Pow(2, float64(len(input)-i-1)))
		}
	}
	return result
}
