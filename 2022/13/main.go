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
		[]interface{}{[]interface{}{2.}},
		[]interface{}{[]interface{}{6.}},
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

func compare(a, b interface{}) int {
	as, aok := a.([]interface{})
	bs, bok := b.([]interface{})

	switch {
	case !aok && !bok:
		return int(a.(float64) - b.(float64))
	case !aok:
		as = []interface{}{a}
	case !bok:
		bs = []interface{}{b}
	}

	for i := 0; i < len(as) && i < len(bs); i++ {
		if c := compare(as[i], bs[i]); c != 0 {
			return c
		}
	}

	return len(as) - len(bs)
}

func getPackages(input string) []interface{} {
	var packages []interface{}
	for _, line := range sh.Lines(input) {
		if len(line) > 0 {
			var x interface{}
			json.Unmarshal([]byte(line), &x)
			packages = append(packages, x)
		}
	}
	return packages
}

// func init() {
// 	for _, p := range strings.Split(input, "\n\n") {
// 		values := strings.Split(p, "\n")
// 		var left, right interface{}
// 		json.Unmarshal([]byte(values[0]), &left)
// 		json.Unmarshal([]byte(values[1]), &right)
// 		pairs = append(pairs, pair{left, right})
// 	}
// }

var test = `[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]`
