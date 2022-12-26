package main

import (
	_ "embed"
	"fmt"

	set "github.com/leonhfr/aoc/shared/setpoint"
)

//go:embed input
var input string

var lines []string

type direction uint8

// from, to
type move [2]set.Point

const (
	N direction = 1 << iota
	S
	W
	E

	ALL  direction = N ^ S ^ W ^ E
	NONE direction = 0
)

var order = [4]direction{N, S, W, E}

var directions = [4][3]set.Point{
	{{-1, -1}, {0, -1}, {1, -1}},
	{{-1, 1}, {0, 1}, {1, 1}},
	{{-1, 1}, {-1, 0}, {-1, -1}},
	{{1, 1}, {1, 0}, {1, -1}},
}

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}

func part1() int {
	elves := set.ParseSharp(input)
	for i := 0; i < 10; i++ {
		moves := proposals(elves, i)
		doMoves(elves, moves)
	}
	return emptyGround(elves)
}

func part2() int {
	elves := set.ParseSharp(input)
	for i := 0; ; i++ {
		moves := proposals(elves, i)
		moved := doMoves(elves, moves)
		if moved == 0 {
			return i + 1
		}
	}
}

func emptyGround(elves *set.SetPoint) int {
	min, max := elves.Boundaries()
	return (max.X-min.X+1)*(max.Y-min.Y+1) - elves.Len()
}

func doMoves(elves *set.SetPoint, moves []move) int {
	var moved int
	dict := make(map[set.Point]int)
	for _, move := range moves {
		dict[move[1]]++
	}
	for _, move := range moves {
		if dict[move[1]] == 1 {
			elves.Del(move[0])
			elves.Add(move[1])
			moved++
		}
	}
	return moved
}

func proposals(elves *set.SetPoint, turn int) []move {
	var moves []move
	for _, elf := range elves.Points() {
		can := canMove(elf, elves)
		if can == ALL || can == NONE {
			continue
		}

		var dir direction
		for i := 0; i < 4; i++ {
			if d := order[(turn+i)%4]; d&can > 0 {
				dir = d
				break
			}
		}

		switch dir {
		case N:
			moves = append(moves, move{elf, set.Point{elf.X, elf.Y - 1}})
		case S:
			moves = append(moves, move{elf, set.Point{elf.X, elf.Y + 1}})
		case W:
			moves = append(moves, move{elf, set.Point{elf.X - 1, elf.Y}})
		case E:
			moves = append(moves, move{elf, set.Point{elf.X + 1, elf.Y}})
		}
	}
	return moves
}

func canMove(elf set.Point, elves *set.SetPoint) direction {
	var can direction
	for i, dir := range order {
		var occupied bool
		for _, v := range directions[i] {
			if elves.Has(set.Point{
				X: elf.X + v.X,
				Y: elf.Y + v.Y,
			}) {
				occupied = true
			}
		}

		if !occupied {
			can ^= dir
		}
	}
	return can
}
