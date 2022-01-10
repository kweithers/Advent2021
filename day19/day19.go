package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//A beacon is simply three integers representing its x,y,z position from a scanner
type beacon [3]int

//A matrix is a 2d array of ints
type matrix [][]int

//A pair of beacons one in the proper orientation, one in a different orientation
//These will be passed to a function that will find the proper orientation for a scanner
type beaconPair struct {
	original beacon
	rotated  beacon
}

//A scanner reads the location of several beacons
type scanner struct {
	beacons []beacon
}

type rotation struct {
	x     int //These are either 0,1,2 representing x,y,z
	x_dir int //These are either 1 or -1 representing positive or negative
	y     int
	y_dir int
	z     int
	z_dir int
}

func abs(x int) int {
	if x > 0 {
		return x
	} else {
		return -x
	}
}

func distance(a beacon, b beacon) int {
	x := abs(a[0] - b[0])
	y := abs(a[1] - b[1])
	z := abs(a[2] - b[2])
	return x + y + z
}

func main() {
	file, _ := os.Open("day19.txt")
	fscanner := bufio.NewScanner(file)
	scanners := make([]scanner, 0)

	var this_scanner scanner
	for fscanner.Scan() {
		line := fscanner.Text()
		if line == "" {
			scanners = append(scanners, this_scanner)
		} else if strings.Contains(line, "scanner") {
			this_scanner = scanner{}
		} else {
			nums := strings.Split(line, ",")
			x, _ := strconv.Atoi(nums[0])
			y, _ := strconv.Atoi(nums[1])
			z, _ := strconv.Atoi(nums[2])
			this_scanner.beacons = append(this_scanner.beacons, beacon{x, y, z})
		}
	}
	//Append the final beacon
	scanners = append(scanners, this_scanner)

	matrices := make([]matrix, 0)
	//Calculate distances matrices for each scanner
	for _, s := range scanners {
		n := len(s.beacons)
		matrix := make([][]int, n)
		for row := 0; row < n; row++ {
			matrix[row] = make([]int, n)
		}
		for y, b1 := range s.beacons {
			for x, b2 := range s.beacons {
				matrix[y][x] = distance(b1, b2)
			}
		}
		matrices = append(matrices, matrix)
	}

	//All 24 valid rotations
	rotations := []rotation{
		{0, 1, 1, 1, 2, 1}, {0, 1, 2, 1, 1, -1}, {0, 1, 1, -1, 2, -1}, {0, 1, 2, -1, 1, 1},
		{0, -1, 1, 1, 2, -1}, {0, -1, 2, -1, 1, -1}, {0, -1, 1, -1, 2, 1}, {0, -1, 2, 1, 1, 1},
		{1, 1, 0, 1, 2, -1}, {1, 1, 2, -1, 0, -1}, {1, 1, 0, -1, 2, 1}, {1, 1, 2, 1, 0, 1},
		{1, -1, 0, 1, 2, 1}, {1, -1, 2, 1, 0, -1}, {1, -1, 0, -1, 2, -1}, {1, -1, 2, -1, 0, 1},
		{2, 1, 0, 1, 1, 1}, {2, 1, 1, 1, 0, -1}, {2, 1, 0, -1, 1, -1}, {2, 1, 1, -1, 0, 1},
		{2, -1, 0, 1, 1, -1}, {2, -1, 1, -1, 0, -1}, {2, -1, 0, -1, 1, 1}, {2, -1, 1, 1, 0, 1}}

	properly_oriented := make(map[int]bool, 0)
	properly_oriented[0] = true
	for len(properly_oriented) < len(scanners) {
		for x := range matrices {
			for y := x + 1; y < len(matrices); y++ {
				if !properly_oriented[x] && !properly_oriented[y] {
					continue
				}
				beacon_pairs := make([]beaconPair, 0)
				for i := range matrices[x] {
					for j := range matrices[y] {
						l := len(Intersection(matrices[x][i], matrices[y][j]))
						if l >= 11 {
							beacon_pairs = append(beacon_pairs, beaconPair{scanners[x].beacons[i], scanners[y].beacons[j]})
						}
					}
				}
				if len(beacon_pairs) >= 12 {
					// fmt.Println(x, y)
					if properly_oriented[x] {
						r, x_diff, y_diff, z_diff := find_rotation(beacon_pairs, rotations)
						scanners[y] = reorient(&scanners, r, y, x_diff, y_diff, z_diff)
						properly_oriented[y] = true
					} else if properly_oriented[y] {
						beacon_pairs := flip_beacon_pairs(beacon_pairs)
						r, x_diff, y_diff, z_diff := find_rotation(beacon_pairs, rotations)
						scanners[x] = reorient(&scanners, r, x, x_diff, y_diff, z_diff)
						properly_oriented[x] = true
					} else {
						panic("This should never happen.")
					}
				}
			}
		}
	}

	//After we have reoriented all of our beacons, count how many unique beacons there are
	beacon_map := make(map[beacon]bool)
	for _, s := range scanners {
		for _, b := range s.beacons {
			beacon_map[b] = true
		}
	}
	fmt.Println(len(beacon_map))
}

