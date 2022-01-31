package main

import (
	"container/heap"
	"fmt"
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
var goal_state = State{hallway: Hallway{46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46}, rooms: [4]Room{{65, 65}, {66, 66}, {67, 67}, {68, 68}}, cost: 0}

// var starting_state = State{hallway: Hallway{46, 46, 46, 46, 46, 46, 46, 46, 46, 46, 46}, rooms: [4]Room{{66, 65}, {67, 68}, {66, 67}, {68, 65}}, cost: 0}

//How much energy it costs each pod to move one square
var costs = map[rune]int{65: 1, 66: 10, 67: 100, 68: 1000}

//Maps room to entrance point in hallway
var entrances = map[int]int{0: 2, 1: 4, 2: 6, 3: 8}

//Map rune value of a pod to its home room
var homes = map[rune]int{65: 0, 66: 1, 67: 2, 68: 3}

func main() {
	dists := make(map[State]int, 1)
	dists[starting_state] = 0

	//Create the priority queue for Dijkstra
	pq := make(PriorityQueue, 1)
	//Keep a map for quick state access
	items := make(map[State]*Item, 0)

	it := Item{state: starting_state, priority: 0, index: 0}
	pq[0] = &it
	items[starting_state] = &it

	heap.Init(&pq)

	//Run Dijkstra
	for len(pq) > 0 {
		//Extract Mind
		u := *(heap.Pop(&pq).(*Item))

		//For each neighbor v of u
		neighbors := generate_moves(u.state)
		for _, neighbor := range neighbors {
			//Calculate edge weight distance for this neighbor
			edge_weight := neighbor.cost
			//Now, set the cost to zero, since all states in the map will have 0 cost
			neighbor.cost = 0

			//Now, Gotta check 2 things
			//1. If it's not in the dists map yet, its dist is infinite so we add_with_priority to the queue
			if _, ok := dists[neighbor]; !ok {
				alt := dists[u.state] + edge_weight
				//Create the item
				it := &Item{state: neighbor, priority: alt}
				items[neighbor] = it
				//Append then update
				heap.Push(&pq, it)
				pq.update(it, it.state, alt)
				dists[neighbor] = alt
			} else /* 2. If it is in the dists map, */ {
				//Calculate the new potential distance
				alt := dists[u.state] + edge_weight
				if alt < dists[neighbor] {
					//If it is smaller than its current distances, replace it and update the priority queue
					dists[neighbor] = alt
					pq.update(items[neighbor], neighbor, alt)
				}
			}
		}
	}
	fmt.Println(dists[goal_state])
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

func room_to_room(st State) []State {
	//Collect all new states that can be created by moving ANY pod in this state
	var new_states []State

	for room_index, room := range st.rooms {
	Room_Walk:
		for room_depth, pod := range room {
			//If this is the bottom space in a room and we are blocked, continue
			if room_depth == 1 && room[0] != 46 {
				continue
			}
			//If this spot is empty, continue
			if pod == 46 {
				continue
			}
			//If this pod is already home, continue
			if homes[pod] == room_index {
				continue
			}
			//Check if this pod's home can be moved into
			if !(st.rooms[homes[pod]] == Room{46, 46} || st.rooms[homes[pod]] == Room{46, pod}) {
				continue
			}

			//Check if all spaces along the way are open
			start := entrances[room_index]
			finish := entrances[homes[pod]]
			if start > finish {
				start, finish = finish, start
			}
			for i := start; i <= finish; i++ {
				if st.hallway[i] != 46 { //If this pod can't get home
					continue Room_Walk
				}
			}

			//Now, we know this pod can get home
			new_state := st
			//Remove this pod from its current spot
			new_state.rooms[room_index][room_depth] = 46
			//Place this pod at the correct depth in its room
			correct_depth := 1
			if st.rooms[homes[pod]][1] == pod {
				correct_depth = 0
			}
			new_state.rooms[homes[pod]][correct_depth] = pod
			//Increase the cost
			move_cost := costs[pod] * (absDiffInt(start, finish) + (correct_depth + 1) + (room_depth + 1))
			new_state.cost += move_cost
			//Record this new possible state
			new_states = append(new_states, new_state)
		}
	}
	return new_states
}

func generate_moves(st State) []State {
	var new_states []State

	new_states = append(new_states, room_to_hallway(st)...)
	new_states = append(new_states, hallway_to_room(st)...)
	new_states = append(new_states, room_to_room(st)...)

	return new_states
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func displayState(st State) {
	outString := ""
	for _, v := range st.hallway {
		outString += string(v)
	}
	fmt.Println(outString, st.cost)
}
