package main

import (
	"bufio"
	"fmt"
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

	// Test Cases

	// bits := HexToBin("D2FE28")
	// bits := HexToBin("38006F45291200")
	// bits := HexToBin("EE00D40C823060")
	// bits := HexToBin("8A004A801A8002F478")
	// bits := HexToBin("620080001611562C8802118E34")
	// bits := HexToBin("C0015000016115A2E0802F182340")
	// bits := HexToBin("A0016C880162017C3686B18A3D4780")

	var version_number_counter int
	process_packet(&bits, 0, &version_number_counter)
	fmt.Println("Part 1:", version_number_counter)
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

//Returns the starting index of the next packet
func process_packet(bits *string, start_index int, version_number_counter *int) int {
	//First three bits are the version number
	*version_number_counter += BinToDec((*bits)[start_index : start_index+3])
	fmt.Println("Version:", BinToDec((*bits)[start_index:start_index+3]))

	//Second three bits are the Packet ID
	packet_type_id := BinToDec((*bits)[start_index+3 : start_index+6])

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
		fmt.Println("Literal Value:", BinToDec(s)) //This is the literal value
		return start_index + 6 + (i * 5)           //Starting index of next packet
	} else if packet_type_id != 4 {
		//If the packet type is NOT 4, this packet is an operator

		//Next Bit is the Length Type ID
		length_type_id := BinToDec((*bits)[start_index+6 : start_index+7])

		//If Zero, then the next 15 bits are a number that represent the total bit
		//length of the sub-packets contained by this packet
		if length_type_id == 0 {
			length_of_subpackets := BinToDec((*bits)[start_index+7 : start_index+22])
			fmt.Println("Length of Subpackets", length_of_subpackets)

			x := start_index + 22
			for x < start_index+22+length_of_subpackets {
				x = process_packet(bits, x, version_number_counter)
			}
			return x
		}
		//If One, the the next 11 bits are a number that represents the number of subpackets immediately
		//contained by this packet
		if length_type_id == 1 {
			number_of_subpackets := BinToDec((*bits)[start_index+7 : start_index+18])
			fmt.Println("Number of Subpackets", number_of_subpackets)

			x := start_index + 18
			for i := 0; i < number_of_subpackets; i++ {
				x = process_packet(bits, x, version_number_counter)
			}

			return x
		}
	}
	return 0
}
