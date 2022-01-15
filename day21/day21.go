package main

import "fmt"

type player struct {
	score int
	space int //1-10 are the valid spaces
}

func (p *player) roll_move_score(die *die) {
	s := 0
	for i := 0; i < 3; i++ {
		s += die.roll()
	}
	p.space += s
	p.space %= 10

	if p.space == 0 {
		p.space = 10
	}

	p.score += p.space
}

type die struct {
	number       int //deterministic die starts at 1 and goes to 100
	times_rolled int
}

func (d *die) roll() int {
	d.times_rolled++
	d.number++
	if d.number > 100 {
		d.number = 1
	}
	return d.number
}

func main() {
	p1 := player{score: 0, space: 7}
	p2 := player{score: 0, space: 1}

	die := die{} //First roll will be a 1
	p1_turn := true
	for {
		if p1_turn {
			p1.roll_move_score(&die)
			fmt.Println("Player 1 moved to ", p1.space, "and their score is ", p1.score)
			if p1.score >= 1000 {
				break
			}
		} else {
			p2.roll_move_score(&die)
			fmt.Println("Player 2 moved to ", p2.space, "and their score is ", p2.score)
			if p2.score >= 1000 {
				break
			}
		}
		p1_turn = !p1_turn
	}

	fmt.Println(die.times_rolled)
	if p1_turn { //If player 1 won
		fmt.Println(die.times_rolled * p2.score)
	} else { //If player 2 won
		fmt.Println(die.times_rolled * p1.score)
	}
}
