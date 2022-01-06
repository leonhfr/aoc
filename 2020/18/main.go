package main

import (
	_ "embed"
	"fmt"
	"strings"

	sh "github.com/leonhfr/aoc/shared"
)

//go:embed input
var input string

var expressions [][]string

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}

func part1() int {
	r := 0
	for _, e := range expressions {
		r += calc(e, calc1)
	}
	return r
}

func part2() int {
	r := 0
	for _, e := range expressions {
		r += calc(e, calc2)
	}
	return r
}

func calc(e []string, f func([]string) string) int {
	var stack []int
	var s int
	var result []string
	for i := range e {
		result = append(result, e[i])
		switch e[i] {
		case "(":
			stack = append(stack, len(result)-1)
		case ")":
			s, stack = stack[len(stack)-1], stack[:len(stack)-1]
			result[s] = f(result[s+1 : len(result)-1])
			result = result[:s+1]
		}
	}
	return sh.ToInt(f(result))
}

func calc1(e []string) string {
	r := sh.ToInt(e[0])
	for i := range e {
		if i+2 < len(e) {
			switch e[i+1] {
			case "+":
				r += sh.ToInt(e[i+2])
			case "*":
				r *= sh.ToInt(e[i+2])
			}
		}
	}
	return fmt.Sprint(r)
}

func calc2(e []string) string {
	for i := 1; i < len(e)-1; i++ {
		if e[i] == "+" {
			e[i-1] = fmt.Sprint(sh.ToInt(e[i-1]) + sh.ToInt(e[i+1]))
			copy(e[i:], e[i+2:])
			e = e[:len(e)-2]
			i--
		}
	}

	for i := 1; i < len(e)-1; i++ {
		if e[i] == "*" {
			e[i-1] = fmt.Sprint(sh.ToInt(e[i-1]) * sh.ToInt(e[i+1]))
			copy(e[i:], e[i+2:])
			e = e[:len(e)-2]
			i--
		}
	}

	return e[0]
}

func init() {
	for _, line := range sh.Lines(input) {
		expressions = append(expressions, strings.Split(strings.ReplaceAll(line, " ", ""), ""))
	}
}
