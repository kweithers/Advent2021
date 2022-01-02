package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type item struct {
	value int
	depth int //starts at 1
}

func main() {
	file, _ := os.Open("day18.txt")
	fscanner := bufio.NewScanner(file)
	lines := make([]string, 0)

	for fscanner.Scan() {
		lines = append(lines, fscanner.Text())
	}

	items := make([]item, 0)
	for i, l := range lines {
		current_depth := 1

		//Increase the depth on all existing items by 1, since we are adding it to the next number
		if i > 1 { //Except for the first two numbers we read
			for i := range items {
				items[i].depth++
			}
		}
		//Append the new number
		for _, v := range l {
			if num, err := strconv.Atoi(string(v)); err == nil {
				items = append(items, item{value: num, depth: current_depth})
			} else if string(v) == "[" {
				current_depth++
			} else if string(v) == "]" {
				current_depth--
			}
		}
		if i == 0 { //If it's the first number, just continue to read the next one
			continue
		}

		//Reduce the number
		done := false
		for !done {
			items, done = apply_action(items)
		}
	}
	fmt.Println("Part 1:", calculate_magnitude(items))

	//Part 2
	max_mag := 0
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines); j++ {
			if i == j {
				continue
			}
			items := make([]item, 0)
			current_depth := 1
			//Append the first number
			for _, v := range lines[i] {
				if num, err := strconv.Atoi(string(v)); err == nil {
					items = append(items, item{value: num, depth: current_depth})
				} else if string(v) == "[" {
					current_depth++
				} else if string(v) == "]" {
					current_depth--
				}
			}
			//Append the second number
			for _, v := range lines[j] {
				if num, err := strconv.Atoi(string(v)); err == nil {
					items = append(items, item{value: num, depth: current_depth})
				} else if string(v) == "[" {
					current_depth++
				} else if string(v) == "]" {
					current_depth--
				}
			}
			//Reduce it
			done := false
			for !done {
				items, done = apply_action(items)
			}
			//Calculate Magnitude
			mag := calculate_magnitude(items)
			//If it's the biggest so far, record it
			if mag > max_mag {
				max_mag = mag
			}
		}
	}
	fmt.Println("Part 2:", max_mag)

}

//This will read through the items and apply the first correct action
//It returns the new list of items, and a bool representing whether we are done reducing
func apply_action(items []item) ([]item, bool) {
	//1. If we find a pair to explode
	for index, it := range items {
		if it.depth == 5 {
			//These are the left and right items of that pair
			left := items[index]
			right := items[index+1]

			//Add the left one to its left neighbor, if it exists
			if index > 0 {
				items[index-1].value += left.value
			}
			//Add the right one to its right neighbor, if it exists
			if index < len(items)-2 {
				items[index+2].value += right.value
			}

			//Remove left and right from the list of items
			new_items := items[:index]
			new_items = append(new_items, item{value: 0, depth: left.depth - 1}) //Add the regular value zero item
			new_items = append(new_items, items[index+2:]...)
			return new_items, false
		}
	}
	//2. If we find a number 10 or greater
	for index, it := range items {
		if it.value >= 10 {
			remainder := it.value % 2
			left := it.value / 2
			right := it.value/2 + remainder

			before := make([]item, len(items[:index]))
			copy(before, items[:index])
			after := make([]item, len(items[index+1:]))
			copy(after, items[index+1:])

			new_items := before
			new_items = append(new_items, item{value: left, depth: it.depth + 1})  //Add the new pair's left item
			new_items = append(new_items, item{value: right, depth: it.depth + 1}) //Add the new pair's right item
			new_items = append(new_items, after...)

			return new_items, false
		}
	}
	return items, true
}

//Calculate the magnitude of a final sum using the recursive formula: 3*Left + 2*Right
//It continues until the entire sum is reduced and we know its magnitude
func calculate_magnitude(items []item) int {
Outer:
	for {
		if len(items) == 1 { //If there's only one item left, it contains the magnitude for the sum
			return items[0].value
		}
		for index := range items {
			if items[index].depth == items[index+1].depth { //If we find a pair at the same depth
				left := items[index]
				right := items[index+1]
				//Remove the two items
				new_items := items[:index]
				//Add the magnitude of this pair
				new_items = append(new_items, item{value: 3*left.value + 2*right.value, depth: left.depth - 1})
				new_items = append(new_items, items[index+2:]...)
				items = new_items
				continue Outer
			}
		}
	}
}
