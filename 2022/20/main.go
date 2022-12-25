package main

import (
	_ "embed"
	"fmt"

	sh "github.com/leonhfr/aoc/shared"
)

//go:embed input
var input string

const KEY = 811589153

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}

func part1() int {
	numbers, root := parse(input)
	mixAll(numbers)
	return sum(root)
}

func part2() int {
	numbers, root := parse(input)
	for _, n := range numbers {
		n.value *= KEY
	}
	for i := 0; i < 10; i++ {
		mixAll(numbers)
	}
	return sum(root)
}

type number struct {
	value      int
	prev, next *number
}

func move(n *number, d int) *number {
	for ; d < 0; d++ {
		n = n.prev
	}
	for ; d > 0; d-- {
		n = n.next
	}
	return n
}

func mix(n *number, l int) {
	p := n.prev
	n.prev.next = n.next
	n.next.prev = n.prev

	p = move(p, n.value%(l-1))
	n.prev = p
	n.next = p.next
	n.prev.next = n
	n.next.prev = n
}

func mixAll(numbers []*number) {
	for _, n := range numbers {
		mix(n, len(numbers))
	}
}

func sum(root *number) int {
	var s int
	for i, n := 0, root; i < 3; i++ {
		n = move(n, 1000)
		s += n.value
	}
	return s
}

func parse(input string) ([]*number, *number) {
	var numbers []*number
	var root *number

	for _, line := range sh.Lines(input) {
		n := &number{value: sh.ToInt(line)}
		if n.value == 0 {
			root = n
		}
		numbers = append(numbers, n)
	}

	numbers[0].prev = numbers[len(numbers)-1]
	numbers[len(numbers)-1].next = numbers[0]

	for i := 1; i < len(numbers); i++ {
		numbers[i].prev = numbers[i-1]
		numbers[i-1].next = numbers[i]
	}

	return numbers, root
}
