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

		//Increase the depth on all existing items by 1
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
		fmt.Println("Before reduction", items)
		for !done {
			items, done = apply_action(items)
		}
		fmt.Println("After reduction", items)
	}
	fmt.Println("Part 1:", calculate_magnitude(items))
}

//This will read through the items and apply the first correct action, then returns the new list of items
func apply_action(items []item) ([]item, bool) {
	//1. If we find a pair to explode
	for index, it := range items {
		if it.depth == 5 {
			//These are the left and right items of that pair
			left := items[index]
			right := items[index+1]
			// fmt.Println(index, left, right)

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
			// fmt.Println("Explode")
			return new_items, false
		}
	}
	//2. If we find a number 10 or greater
	for index, it := range items {
		if it.value >= 10 {
			remainder := it.value % 2
			left := it.value / 2
			right := it.value/2 + remainder

			//Jank but works
			before := make([]item, len(items[:index]))
			copy(before, items[:index])
			after := make([]item, len(items[index+1:]))
			copy(after, items[index+1:])

			new_items := before
			new_items = append(new_items, item{value: left, depth: it.depth + 1})  //Add the new pair's left item
			new_items = append(new_items, item{value: right, depth: it.depth + 1}) //Add the new pair's right item
			new_items = append(new_items, after...)
			// fmt.Println("Split")

			return new_items, false
		}
	}
	return items, true
}

//Calculate the value of a number: 3*Left + 2*Right
func calculate_magnitude(items []item) int {
Outer:
	for {
		if len(items) == 1 {
			return items[0].value
		}
		for index := range items {
			if items[index].depth == items[index+1].depth {
				left := items[index]
				right := items[index+1]
				new_items := items[:index]
				new_items = append(new_items, item{value: 3*left.value + 2*right.value, depth: left.depth - 1}) //Add the regular value zero item
				new_items = append(new_items, items[index+2:]...)
				items = new_items
				continue Outer
			}
		}
	}
}
