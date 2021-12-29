package main

import (
	"fmt"

	sh "github.com/leonhfr/aoc/shared"
)

type start struct{ player1, player2 int }

var (
	sample = start{4, 8}
	input  = start{10, 6}
)

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}

func part1() int {
	g := deterministicGame(input)
	return sh.Min(g.score1, g.score2) * 3 * g.turns
}

func part2() int {
	win1, win2 := diracGames(input)
	return sh.Max(win1, win2)
}

func deterministicGame(input start) game {
	g, dice := new(input), deterministicDice()
	for g.score1 < 1000 && g.score2 < 1000 {
		a, b, c := dice(), dice(), dice()
		g = nextGame(g, a+b+c, g.turns%2 == 0)
	}
	return g
}

type state struct {
	g game
	n int
}

func diracGames(input start) (win1, win2 int) {
	rolls := rollsMap()
	stack := []state{{new(input), 1}}
	for len(stack) > 0 {
		state := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		switch {
		case state.g.score1 >= 21:
			win1 += state.n
		case state.g.score2 >= 21:
			win2 += state.n
		default:
			stack = append(stack, nextStates(state, rolls)...)
		}
	}
	return
}

func nextStates(s state, rolls map[int]int) []state {
	var states []state
	for roll, n := range rolls {
		states = append(states, state{
			nextGame(s.g, roll, s.g.turns%2 == 0),
			n * s.n,
		})
	}
	return states
}

type game struct {
	player1, player2 int
	score1, score2   int
	turns            int
}

func new(input start) game {
	return game{input.player1, input.player2, 0, 0, 0}
}

func nextGame(g game, rolls int, player1 bool) game {
	if player1 {
		g.player1 = nextPosition(g.player1, rolls)
		g.score1 += g.player1
	} else {
		g.player2 = nextPosition(g.player2, rolls)
		g.score2 += g.player2
	}
	g.turns++
	return g
}

func nextPosition(pos, rolls int) int {
	pos = pos + rolls
	if pos%10 == 0 {
		return 10
	}
	return pos % 10
}

func deterministicDice() func() int {
	roll := 0
	return func() int {
		roll++
		if roll%100 == 0 {
			return 100
		}
		return roll % 100
	}
}

func rollsMap() map[int]int {
	init := []int{1, 2, 3}
	list := combine(combine(init, init), init)
	return sh.IntDict(list)
}

func combine(a, b []int) (c []int) {
	for _, k := range a {
		for _, l := range b {
			c = append(c, k+l)
		}
	}
	return
}
