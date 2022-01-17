package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x, y, z int
}

type instruction struct {
	on                                 bool
	xmin, xmax, ymin, ymax, zmin, zmax int
}

func main() {
	file, _ := os.Open("day22.txt")
	fscanner := bufio.NewScanner(file)
	inst := make([]instruction, 0)

	//Read in the data
	for fscanner.Scan() {
		line := fscanner.Text()
		parts := strings.Split(line, " ")

		var on bool
		if parts[0] == "on" {
			on = true
		}
		nums := strings.Split(parts[1], ",")

		xs := strings.Split(strings.Split(nums[0], "=")[1], "..")
		ys := strings.Split(strings.Split(nums[1], "=")[1], "..")
		zs := strings.Split(strings.Split(nums[2], "=")[1], "..")

		xmin, _ := strconv.Atoi(xs[0])
		xmax, _ := strconv.Atoi(xs[1])
		ymin, _ := strconv.Atoi(ys[0])
		ymax, _ := strconv.Atoi(ys[1])
		zmin, _ := strconv.Atoi(zs[0])
		zmax, _ := strconv.Atoi(zs[1])
		inst = append(inst, instruction{on, xmin, xmax, ymin, ymax, zmin, zmax})
	}

	//Apply the operations
	m := make(map[point]bool)

	for _, in := range inst {
		process_instruction(m, in)
	}
	fmt.Println(count_lights(m))
}

func process_instruction(m map[point]bool, in instruction) {
	fmt.Println(in)
	if (in.xmax < -50 || in.xmin > 50) && (in.ymax < -50 || in.ymin > 50) && (in.zmax < -50 || in.zmin > 50) {
		fmt.Println("SKIPPING")
		return
	}

	for i := max(-50, in.xmin); i <= min(50, in.xmax); i++ {
		for j := max(-50, in.ymin); j <= min(50, in.ymax); j++ {
			for k := max(-50, in.zmin); k <= min(50, in.zmax); k++ {
				m[point{i, j, k}] = in.on
			}
		}
	}
}

func count_lights(m map[point]bool) int {
	c := 0
	for _, v := range m {
		if v {
			c++
		}
	}
	return c
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
