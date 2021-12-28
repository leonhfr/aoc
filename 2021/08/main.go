package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"

	sh "github.com/leonhfr/aoc/shared"
)

//go:embed input
var input string

var entries []entry

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}

func part1() int {
	var count int
	for _, e := range entries {
		for _, o := range e.output {
			l := len(o)
			if l == 2 || l == 4 || l == 3 || l == 7 {
				count++
			}
		}
	}
	return count
}

func part2() int {
	var sum int
	for _, e := range entries {
		sum += e.crack()
	}
	return sum
}

func (e entry) crack() int {
	knowns := map[int]int{2: 1, 4: 4, 3: 7, 7: 8}
	intToString := make(map[int]string)
	stringToInt := make(map[string]string)
	var unknowns []string

	for _, v := range e.patterns {
		s, ok := knowns[len(v)]
		if ok {
			intToString[s] = v
			stringToInt[v] = strconv.Itoa(s)
		} else {
			unknowns = append(unknowns, v)
		}
	}

	for _, c := range unknowns {
		switch len(c) {
		case 5:
			if contains(c, intToString[1]) {
				intToString[3] = c
				stringToInt[c] = "3"
			} else if contains(c, difference(intToString[4], intToString[1])) {
				intToString[5] = c
				stringToInt[c] = "5"
			} else {
				intToString[2] = c
				stringToInt[c] = "2"
			}
		case 6:
			if contains(c, intToString[4]) {
				intToString[9] = c
				stringToInt[c] = "9"
			} else if !contains(c, intToString[1]) {
				intToString[6] = c
				stringToInt[c] = "6"
			} else {
				intToString[0] = c
				stringToInt[c] = "0"
			}
		}
	}

	var res string
	for _, o := range e.output {
		res += stringToInt[o]
	}
	return sh.ToInt(res)
}

func difference(a, b string) string {
	var c []rune
	for _, r := range a {
		if !strings.ContainsRune(b, r) {
			c = append(c, r)
		}
	}
	return string(c)
}

func contains(s, t string) bool {
	for _, r := range t {
		if !strings.ContainsRune(s, r) {
			return false
		}
	}
	return true
}

type entry struct {
	patterns, output []string
}

func init() {
	lines := sh.Lines(input)
	for _, line := range lines {
		s := strings.Split(line, " | ")
		patterns := sortStrings(strings.Split(s[0], " "))
		output := sortStrings(strings.Split(s[1], " "))
		entries = append(entries, entry{patterns, output})
	}
}

func sortStrings(input []string) (output []string) {
	for _, i := range input {
		r := runes(i)
		sort.Sort(r)
		output = append(output, string(r))
	}
	return
}

type runes []rune

func (r runes) Len() int           { return len(r) }
func (r runes) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r runes) Less(i, j int) bool { return r[i] < r[j] }
