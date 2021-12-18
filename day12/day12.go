package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	file, _ := os.Open("day12.txt")
	fscanner := bufio.NewScanner(file)
	//The graph, in adjacency list format
	vertices := make(map[string][]string, 0)

	//Read in the edges
	for fscanner.Scan() {
		line := strings.Split(fscanner.Text(), "-")
		a := line[0]
		b := line[1]

		if b == "start" || a == "end" {
			a, b = b, a
		}
		vertices[a] = append(vertices[a], b)

		//Edges are bidirectional, unless they contain start or end
		if a != "start" && b != "end" {
			vertices[b] = append(vertices[b], a)
		}
	}

	//Solve using backtracking with a stack
	stack := make([]string, 0)
	small_caves_visited := make(map[string]bool, 0)

	//Part 1
	path_counter := 0
	double_cave_used := true
	count_paths(&stack, "start", "end", &vertices, &path_counter, &small_caves_visited, &double_cave_used)
	fmt.Println("Part 1:", path_counter)
	//Part 2
	path_counter = 0
	double_cave_used = false
	count_paths(&stack, "start", "end", &vertices, &path_counter, &small_caves_visited, &double_cave_used)
	fmt.Println("Part 2:", path_counter)

}

func count_paths(stack *[]string, node string, dest string, graph *map[string][]string, path_counter *int, small_caves_visited *map[string]bool, double_cave_used *bool) {
	this_was_double_cave := false

	//If this is a small cave that we have visited
	if unicode.IsLower(rune(node[0])) && (*small_caves_visited)[node] {
		//If we have used our double cave already, then return since we can not use this cave
		if *double_cave_used {
			return
		}
		//If we havent used our double cave, use it here and keep track of it
		this_was_double_cave = true
		*double_cave_used = true
	}

	//Add cave to stack
	*stack = append(*stack, node)

	//If we have reached the end, increase the counter and print the path
	if node == "end" {
		*path_counter++
	}

	//If its a small cave, mark it as visited
	if unicode.IsLower(rune(node[0])) {
		(*small_caves_visited)[node] = true
	}

	//If we haven't reached the end, keep searching
	for _, d := range (*graph)[node] {
		count_paths(stack, d, "end", graph, path_counter, small_caves_visited, double_cave_used)
	}

	//Finally, pop this node from the stack
	n := len(*stack) - 1
	*stack = (*stack)[:n]

	//If its a small cave, and this was not the double cave, mark it as NOT visited
	if unicode.IsLower(rune(node[0])) && !this_was_double_cave {
		(*small_caves_visited)[node] = false
	}

	//If this was the double cave, unmark the double cave flag
	if this_was_double_cave {
		*double_cave_used = false
	}
}
