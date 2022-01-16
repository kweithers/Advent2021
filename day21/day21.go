package main

import "fmt"

//Part 1 stuff
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

func play_game(p1, p2 *player) {
	die := die{} //First roll will be a 1
	p1_turn := true
	for {
		if p1_turn {
			p1.roll_move_score(&die)
			// fmt.Println("Player 1 moved to ", p1.space, "and their score is ", p1.score)
			if p1.score >= 1000 {
				break
			}
		} else {
			p2.roll_move_score(&die)
			// fmt.Println("Player 2 moved to ", p2.space, "and their score is ", p2.score)
			if p2.score >= 1000 {
				break
			}
		}
		p1_turn = !p1_turn
	}
	if p1_turn { //If player 1 won
		fmt.Println("Part 1:", die.times_rolled*p2.score)
	} else { //If player 2 won
		fmt.Println("Part 1:", die.times_rolled*p1.score)
	}
}

//Part 2 stuff
type state struct {
	p1, p2 player
}

type result struct {
	p1_wins int
	p2_wins int
}

type rolls struct {
	first, second, third int
}

func (p *player) roll_move_score_dirac(die_sum int) {
	p.space += die_sum
	p.space %= 10
	if p.space == 0 {
		p.space = 10
	}
	p.score += p.space
}

func play_dirac_game(cache map[state]result, p1, p2 player) result {
	//Check the cache
	if val, ok := cache[state{p1, p2}]; ok {
		return val
	}
	result := result{}

	//Roll the die three times
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			for k := 1; k <= 3; k++ {
				new_p1 := p1
				new_p1.roll_move_score_dirac(i + j + k)
				if new_p1.score >= 21 {
					result.p1_wins += 1
				} else {
					r := play_dirac_game(cache, p2, new_p1)
					result.p1_wins += r.p2_wins
					result.p2_wins += r.p1_wins
				}
			}
		}
	}
	cache[state{p1, p2}] = result
	return result
}

func main() {
	//Part 1
	p1 := player{space: 7}
	p2 := player{space: 1}
	play_game(&p1, &p2)

	//Part 2
	p1 = player{space: 7}
	p2 = player{space: 1}

	cache := make(map[state]result)
	result := play_dirac_game(cache, p1, p2)
	fmt.Println("Part 2:", result)
}
