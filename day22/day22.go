package main

import (
	"bufio"
	"fmt"
	"math"
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
	file, _ := os.Open("test.txt")
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
	fmt.Println("Part 1", count_lights(m))

	//Part 2
	vols := make([]int, 0)
	//infinite off instruction: represents starting point
	start := instruction{false, math.MinInt, math.MaxInt, math.MinInt, math.MaxInt, math.MinInt, math.MaxInt}
	vols = append(vols, intersect(start, inst[0]))
	fmt.Println(vols)

	for i := 1; i < len(inst); i++ {
		for j := 0; j < i; j++ {
			vol := intersect(inst[i], inst[j])
			if vol != 0 {
				vols = append(vols, vol)
			}
		}
		// process_instruction(m, in)
	}
	fmt.Println(len(vols))

}

func intersect(i1, i2 instruction) int {

	xmin := max(i1.xmin, i2.xmin)
	xmax := min(i1.xmax, i2.xmax)

	ymin := max(i1.ymin, i2.ymin)
	ymax := min(i1.ymax, i2.ymax)

	zmin := max(i1.zmin, i2.zmin)
	zmax := min(i1.zmax, i2.zmax)

	if xmin > xmax || ymin > ymax || zmin > zmax {
		return 0
	}

	return (xmax - xmin) * (ymax - ymin) * (zmax - zmin)

}

func process_instruction(m map[point]bool, in instruction) {
	if (in.xmax < -50 || in.xmin > 50) && (in.ymax < -50 || in.ymin > 50) && (in.zmax < -50 || in.zmin > 50) {
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
