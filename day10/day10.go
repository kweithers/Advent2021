package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	file, _ := os.Open("day10.txt")
	fscanner := bufio.NewScanner(file)
	chunks := make([]string, 0)

	//Read in the chunks
	for fscanner.Scan() {
		chunks = append(chunks, fscanner.Text())
	}

	//Value of the illegal characters for part 1
	p := make(map[string]int, 0)
	p[")"] = 3
	p["]"] = 57
	p["}"] = 1197
	p[">"] = 25137

	//Map containing pairs of characters
	pair := make(map[string]string, 0)
	pair["("] = ")"
	pair["["] = "]"
	pair["{"] = "}"
	pair["<"] = ">"

	//Value of adding characters for part 2
	p2 := make(map[string]int, 0)
	p2[")"] = 1
	p2["]"] = 2
	p2["}"] = 3
	p2[">"] = 4

	//Iterate through each line
	part1 := 0
	part2 := make([]int, 0)

OuterLoop:
	for _, c := range chunks {
		stack := make([]string, 0) // Use a stack to keep track of chars

		for _, x := range c {
			xx := string(x)
			if xx == "(" || xx == "[" || xx == "{" || xx == "<" {
				stack = append(stack, xx) //Push to stack
			} else {
				n := len(stack) - 1
				top_of_stack := stack[n]

				if pair[top_of_stack] != xx { //If the top of the stack does not pair with this new element, we know the illegal character
					part1 += p[xx]
					continue OuterLoop
				} else { //If they do match, pop the first char of the pair from the stack and continue parsing
					stack = stack[:n] //Pop from stack
				}
			}
		}
		//If a line gets to this point, we know it is an incomplete line.
		//We pop from the stack one by one, adding the correct closing chars and updating the score
		local_score := 0

		for len(stack) > 0 {
			n := len(stack) - 1
			top_of_stack := stack[n]

			local_score *= 5
			local_score += p2[pair[top_of_stack]]
			stack = stack[:n]
		}
		part2 = append(part2, local_score)
	}
	fmt.Println("Part 1:", part1)
	sort.Ints(part2)
	fmt.Println("Part 2:", part2[len(part2)/2])
}
