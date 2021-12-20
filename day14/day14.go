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
	var polymer strings.Builder
	// polymer.WriteString("NNCB")
	polymer.WriteString("CFFPOHBCVVNPHCNBKVNV")

	for fscanner.Scan() {
		line := strings.Split(fscanner.Text(), " -> ")
		rules[line[0]] = line[1]
	}

	//Simulate ten time steps
	for t := 1; t <= 10; t++ {
		fmt.Println(t)
		var new_polymer strings.Builder
		s := polymer.String()

		//Write the first char
		new_polymer.WriteString(string(s[0]))

		//For every char pair
		for i := 1; i <= len(s)-1; i++ {
			a := string(s[i-1])
			b := string(s[i])
			//Insert the new char
			new_polymer.WriteString(rules[a+b])
			//Write the next char
			new_polymer.WriteString(b)
		}
		polymer = new_polymer
	}
	final_polymer := polymer.String()
	elements := make(map[string]int, 0)
	for _, e := range final_polymer {
		elements[string(e)] += 1
	}

	min := math.MaxInt
	max := 0
	for _, v := range elements {
		if v > max {
			max = v
		} else if v < min {
			min = v
		}
	}
	fmt.Println(max - min)
}
