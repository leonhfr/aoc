package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strings"

	sh "github.com/leonhfr/aoc/shared"
)

//go:embed input
var input string

var polymer state

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}

func part1() int {
	for i := 0; i < 10; i++ {
		polymer = polymer.step()
	}
	return polymer.statistics()
}

func part2() int {
	for i := 0; i < 40; i++ {
		polymer = polymer.step()
	}
	return polymer.statistics()
}

type state struct {
	counter map[string]int
	pairs   map[string]int
	rules   map[string]string
}

func (s state) step() state {
	pairs := make(map[string]int)
	for pair, count := range s.pairs {
		ch := s.rules[pair]
		a, b := pair[0:1]+ch, ch+pair[1:2]
		s.counter[ch] += count
		pairs[a] += count
		pairs[b] += count
	}
	return state{
		counter: s.counter,
		pairs:   pairs,
		rules:   s.rules,
	}
}

func (s state) statistics() int {
	var counts []int
	for _, v := range s.counter {
		counts = append(counts, v)
	}
	sort.Ints(counts)
	return counts[len(counts)-1] - counts[0]
}

func init() {
	lines := sh.Lines(input)
	counter, pairs := parseTemplate(lines[0])
	rules := parseRules(lines[2:])
	polymer = state{counter, pairs, rules}
}

func parseTemplate(template string) (map[string]int, map[string]int) {
	counter, pairs := make(map[string]int), make(map[string]int)
	for i := 0; i < len(template); i++ {
		counter[template[i:i+1]]++

		if i != 0 {
			pairs[template[i-1:i+1]]++
		}
	}
	return counter, pairs
}

func parseRules(lines []string) map[string]string {
	rules := make(map[string]string)
	for _, l := range lines {
		elements := strings.Split(l, " -> ")
		rules[elements[0]] = elements[1]
	}
	return rules
}
