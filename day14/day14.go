package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("rules.txt")
	fscanner := bufio.NewScanner(file)
	rules := make(map[string]string, 0)
	for fscanner.Scan() {
		line := strings.Split(fscanner.Text(), " -> ")
		rules[line[0]] = line[1]
	}

	polymer := "CFFPOHBCVVNPHCNBKVNV"
	polymer_map := make(map[string]int, 0)

	//Create map representation of the polymer
	for i := 1; i <= len(polymer)-1; i++ {
		polymer_map[polymer[i-1:i+1]] += 1
	}

	//Simulate steps
	for t := 1; t <= 40; t++ {
		//Create the map for the next step
		new_polymer_map := make(map[string]int, 0)

		//For every char pair
		for k, v := range polymer_map {
			a := string(k[0])
			b := string(k[1])
			//Find the new char to be inserted between ab
			new_char := rules[k]
			//Add the two new pairs a-new and new-b to the new map v times
			new_polymer_map[a+new_char] += v
			new_polymer_map[new_char+b] += v

		}
		//Set the map for the next step
		polymer_map = new_polymer_map
		if t == 10 {
			fmt.Println("Part 1:", find_difference(polymer, polymer_map))
		}
	}
	fmt.Println("Part 2:", find_difference(polymer, polymer_map))
}

func find_difference(polymer string, polymer_map map[string]int) int {
	//Count how many times each element occurs in the polymer
	elements := make(map[string]int, 0)

	//For each pair in the map
	for k, v := range polymer_map {
		//Increase both char counts
		elements[string(k[0])] += v
		elements[string(k[1])] += v
	}

	//Halve them since we are double counting
	for k := range elements {
		elements[k] /= 2
	}
	//Except the first and last element, which never change
	elements[string(polymer[0])]++
	elements[string(polymer[len(polymer)-1])]++

	//Find the difference between the max and min
	min := math.MaxInt
	max := 0
	for _, v := range elements {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return max - min
}
