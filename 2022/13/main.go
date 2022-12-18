package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"sort"

	sh "github.com/leonhfr/aoc/shared"
)

//go:embed input
var input string

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}

func part1() int {
	var count int
	packages := getPackages(input)
	for i := 0; i < len(packages); i += 2 {
		if compare(packages[i], packages[i+1]) <= 0 {
			count += i/2 + 1
		}
	}
	return count
}

func part2() int {
	packages := getPackages(input)
	packages = append(packages,
		[]any{[]any{2.}},
		[]any{[]any{6.}},
	)

	sort.Slice(packages,
		func(i, j int) bool { return compare(packages[i], packages[j]) < 0 },
	)

	result := 1
	for i, p := range packages {
		if fmt.Sprint(p) == "[[2]]" || fmt.Sprint(p) == "[[6]]" {
			result *= i + 1
		}
	}
	return result
}

func compare(a, b any) int {
	as, aok := a.([]any)
	bs, bok := b.([]any)

	switch {
	case !aok && !bok:
		return int(a.(float64) - b.(float64))
	case !aok:
		as = []any{a}
	case !bok:
		bs = []any{b}
	}

	for i := 0; i < len(as) && i < len(bs); i++ {
		if c := compare(as[i], bs[i]); c != 0 {
			return c
		}
	}

	return len(as) - len(bs)
}

func getPackages(input string) []any {
	var packages []any
	for _, line := range sh.Lines(input) {
		if len(line) > 0 {
			var x any
			json.Unmarshal([]byte(line), &x)
			packages = append(packages, x)
		}
	}
	return packages
}
