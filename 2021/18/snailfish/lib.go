package snailfish

import (
	"fmt"

	sh "github.com/leonhfr/aoc/shared"
)

type SnailFish struct {
	value       *int
	left, right *SnailFish
}

func New(str string) *SnailFish {
	sf, _ := parse(str, 0)
	return sf
}

func Add(a, b *SnailFish) *SnailFish {
	n := &SnailFish{nil, a, b}
	for reduce := true; reduce; {
		if n.Explode() {
			continue
		}

		reduce = n.Split()
	}
	return n
}

func (sf *SnailFish) Explode() bool {
	if explode(1, sf.left, nil, sf.right) {
		return true
	}

	return explode(1, sf.right, sf.left, nil)
}

func explode(n int, sf, left, right *SnailFish) bool {
	if sf.value != nil {
		return false
	}

	if n == 4 {
		l, r := sf.left.value, sf.right.value
		if l == nil || r == nil {
			panic("Exploding pair does not consist of two regular numbers.")
		}

		if left != nil {
			rm := rightMost(left)
			tmp := *rm.value + *l
			rm.value = &tmp
		}

		if right != nil {
			lm := leftMost(right)
			tmp := *lm.value + *r
			lm.value = &tmp
		}

		zero := 0
		sf.value = &zero
		sf.left = nil
		sf.right = nil

		return true
	}

	if explode(n+1, sf.left, left, sf.right) {
		return true
	}

	return explode(n+1, sf.right, sf.left, right)
}

func leftMost(sf *SnailFish) *SnailFish {
	if sf.value != nil {
		return sf
	}

	return leftMost(sf.left)
}

func rightMost(sf *SnailFish) *SnailFish {
	if sf.value != nil {
		return sf
	}

	return rightMost(sf.right)
}

func (sf *SnailFish) Split() bool {
	stack := []*SnailFish{sf}
	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if current.value != nil && *current.value > 9 {
			l, r := *current.value/2, *current.value-*current.value/2
			left, right := &SnailFish{&l, nil, nil}, &SnailFish{&r, nil, nil}
			new := SnailFish{nil, left, right}
			*current = new
			return true
		}

		if current.value == nil {
			stack = append(stack, current.right)
			stack = append(stack, current.left)
		}
	}

	return false
}

func (sf *SnailFish) Magnitude() int {
	if sf.value != nil {
		return *sf.value
	}

	return 3*sf.left.Magnitude() + 2*sf.right.Magnitude()
}

func (sf *SnailFish) String() string {
	if sf.value != nil {
		return fmt.Sprint(*sf.value)
	}
	left, right := sf.left.String(), sf.right.String()
	return fmt.Sprintf("[%v,%v]", left, right)
}

func parse(str string, n int) (*SnailFish, int) {
	if str[n] == '[' {
		left, n := parse(str, n+1)
		right, n := parse(str, n+1)
		return &SnailFish{nil, left, right}, n + 1
	}

	// We always parse single digit numbers
	value := sh.ToInt(str[n : n+1])
	return &SnailFish{value: &value}, n + 1
}
