package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//How to use the empty interface to handle different possible types//
/* func print_out_type(x interface{}) string {
    switch v := x.(type) {
        case string:
             return "A string"
        case int32:
             return "An Integer"
        default:
             return "A default"
    }
}
*/

type alu struct {
	w, x, y, z int //Contains four data registers
}

type instruction struct {
	i string      //One of six legal instructions
	a string      //Always one of w,x,y,z. Also where the result of the operation in stored
	b interface{} //Either 1) one of w,x,y,z or 2) an int
}

func (a *alu) inp(i instruction, v int) {
	switch i.a {
	case "w":
		a.w = v
	case "x":
		a.x = v
	case "y":
		a.y = v
	case "z":
		a.z = v
	default:
		panic("STINKY")
	}
}

func (a *alu) add(i instruction) {
	var temp int

	switch t := i.b.(type) {
	case string:
		switch i.b {
		case "w":
			temp = a.w
		case "x":
			temp = a.x
		case "y":
			temp = a.y
		case "z":
			temp = a.z
		default:
			panic("STINKY")
		}
	case int:
		temp = i.b.(int)
	default:
		panic(t)
	}

	switch i.a {
	case "w":
		a.w += temp
	case "x":
		a.x += temp
	case "y":
		a.y += temp
	case "z":
		a.z += temp
	}
}

func (a *alu) mul(i instruction) {
	var temp int

	switch t := i.b.(type) {
	case string:
		switch i.b {
		case "w":
			temp = a.w
		case "x":
			temp = a.x
		case "y":
			temp = a.y
		case "z":
			temp = a.z
		default:
			panic("STINKY")
		}
	case int:
		temp = i.b.(int)
	default:
		panic(t)
	}

	switch i.a {
	case "w":
		a.w *= temp
	case "x":
		a.x *= temp
	case "y":
		a.y *= temp
	case "z":
		a.z *= temp
	}
}

func (a *alu) div(i instruction) {
	var temp int

	switch t := i.b.(type) {
	case string:
		switch i.b {
		case "w":
			temp = a.w
		case "x":
			temp = a.x
		case "y":
			temp = a.y
		case "z":
			temp = a.z
		default:
			panic("STINKY")
		}
	case int:
		temp = i.b.(int)
	default:
		panic(t)
	}

	switch i.a {
	case "w":
		a.w /= temp
	case "x":
		a.x /= temp
	case "y":
		a.y /= temp
	case "z":
		a.z /= temp
	}
}

func (a *alu) mod(i instruction) {
	var temp int

	switch t := i.b.(type) {
	case string:
		switch i.b {
		case "w":
			temp = a.w
		case "x":
			temp = a.x
		case "y":
			temp = a.y
		case "z":
			temp = a.z
		default:
			panic("STINKY")
		}
	case int:
		temp = i.b.(int)
	default:
		panic(t)
	}
	switch i.a {
	case "w":
		a.w %= temp
	case "x":
		a.x %= temp
	case "y":
		a.y %= temp
	case "z":
		a.z %= temp
	}
}

func (a *alu) eq(i instruction) {
	var temp int

	switch t := i.b.(type) {
	case string:
		switch i.b {
		case "w":
			temp = a.w
		case "x":
			temp = a.x
		case "y":
			temp = a.y
		case "z":
			temp = a.z
		}
	case int:
		temp = i.b.(int)
	default:
		panic(t)
	}

	switch i.a {
	case "w":
		if a.w == temp {
			a.w = 1
		} else {
			a.w = 0
		}
	case "x":
		if a.x == temp {
			a.x = 1
		} else {
			a.x = 0
		}
	case "y":
		if a.y == temp {
			a.y = 1
		} else {
			a.y = 0
		}
	case "z":
		if a.z == temp {
			a.z = 1
		} else {
			a.z = 0
		}
	}
}

func main() {
	file, _ := os.Open("day24.txt")
	fscanner := bufio.NewScanner(file)
	var instructions []instruction

	//Read in the instructions
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
	}

	for n := 1; n < 10; n++ {
		a := alu{}
		str_n := strconv.Itoa(n)
		current_index := 0

		for _, in := range instructions {
			switch in.i {
			case "inp":
				digit, _ := strconv.Atoi(string(str_n[current_index : current_index+1]))
				a.inp(in, digit)
				current_index++
			case "add":
				a.add(in)
			case "mul":
				a.mul(in)
			case "div":
				a.div(in)
			case "mod":
				a.mod(in)
			case "eq":
				a.eq(in)
			}
		}

		fmt.Println(n, a.z)
	}

}
