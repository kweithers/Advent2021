package main

import (
	"fmt"
	"math"
)

//Rune values of A,B,C,D,. : 65,66,67,68,46

//The position of the pods
type Hallway [11]rune
type Room [2]rune

//A state is the positions of the pods, with how much energy it took to get there
type State struct {
	hallway Hallway
	rooms   [4]Room
	cost    int
}

/* ********************** Rules **********************
1. Can not stop at positions 2,4,6,8: These are the spaces immediately outside a room
2. Can only move into a room if a. both spots are empty, b. bottom spot is filled with correct pod and top spot is empty
3. Only three types of moves: room to hallway, hallway to room, room to room (i.e. no hallway to hallway!)
*/

var starting_state = State{hallway: Hallway{46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46}, rooms: [4]Room{{68, 68}, {65, 67}, {67, 66}, {65, 66}}, cost: 0}
var goal_state = State{hallway: Hallway{46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46}, rooms: [4]Room{{65, 65}, {66, 66}, {67, 67}, {68, 68}}, cost: math.MaxInt}
var test_state = State{hallway: Hallway{66, 46, 46, 46, 46, 66, 46, 46, 46, 46, 46}, rooms: [4]Room{{65, 65}, {46, 46}, {67, 67}, {68, 68}}, cost: 0}

//How much energy it costs each pod to move one square
var costs = map[rune]int{65: 1, 66: 10, 67: 100, 68: 1000}

//Maps room to entrance point in hallway
var entrances = map[int]int{0: 2, 1: 4, 2: 6, 3: 8}

//Map rune value of a pod to its home room
var homes = map[rune]int{65: 0, 66: 1, 67: 2, 68: 3}

func main() {
	x := hallway_to_room(test_state)
	fmt.Println("START")
	for _, st := range x {
		displayState(st)
	}

}

func room_to_hallway(st State) []State {
	//Collect all new states that can be created by moving ANY pod in this state
	var new_states []State

	for room_index, room := range st.rooms {
		for room_depth, pod := range room {
			//If this is the bottom space in a room and we are blocked, continue
			if room_depth == 1 && room[0] != 46 {
				continue
			}
			//If this spot is empty, continue
			if pod == 46 {
				continue
			}

			//Check how far we can go in both directions
			// 1. Go left, starting at the room entrance
			for i := entrances[room_index]; i >= 0; i-- {
				//if we are directly outside a room, continue since we can't stop there
				if i == 2 || i == 4 || i == 6 || i == 8 {
					continue
				}
				//if this spot is occupied by another pod, break out of the loop
				if st.hallway[i] != 46 {
					break
				}
				//If we can move to this spot, Create the new state
				new_state := st
				//Remove the pod from the room
				new_state.rooms[room_index][room_depth] = 46
				//Place the pod at this location
				new_state.hallway[i] = pod
				//Increase the cost
				move_cost := costs[pod] * (absDiffInt(entrances[room_index], i) + (room_depth + 1))
				new_state.cost += move_cost
				//Record this new possible state
				new_states = append(new_states, new_state)
			}

			// 2. Go right, starting at the room entrance
			for i := entrances[room_index]; i <= 10; i++ {
				//if we are directly outside a room, continue since we can't stop there
				if i == 2 || i == 4 || i == 6 || i == 8 {
					continue
				}
				//if this spot is occupied by another pod, break out of the loop
				if st.hallway[i] != 46 {
					break
				}
				//If we can move to this spot, Create the new state
				new_state := st
				//Remove the pod from the room
				new_state.rooms[room_index][room_depth] = 46
				//Place the pod at this location
				new_state.hallway[i] = pod
				//Increase the cost
				move_cost := costs[pod] * (absDiffInt(entrances[room_index], i) + (room_depth + 1))
				new_state.cost += move_cost
				//Record this new possible state
				new_states = append(new_states, new_state)
			}
		}
	}
	return new_states
}

func hallway_to_room(st State) []State {
	var new_states []State

Hallway_Walk:
	for hallway_index, pod := range st.hallway {
		//If this spot is empty, continue
		if pod == 46 {
			continue
		}
		//Check if this pod's home can be moved into
		if !(st.rooms[homes[pod]] == Room{46, 46} || st.rooms[homes[pod]] == Room{46, pod}) {
			continue
		}

		//Check if all spaces along the way are open
		start := hallway_index
		finish := entrances[homes[pod]]
		fmt.Println(start, finish)
		if start > finish {
			start, finish = finish, start
		}
		for i := start; i <= finish; i++ {
			if st.hallway[i] != 46 && i != hallway_index { //If this pod can't get home
				continue Hallway_Walk
			}
		}

		//Now, we know this pod can get home
		new_state := st
		//Remove this pod from the hallway
		new_state.hallway[hallway_index] = 46
		//Place this pod at the correct depth in its room
		correct_depth := 1
		if st.rooms[homes[pod]][1] == pod {
			correct_depth = 0
		}
		new_state.rooms[homes[pod]][correct_depth] = pod
		//Increase the cost
		move_cost := costs[pod] * (absDiffInt(entrances[homes[pod]], hallway_index) + (correct_depth + 1))
		new_state.cost += move_cost
		//Record this new possible state
		new_states = append(new_states, new_state)
	}
	return new_states
}

// func room_to_room(st State) []State {}
// func generate_moves(State) []State {}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func displayState(st State) {
	// don't need index so _
	outString := ""
	for _, v := range st.hallway {
		outString += string(v)
	}
	fmt.Println(outString, st.cost)
}
