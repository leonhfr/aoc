package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	sh "github.com/leonhfr/aoc/shared"
)

//go:embed input
var input string

type monkey struct {
	d    string
	a, b string
	op   operation
}

type operation rune

const (
	ADDITION       operation = '+'
	SUBTRACTION    operation = '-'
	MULTIPLICATION operation = '*'
	DIVISION       operation = '/'
	ASSIGNMENT     operation = '='
)

const (
	ROOT  = "root"
	HUMAN = "humn"
)

var lines []string

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}

func part1() int {
	m, queue := parse(input)
	m = solve(m, queue)
	return m[ROOT]
}

func part2() int {
	m, queue := parse(input)
	delete(m, HUMAN)
	queue = correct(queue)

	var q monkey
	for len(queue) > 0 {
		q, queue = queue[0], queue[1:]
		a, aok := m[q.a]
		b, bok := m[q.b]
		d, dok := m[q.d]

		switch {
		case q.op == ASSIGNMENT && aok:
			m[q.b] = a
		case q.op == ASSIGNMENT && bok:
			m[q.a] = b
		case q.op == ADDITION && aok && bok:
			m[q.d] = a + b
		case q.op == ADDITION && bok && dok:
			m[q.a] = d - b
		case q.op == ADDITION && dok && aok:
			m[q.b] = d - a
		case q.op == SUBTRACTION && aok && bok:
			m[q.d] = a - b
		case q.op == SUBTRACTION && bok && dok:
			m[q.a] = d + b
		case q.op == SUBTRACTION && dok && aok:
			m[q.b] = a - d
		case q.op == MULTIPLICATION && aok && bok:
			m[q.d] = a * b
		case q.op == MULTIPLICATION && bok && dok:
			m[q.a] = d / b
		case q.op == MULTIPLICATION && dok && aok:
			m[q.b] = d / a
		case q.op == DIVISION && aok && bok:
			m[q.d] = a / b
		case q.op == DIVISION && bok && dok:
			m[q.a] = d * b
		case q.op == DIVISION && dok && aok:
			m[q.b] = a / d
		default:
			queue = append(queue, q)
		}
	}

	return m[HUMAN]
}

func solve(m map[string]int, queue []monkey) map[string]int {
	var q monkey
	for len(queue) > 0 {
		q, queue = queue[0], queue[1:]
		a, aok := m[q.a]
		b, bok := m[q.b]

		if !aok || !bok {
			queue = append(queue, q)
			continue
		}

		switch q.op {
		case ADDITION:
			m[q.d] = a + b
		case SUBTRACTION:
			m[q.d] = a - b
		case MULTIPLICATION:
			m[q.d] = a * b
		case DIVISION:
			m[q.d] = a / b
		default:
			panic("unknown operation")
		}
	}
	return m
}

func correct(monkeys []monkey) []monkey {
	for i, m := range monkeys {
		if m.d == ROOT {
			monkeys[i].op = ASSIGNMENT
		}
	}
	return monkeys
}

func parse(input string) (map[string]int, []monkey) {
	var monkeys []monkey
	m := make(map[string]int)
	for _, line := range sh.Lines(input) {
		fields := strings.Split(line, ": ")
		n, err := strconv.Atoi(fields[1])
		if err != nil {
			parts := strings.Fields(fields[1])
			monkeys = append(monkeys, monkey{
				fields[0],
				parts[0],
				parts[2],
				operation(parts[1][0]),
			})
			continue
		}

		m[fields[0]] = n
	}
	return m, monkeys
}
