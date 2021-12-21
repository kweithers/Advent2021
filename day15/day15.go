package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	y int
	x int
}

func main() {
	file, _ := os.Open("day15.txt")
	fscanner := bufio.NewScanner(file)
	weights := make(map[Point]int)

	//Read in the weights
	j := 0
	for fscanner.Scan() {
		line := strings.Split(fscanner.Text(), "")

		for i, h := range line {
			val, _ := strconv.Atoi(h)
			weights[Point{j, i}] = val
		}
		j++

	}

	//Part 2 - Enlarge the weights to 500x500
	for y := 0; y < 100; y++ {
		for x := 0; x < 100; x++ {
			for j := 0; j < 5; j++ {
				for i := 0; i < 5; i++ {
					if i == 0 && j == 0 {
						continue
					}
					new := (weights[Point{y, x}] + i + j)
					if new > 9 {
						new -= 9
					}
					weights[Point{y + 100*j, x + 100*i}] = new
				}
			}
		}
	}

	//Create Adjacency List
	graph := make(map[Point][]Point, 1)
	//Also, keep track of distances for Dijkstra
	dists := make(map[Point]int, 1)
	for y := 0; y < 500; y++ {
		for x := 0; x < 500; x++ {
			dist := math.MaxInt32
			if x == 0 && y == 0 {
				dist = 0
			}
			graph[Point{y, x}] = generate_neighbors(y, x)
			dists[Point{y, x}] = dist
		}
	}

	//Create the priority queue for Dijkstra
	pq := make(PriorityQueue, 500*500)
	//Keep a map for quick item access
	items := make(map[Point]*Item, 0)

	i := 0
	for point := range graph {
		it := Item{
			point:    point,
			priority: dists[point],
			index:    i,
		}
		pq[i] = &it
		items[point] = &it
		i++
	}
	heap.Init(&pq)

	//Run Dijkstra
	for len(pq) > 0 {
		//Extract Min
		u := *(heap.Pop(&pq).(*Item))

		//For each neighbor v of u
		for _, neighbor := range graph[u.point] {
			//Calculate new distance for this neighbor
			alt := dists[u.point] + weights[neighbor]
			//If it is smaller than its current distance, replace it and update the priority queue
			if alt < dists[neighbor] {
				dists[neighbor] = alt
				pq.update(items[neighbor], neighbor, alt)
			}
		}
	}

	fmt.Println("Part 1:", dists[Point{99, 99}])
	fmt.Println("Part 2:", dists[Point{499, 499}])
}

func generate_neighbors(y int, x int) []Point {
	dirs := []int{-1, 0, 1}
	points := make([]Point, 0)
	for _, j := range dirs {
		for _, i := range dirs {
			if i == 0 && j == 0 || i != 0 && j != 00 { //Exclude diagonals and self
				continue
			}
			if y+j >= 0 && y+j < 500 && x+i >= 0 && x+i < 500 { //If the point is in bounds, add it to the slice
				points = append(points, Point{y + j, x + i})
			}
		}
	}
	return points
}
