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

type resource uint8

const (
	ORE resource = iota
	CLAY
	OBSIDIAN
	GEODE
)

const (
	N_RESOURCES = 4
	INDEX_TIME  = 8
	TIME_TARGET = 24
)

// cost = robot*4+resource
type blueprint [16]int

var blueprints []blueprint

// triangular numbers
var triangular = [TIME_TARGET + 1]int{
	0, 1, 3, 6, 10, 15,
	21, 28, 36, 45, 55, 66,
	78, 91, 105, 120, 136, 153,
	171, 190, 210, 231, 253, 276,
	300,
}

// state 0-3 resources, 4-7 robots, 8 time
type state [9]int

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}

func part1() int {
	fmt.Println(blueprints)
	s := state{0, 0, 0, 0, 1, 0, 0, 0, 0}
	for r := ORE; r <= GEODE; r++ {
		for _, b := range blueprints {
			fmt.Println(b.makeRobot(r, s))
		}
	}

	for _, b := range blueprints {
		fmt.Println(maxGeodes(b))
	}
	return 0
}

func part2() int {
	return 0
}

func maxGeodes(b blueprint) int {
	var s state
	max := math.MinInt
	stack := []state{{0, 0, 0, 0, 1, 0, 0, 0, 0}}

	for len(stack) > 0 {
		s, stack = stack[len(stack)-1], stack[:len(stack)-1]
		remaining := TIME_TARGET - s[INDEX_TIME]

		// reached target time
		if remaining == 0 {
			// collected more geodes than previous maximum
			if s[GEODE] > max {
				max = s[GEODE]
			}
			fmt.Println("reached target", max, s[GEODE])
			continue
		}

		theoreticalGeodes := s[GEODE] + remaining*s[N_RESOURCES+GEODE] + triangular[remaining]

		// not possible to collect more geodes than the max
		if theoreticalGeodes <= max {
			// fmt.Println("skipped theory")
			continue
		}

		orNext, orOk := b.makeRobot(ORE, s)
		clNext, clOk := b.makeRobot(CLAY, s)
		obNext, obOk := b.makeRobot(OBSIDIAN, s)
		geNext, geOk := b.makeRobot(GEODE, s)

		switch {
		case geOk:
			stack = append(stack, geNext)
		case obOk:
			stack = append(stack, obNext)
		default:
			if (or)
		}

		// for r := ORE; r <= GEODE; r++ {
		// 	next, ok := b.makeRobot(r, s)
		// 	if ok {
		// 		stack = append(stack, next)
		// 	}
		// }

		stack = append(stack, skip(s))

		fmt.Println(len(stack), max, s[INDEX_TIME], s)
	}

	return max
}

// makeRobot returns the state after the next robot creation of the selected resource
// bool is false if the creation is not possible
func (b blueprint) makeRobot(robot resource, s state) (state, bool) {
	var next state
	time := math.MinInt

	for r := ORE; r <= GEODE; r++ {
		need := b[robot*N_RESOURCES+r]
		rate := s[N_RESOURCES+r]

		// no robots to collect the resource
		if need > 0 && rate == 0 {
			return next, false
		}

		if rate == 0 {
			continue
		}

		got := s[r]
		want := need - got
		needTime := int(math.Ceil(float64(want) / float64(rate)))

		// not enough time to create the robot
		if TIME_TARGET-s[INDEX_TIME] < needTime {
			return next, false
		}

		time = sh.Max(time, needTime)
	}

	// update next state
	for r := ORE; r <= GEODE; r++ {
		need := b[robot*N_RESOURCES+r]
		rate := s[N_RESOURCES+r]

		next[r] = s[r] + time*rate - need
		next[N_RESOURCES+r] = s[N_RESOURCES+r]
	}

	next[N_RESOURCES+robot] = s[N_RESOURCES+robot] + 1
	next[INDEX_TIME] = s[INDEX_TIME] + time

	return next, true
}

// skip returns the state at the end of the time target
// without creating any more robots
func skip(s state) state {
	var next state
	for r := ORE; r <= GEODE; r++ {
		next[r] = s[r] + (TIME_TARGET-s[INDEX_TIME])*s[N_RESOURCES+r]
		next[N_RESOURCES+r] = s[N_RESOURCES+r]
	}
	next[INDEX_TIME] = TIME_TARGET
	return next
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
			values[1], 0, 0, 0,
			values[2], 0, 0, 0,
			values[3], values[4], 0, 0,
			values[5], 0, values[6], 0,
		})
	}
}

var test = `Blueprint 1: Each ore robot costs 4 ore. Each clay robot costs 2 ore. Each obsidian robot costs 3 ore and 14 clay. Each geode robot costs 2 ore and 7 obsidian.
Blueprint 2: Each ore robot costs 2 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 8 clay. Each geode robot costs 3 ore and 12 obsidian.`
