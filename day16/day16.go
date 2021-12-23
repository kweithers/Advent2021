package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("day16.txt")
	fscanner := bufio.NewScanner(file)

	bits := ""

	for fscanner.Scan() {
		line := strings.Split(fscanner.Text(), "")
		for _, c := range line {
			bits += HexToBin(c)
		}
	}

	var version_number_counter int
	_, part2 := process_packet(&bits, 0, &version_number_counter)
	fmt.Println("Part 1:", version_number_counter)
	fmt.Println("Part 2:", part2)
}

func HexToBin(hex string) string {
	s := ""
	for _, c := range hex {
		ui, _ := strconv.ParseUint(string(c), 16, 64)
		// %04b indicates base 2, zero padded, with 4 characters
		s += fmt.Sprintf("%04b", ui)
	}
	return s
}

func BinToDec(bin string) int {
	ui, _ := strconv.ParseUint(bin, 2, 64)
	return int(ui)
}

//Returns (the starting index of the next packet, value of the packet)
func process_packet(bits *string, start_index int, version_number_counter *int) (int, int) {
	//First three bits are the version number
	*version_number_counter += BinToDec((*bits)[start_index : start_index+3])

	//Second three bits are the packet type id
	packet_type_id := BinToDec((*bits)[start_index+3 : start_index+6])

	//Keep track of our place in the bit string
	var x int
	//Record the value of a packet
	var value int
	//Record the values of the sub packets
	sub_packet_values := make([]int, 0)
	//If the packet type id is 4, this is a literal value
	if packet_type_id == 4 {
		i := 0
		s := ""
		for {
			//Append the last four bits
			s += (*bits)[start_index+7+(i*5) : start_index+11+(i*5)]
			//If the first bit is one, this is not the last group
			if (*bits)[start_index+6+(i*5):start_index+7+(i*5)] == "1" {
				i++
			} else { //If the first bit is zero, this it the last group
				i++
				break
			}
		}
		return start_index + 6 + (i * 5), BinToDec(s) //Starting index of next packet, literal value
	} else if packet_type_id != 4 {
		//If the packet type is NOT 4, this packet is an operator

		//Next Bit is the Length Type ID
		length_type_id := BinToDec((*bits)[start_index+6 : start_index+7])

		//If Zero, then the next 15 bits are a number that represent the total bit
		//length of the sub-packets contained by this packet
		if length_type_id == 0 {
			length_of_subpackets := BinToDec((*bits)[start_index+7 : start_index+22])

			x = start_index + 22
			for x < start_index+22+length_of_subpackets {
				x, value = process_packet(bits, x, version_number_counter)
				sub_packet_values = append(sub_packet_values, value)
			}
		}
		//If One, the the next 11 bits are a number that represents the number of subpackets immediately
		//contained by this packet
		if length_type_id == 1 {
			number_of_subpackets := BinToDec((*bits)[start_index+7 : start_index+18])

			x = start_index + 18
			for i := 0; i < number_of_subpackets; i++ {
				x, value = process_packet(bits, x, version_number_counter)
				sub_packet_values = append(sub_packet_values, value)
			}
		}
	}

	//Now, process the contents of this packet and apply the correct operation
	if packet_type_id == 0 { //Sum
		sum := 0
		for _, v := range sub_packet_values {
			sum += v
		}
		return x, sum
	} else if packet_type_id == 1 { //Product
		prod := 1
		for _, v := range sub_packet_values {
			prod *= v
		}
		return x, prod
	} else if packet_type_id == 2 { //Minimum
		min := math.MaxInt
		for _, v := range sub_packet_values {
			if v < min {
				min = v
			}
		}
		return x, min
	} else if packet_type_id == 3 { //Maximum
		max := math.MinInt
		for _, v := range sub_packet_values {
			if v > max {
				max = v
			}
		}
		return x, max
	} else if packet_type_id == 5 { //Greater than
		v := 0
		if sub_packet_values[0] > sub_packet_values[1] {
			v = 1
		}
		return x, v
	} else if packet_type_id == 6 { //Less than
		v := 0
		if sub_packet_values[0] < sub_packet_values[1] {
			v = 1
		}
		return x, v
	} else if packet_type_id == 7 { //Equal to
		v := 0
		if sub_packet_values[0] == sub_packet_values[1] {
			v = 1
		}
		return x, v
	}

	return x, -1 //Will never reach here
}
