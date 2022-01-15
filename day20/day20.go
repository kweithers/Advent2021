package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const n = 100
const algo = "#.#.....##...##..#.#......#.#...#.#.#...###.##......###.##.##..##.#...#.....###.#.....#.#.#...#.#..###.###..###..#..##..###..##..##.##.#..###########.##....#.#......#...#.###..###...#.####..########.#####.#.#..##.##.##..###.##.####.#..##.##..#...#####..#.#.##.##...##..#..##.....###.#.#....####.##.#...##.########.#.##.#.....###....#..###.####....############.#.##...#.####...#...##.#.#..#..#......#..##...#.########.#.#...#####..#..######.#.#.....#####...##.###.#.#.##.........#.#.##..##.#..#..##..##.###.##.##."

type point struct {
	y int
	x int
}

type image struct {
	points         map[point]bool
	x_min          int
	x_max          int
	y_min          int
	y_max          int
	infinity_point bool
}

func (i image) get(y, x int) bool {
	if y < i.y_min || y > i.y_max || x < i.x_min || x > i.x_max {
		return i.infinity_point
	} else {
		return i.points[point{y, x}]
	}
}

func main() {
	file, _ := os.Open("day20.txt")
	fscanner := bufio.NewScanner(file)
	m := make(map[point]bool, 0)
	j := 0
	for fscanner.Scan() {
		line := fscanner.Text()
		for i := 0; i < n; i++ {
			if line[i] == 35 { //This is the rune value of #
				m[point{j, i}] = true
			}
		}
		j++
	}

	i := image{m, 0, n - 1, 0, n - 1, false}
	for t := 0; t < 50; t++ {
		i = enhance(i)
		if t == 1 {
			fmt.Println("Part 1:", count_lights(i))
		}
	}
	fmt.Println("Part 2:", count_lights(i))

}

func BinToDec(bin string) int {
	x, _ := strconv.ParseInt(bin, 2, 64)
	return int(x)
}

func enhance(im image) image {
	new_map := make(map[point]bool, 0)
	for j := im.y_min - 1; j <= im.y_max+1; j++ {
		for i := im.x_min - 1; i <= im.x_max+1; i++ {
			new_map[point{j, i}] = process_point(im, j, i)
		}
	}

	return image{
		points:         new_map,
		x_min:          im.x_min - 1,
		x_max:          im.x_max + 1,
		y_min:          im.y_min - 1,
		y_max:          im.y_max + 1,
		infinity_point: !im.infinity_point}
}

func process_point(im image, y, x int) bool {
	s := ""
	for j := -1; j <= 1; j++ {
		for i := -1; i <= 1; i++ {
			if im.get(y+j, x+i) {
				s += "1"
			} else {
				s += "0"
			}
		}
	}
	if algo[BinToDec(s)] == 35 {
		return true
	} else {
		return false
	}
}

func count_lights(im image) int {
	c := 0
	for _, v := range im.points {
		if v {
			c++
		}
	}
	return c
}