func Intersection(a, b []int) (c []int) {
	m := make(map[int]bool)

	for _, item := range a {
		if item == 0 {
			continue //skip distance to self
		}
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; ok {
			c = append(c, item)
		}
	}
	return
}

//Take two beacons and find the proper orientatoin of the second one relative to the first one
func find_rotation(beacon_pairs []beaconPair, rotations []rotation) (rotation, int, int, int) {
	for _, r := range rotations {
		x_vals := make([]int, 0)
		y_vals := make([]int, 0)
		z_vals := make([]int, 0)
		for _, bp := range beacon_pairs {
			x_vals = append(x_vals, (r.x_dir*bp.rotated[r.x])-bp.original[0])
			y_vals = append(y_vals, (r.y_dir*bp.rotated[r.y])-bp.original[1])
			z_vals = append(z_vals, (r.z_dir*bp.rotated[r.z])-bp.original[2])
		}
		//If all three of these are all the same, it is the location of the rotated scanner relative to the original scanner
		if allSameValue(x_vals) && allSameValue(y_vals) && allSameValue(z_vals) {
			// fmt.Println("SOLVED", r, x_vals[0], y_vals[0], z_vals[0])
			return r, x_vals[0], y_vals[0], z_vals[0] //These are the diffs we can use for the equation
			//Diff = dir*rotated - original
			//Original = dir*rotated - diff
		}
	}
	return rotation{}, -1, -1, -1
}

func allSameValue(a []int) bool {
	for i := 1; i < len(a); i++ {
		if a[i] != a[0] {
			return false
		}
	}
	return true
}

//x is the scanner number in the proper orientation, y is the scanner number in a rotated orientation
//This re-orients a scanner's beacons by applying the correct rotation found using find_rotations
//Uses the formula: Original_orientation = dir*rotated_orientation - diff to convert a rotated point to a properly oriented point
func reorient(scanners *[]scanner, r rotation, y int, x_diff, y_diff, z_diff int) scanner {
	reoriented_beacons := make([]beacon, 0)
	for _, beacon := range (*scanners)[y].beacons {
		new_x := r.x_dir*beacon[r.x] - x_diff
		new_y := r.y_dir*beacon[r.y] - y_diff
		new_z := r.z_dir*beacon[r.z] - z_diff
		// fmt.Println(new_x, new_y, new_z)
		reoriented_beacons = append(reoriented_beacons, [3]int{new_x, new_y, new_z})
	}
	return scanner{reoriented_beacons}
}

//switch oriented to rotated and vice versa
func flip_beacon_pairs(bp []beaconPair) []beaconPair {
	new_pairs := make([]beaconPair, 0)
	for _, b := range bp {
		new_pairs = append(new_pairs, beaconPair{original: b.rotated, rotated: b.original})
	}
	return new_pairs
}
