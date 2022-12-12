package main

import (
	"fmt"
	"sort"
)

type monkey struct {
	items       []int
	inspections int
	operation   func(int) int
	test        [3]int
}

func (m *monkey) throwTo(worry int) int {
	if worry%m.test[0] == 0 {
		return m.test[1]
	}
	return m.test[2]
}

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}

func part1() int {
	troop := newInputTroop()
	for i := 0; i < 20; i++ {
		for m := range troop {
			tot := len(troop[m].items)
			for index := 0; index < tot; index++ {
				troop[m].inspections++
				item := troop[m].items[0]
				troop[m].items = troop[m].items[1:]
				worry := troop[m].operation(item) / 3
				next := troop[m].throwTo(worry)
				troop[next].items = append(troop[next].items, worry)
			}
		}
	}
	return inspections(troop)
}

const LCM = 2 * 3 * 5 * 7 * 11 * 13 * 17 * 19

func part2() int {
	troop := newInputTroop()
	for i := 0; i < 10000; i++ {
		for m := range troop {
			troop[m].inspections += len(troop[m].items)
			for _, item := range troop[m].items {
				worry := troop[m].operation(item) % LCM
				next := troop[m].throwTo(worry)
				troop[next].items = append(troop[next].items, worry)
			}
			troop[m].items = []int{}
		}
	}
	return inspections(troop)
}

func inspections(troop []monkey) int {
	var inspections []int
	for _, monkey := range troop {
		inspections = append(inspections, monkey.inspections)
	}
	sort.Ints(inspections)
	return inspections[len(inspections)-1] * inspections[len(inspections)-2]
}

func newInputTroop() []monkey {
	return []monkey{
		{test: [3]int{13, 4, 7}, operation: func(i int) int { return 11 * i }, items: []int{63, 84, 80, 83, 84, 53, 88, 72}},
		{test: [3]int{11, 5, 3}, operation: func(i int) int { return i + 4 }, items: []int{67, 56, 92, 88, 84}},
		{test: [3]int{2, 3, 1}, operation: func(i int) int { return i * i }, items: []int{52}},
		{test: [3]int{5, 5, 6}, operation: func(i int) int { return i + 2 }, items: []int{59, 53, 60, 92, 69, 72}},
		{test: [3]int{7, 7, 2}, operation: func(i int) int { return i + 3 }, items: []int{61, 52, 55, 61}},
		{test: [3]int{3, 0, 6}, operation: func(i int) int { return i + 1 }, items: []int{79, 53}},
		{test: [3]int{19, 4, 0}, operation: func(i int) int { return i + 5 }, items: []int{59, 86, 67, 95, 92, 77, 91}},
		{test: [3]int{17, 2, 1}, operation: func(i int) int { return 19 * i }, items: []int{58, 83, 89}},
	}
}
