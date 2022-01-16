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
	p1_pos, p2_pos, p1_score, p2_score int
}

type result struct {
	p1_wins int
	p2_wins int
}

type rolls struct {
	first, second, third int
}

func roll_move_dirac(space, score, die_sum int) int { //returns the landing square
	space += die_sum
	space %= 10
	if space == 0 {
		space = 10
	}
	return space
}
func play_dirac_game(cache map[state]result, r []rolls, p1_pos, p2_pos, p1_score, p2_score int) result {
	//Check the cache
	if val, ok := cache[state{p1_pos, p2_pos, p1_score, p2_score}]; ok {
		return val
	}

	result := result{}
	for _, roll := range r {
		new_pos_1 := roll_move_dirac(p1_pos, p1_score, roll.first+roll.second+roll.third)
		new_score_1 := p1_score + new_pos_1
		if new_score_1 >= 21 {
			result.p1_wins += 1
		} else {
			r := play_dirac_game(cache, r, p2_pos, new_pos_1, p2_score, new_score_1)
			result.p1_wins += r.p2_wins
			result.p2_wins += r.p1_wins
		}
	}
	cache[state{p1_pos, p2_pos, p1_score, p2_score}] = result
	return result
}

func main() {
	//Part 1
	p1 := player{space: 7}
	p2 := player{space: 1}
	play_game(&p1, &p2)

	//Part 2
	p1_pos := 7
	p2_pos := 1
	p1_score := 0
	p2_score := 0
	rolls := []rolls{
		{1, 1, 1}, {1, 1, 2}, {1, 1, 3}, {1, 2, 1}, {1, 2, 2}, {1, 2, 3}, {1, 3, 1}, {1, 3, 2}, {1, 3, 3},
		{2, 1, 1}, {2, 1, 2}, {2, 1, 3}, {2, 2, 1}, {2, 2, 2}, {2, 2, 3}, {2, 3, 1}, {2, 3, 2}, {2, 3, 3},
		{3, 1, 1}, {3, 1, 2}, {3, 1, 3}, {3, 2, 1}, {3, 2, 2}, {3, 2, 3}, {3, 3, 1}, {3, 3, 2}, {3, 3, 3}}

	cache := make(map[state]result)
	result := play_dirac_game(cache, rolls, p1_pos, p2_pos, p1_score, p2_score)
	fmt.Println("Part 2:", result)
}
