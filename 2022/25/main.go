package main

import (
	_ "embed"
	"fmt"

	sh "github.com/leonhfr/aoc/shared"
)

//go:embed input
var input string

var lines []string

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}

func part1() string {
	var sum snafu
	for _, req := range parse(input) {
		sum.val += req.val
	}
	return sum.toSnafu()
}

func part2() int {
	return 0
}

func parse(input string) []snafu {
	var reqs []snafu
	for _, line := range sh.Lines(input) {
		reqs = append(reqs, parseSnafu(line))
	}
	return reqs
}

type snafu struct {
	val int
}

func parseSnafu(str string) snafu {
	var s snafu
	for i := 0; i < len(str); i++ {
		switch p := sh.Pow(5, i); str[len(str)-i-1] {
		case '=':
			s.val -= 2 * p
		case '-':
			s.val -= p
		case '0':
		case '1':
			s.val += p
		case '2':
			s.val += 2 * p
		default:
			panic("unknown rune")
		}
	}
	return s
}

func (s snafu) toSnafu() string {
	var carry int
	var res string
	for r := s.val; r != 0; r = r/5 + carry {
		carry = 0
		switch r % 5 {
		case 0:
			res = "0" + res
		case 1:
			res = "1" + res
		case 2:
			res = "2" + res
		case 3:
			res = "=" + res
			carry = 1
		case 4:
			res = "-" + res
			carry = 1
		}
	}
	return res
}

func (s snafu) toDecimal() int {
	return s.val
}
