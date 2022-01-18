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

type cube struct {
	on                                 bool
	xmin, xmax, ymin, ymax, zmin, zmax int
}

func main() {
	file, _ := os.Open("day22.txt")
	fscanner := bufio.NewScanner(file)
	inst := make([]cube, 0)

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
		inst = append(inst, cube{on, xmin, xmax, ymin, ymax, zmin, zmax})
	}

	//Apply the operations
	m := make(map[point]bool)

	for _, in := range inst {
		process_instruction(m, in)
	}
	fmt.Println("Part 1:", count_lights(m))

	//Part 2
	var volumes []cube

	for _, c := range inst {
		var overlaps []cube

		for _, finalCube := range volumes {
			intersection, intersected := finalCube.getIntersection(c)
			if intersected {
				overlaps = append(overlaps, intersection)
			}
		}

		if c.on {
			overlaps = append(overlaps, c)
		}

		volumes = append(volumes, overlaps...)
	}

	var cubes_on int
	for _, c := range volumes {
		cubes_on += c.getVolume()
	}
	fmt.Println("Part 2:", cubes_on)

}

func (c cube) getIntersection(c2 cube) (cube, bool) {
	xmin := max(c.xmin, c2.xmin)
	xmax := min(c.xmax, c2.xmax)
	ymin := max(c.ymin, c2.ymin)
	ymax := min(c.ymax, c2.ymax)
	zmin := max(c.zmin, c2.zmin)
	zmax := min(c.zmax, c2.zmax)

	if xmin > xmax || ymin > ymax || zmin > zmax {
		return cube{}, false
	}
	var on bool
	if c.on && c2.on {
		on = false
	} else if !c.on && !c2.on {
		on = true
	} else {
		on = c2.on
	}
	return cube{on, xmin, xmax, ymin, ymax, zmin, zmax}, true
}

func (c cube) getVolume() int {
	vol := (c.xmax - c.xmin + 1) * (c.ymax - c.ymin + 1) * (c.zmax - c.zmin + 1)
	if c.on {
		return vol
	}
	return -vol
}

func process_instruction(m map[point]bool, in cube) {
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
