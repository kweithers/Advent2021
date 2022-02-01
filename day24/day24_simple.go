package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type alu struct {
	w, x, y, z int //Contains four data registers
}

type instruction struct {
	i string      //One of six legal instructions
	a string      //Always one of w,x,y,z. Also where the result of the operation in stored
	b interface{} //Either 1) one of w,x,y,z or 2) an int
}

type pair struct {
	first, second int
}

func main() {
	file, _ := os.Open("day24_full.txt")
	fscanner := bufio.NewScanner(file)
	var instructions []instruction

	var As, Bs, Cs []int
	//Read in the instructions
	i := 1
	for fscanner.Scan() {
		line := fscanner.Text()
		parts := strings.Split(line, " ")

		var in instruction
		if parts[0] == "inp" {
			in = instruction{i: parts[0], a: parts[1]}
		} else {
			v, err := strconv.Atoi(parts[2])
			if err == nil {
				in = instruction{i: parts[0], a: parts[1], b: v}
			} else {
				in = instruction{i: parts[0], a: parts[1], b: parts[2]}
			}
		}
		instructions = append(instructions, in)
		if i%18 == 5 {
			As = append(As, in.b.(int))
		} else if i%18 == 6 {
			Bs = append(Bs, in.b.(int))
		} else if i%18 == 16 {
			Cs = append(Cs, in.b.(int))
		}
		i++
	}

	pairs := []pair{{0, 13}, {1, 12}, {2, 3}, {4, 5}, {6, 7}, {8, 11}, {9, 10}}

	var results []pair
	for _, p := range pairs {
		results = append(results, find_match_max(As[p.first], Bs[p.first], Cs[p.first], As[p.second], Bs[p.second], Cs[p.second]))
	}
	fmt.Println(pairs)
	fmt.Println(results)
}

func find_match_max(a, b, c, a2, b2, c2 int) pair {
	var x, z int

	for w := 9; w > 0; w-- {
		for w2 := 9; w2 > 0; w2-- {

			x, z = 0, 0
			x_bool := ((z%26)+b != w)
			if x_bool {
				x = 1
			}
			z /= a
			z *= 25*x + 1
			z += (w + c) * x

			x = 0
			x_bool = ((z%26)+b2 != w2)
			if x_bool {
				x = 1
			}
			z /= a2
			z *= 25*x + 1
			z += (w2 + c2) * x

			if z == 0 {
				return pair{w, w2}
			}
		}
	}
	panic("STINKY")
}
