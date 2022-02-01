package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type pair struct {
	first, second int
}

func main() {
	file, _ := os.Open("day24.txt")
	fscanner := bufio.NewScanner(file)

	var As, Bs, Cs []int

	//Read in the instructions
	i := 1
	for fscanner.Scan() {
		line := fscanner.Text()
		parts := strings.Split(line, " ")

		if i%18 == 1 {
			i++
			continue
		}
		v, _ := strconv.Atoi(parts[2])

		//The only instructions that change are lines 5,6,16
		if i%18 == 5 {
			As = append(As, v)
		} else if i%18 == 6 {
			Bs = append(Bs, v)
		} else if i%18 == 16 {
			Cs = append(Cs, v)
		}
		i++
	}

	pairs := []pair{{0, 13}, {1, 12}, {2, 3}, {4, 5}, {6, 7}, {8, 11}, {9, 10}}

	var results []pair
	for _, p := range pairs {
		results = append(results, find_match(As[p.first], Bs[p.first], Cs[p.first], As[p.second], Bs[p.second], Cs[p.second], true))
	}
	fmt.Println("MAX")
	fmt.Println(pairs)
	fmt.Println(results)

	results = []pair{}
	for _, p := range pairs {
		results = append(results, find_match(As[p.first], Bs[p.first], Cs[p.first], As[p.second], Bs[p.second], Cs[p.second], false))
	}
	fmt.Println("MIN")
	fmt.Println(pairs)
	fmt.Println(results)

}

func find_match(a, b, c, a2, b2, c2 int, max bool) pair {
	var x, z int

	if max {
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
	} else {
		for w := 1; w < 10; w++ {
			for w2 := 1; w2 < 10; w2++ {

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
	}
	panic("NO MATCH")
}
