package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strings"

	sh "github.com/leonhfr/aoc/shared"
)

//go:embed input
var input string

type valve struct {
	name  string
	rate  int
	edges []string
}

func (v *valve) String() string {
	return fmt.Sprintf("%s", v.name)
}

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}

func part1() int {
	valves := parse(test)
	fmt.Println(valves)
	return 0
}

func part2() int {
	return 0
}

func parse(input string) []*valve {
	var valves []*valve
	regex := regexp.MustCompile(`^Valve ([A-Z]{2}) has flow rate=(\d+); tunnels? leads? to valves? (.*)$`)
	for _, line := range sh.Lines(input) {
		if regex.MatchString(line) {
			fields := regex.FindStringSubmatch(line)
			valves = append(valves, &valve{
				name:  fields[1],
				rate:  sh.ToInt(fields[2]),
				edges: strings.Split(fields[3], ", "),
			})
		}
	}
	return valves
}

var test = `Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
Valve BB has flow rate=13; tunnels lead to valves CC, AA
Valve CC has flow rate=2; tunnels lead to valves DD, BB
Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE
Valve EE has flow rate=3; tunnels lead to valves FF, DD
Valve FF has flow rate=0; tunnels lead to valves EE, GG
Valve GG has flow rate=0; tunnels lead to valves FF, HH
Valve HH has flow rate=22; tunnel leads to valve GG
Valve II has flow rate=0; tunnels lead to valves AA, JJ
Valve JJ has flow rate=21; tunnel leads to valve II`
