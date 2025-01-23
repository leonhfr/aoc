package main

import (
	_ "embed"
	"fmt"
	"math"
	"regexp"

	sh "github.com/leonhfr/aoc/shared"
)

//go:embed input
var input string

const TARGET = 24

type resource uint8

const (
	ORE resource = iota
	CLAY
	OBSIDIAN
	GEODE
)

type blueprint [4][4]int

var blueprints []blueprint

type state [8]int // 0-3 robots 4-7 resources

func next(s state, b blueprint) []state {
	states := []state{
		{
			s[0], s[1], s[2], s[3],
			s[0] + s[4], s[1] + s[5], s[2] + s[6], s[3] + s[7],
		},
	}

	for buy := 0; buy < 4; buy++ {
		r1 := s[4] - b[buy][0]
		r2 := s[5] - b[buy][1]
		r3 := s[6] - b[buy][2]
		r4 := s[7] - b[buy][3]

		if sh.Min(r1, r2, r3, r4) < 0 {
			continue
		}

		n := state{
			s[0], s[1], s[2], s[3],
			r1, r2, r3, r4,
		}
		n[buy]++
		states = append(states, n)
	}

	return states
}

func maxGeodes(b blueprint) int {
	states := []state{{1}}

	for i := 0; i < TARGET; i++ {
		var n []state
		for _, s := range states {
			n = append(n, next(s, b)...)
		}
		states = n
	}

	max := math.MinInt
	for _, s := range states {
		max = sh.Max(max, s[7])
	}
	return max
}

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}

func part1() int {
	var quality int
	for i, b := range blueprints {
		max := maxGeodes(b)
		quality += (i + 1) * max
		fmt.Println(quality)
	}
	return quality
}

func part2() int {
	return 0
}

func init() {
	regex := regexp.MustCompile(`(\d+)`)
	for _, line := range sh.Lines(test) {
		fields := regex.FindAllStringSubmatch(line, 7)
		var values []int
		for _, field := range fields {
			values = append(values, sh.ToInt(field[0]))
		}

		blueprints = append(blueprints, blueprint{
			{values[1], 0, 0, 0},
			{values[2], 0, 0, 0},
			{values[3], values[4], 0, 0},
			{values[5], 0, values[6], 0},
		})
	}
}

var test = `Blueprint 1: Each ore robot costs 4 ore. Each clay robot costs 2 ore. Each obsidian robot costs 3 ore and 14 clay. Each geode robot costs 2 ore and 7 obsidian.
Blueprint 2: Each ore robot costs 2 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 8 clay. Each geode robot costs 3 ore and 12 obsidian.`

// func mix(tuples []tuple) []tuple {
// 	for i := 0; i < len(tuples); i++ {
// 		index := originalIndex(tuples, i)
// 		value := tuples[index]
// 		next := modulo(index+value.number, len(tuples))
// 		fmt.Println("next", i, index, value, next)

// 		removed := append(tuples[:index], tuples[index+1:]...)
// 		tuples = append(removed[:next], append([]tuple{value}, removed[next:]...)...)
// 		fmt.Println(tuples)
// 	}
// 	return tuples
// }

// func modulo(d, m int) int {
// 	v := d % m
// 	if v < 0 {
// 		return v + m
// 	}
// 	return v
// }

// func coordinates(tuples []tuple) int {
// 	var sum int
// 	o := origin(tuples)
// 	l := len(tuples)
// 	for _, index := range indices {
// 		a := index % l
// 		if o+a < l {
// 			sum += tuples[o+a].number
// 		} else {
// 			sum += tuples[a-(l-o)].number
// 		}
// 	}
// 	return sum
// }

// func originalIndex(tuples []tuple, index int) int {
// 	for index, tuple := range tuples {
// 		if tuple.index == index {
// 			return index
// 		}
// 	}
// 	fmt.Println(index)
// 	panic("index not found")
// }

// func origin(tuples []tuple) int {
// 	for index, tuple := range tuples {
// 		if tuple.number == ORIGIN {
// 			return index
// 		}
// 	}
// 	panic("origin not found")
// }

// func parse(input string) []tuple {
// 	var tuples []tuple
// 	for index, line := range sh.Lines(input) {
// 		tuples = append(tuples, tuple{index, sh.ToInt(line)})
// 	}
// 	return tuples
// }
