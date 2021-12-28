package shared

import (
	"strconv"
	"strings"
)

func Lines(str string) []string {
	return strings.Split(strings.TrimRight(str, "\n"), "\n")
}

func IntList(str string) []int {
	lines := Lines(str)
	return ToInts(strings.Split(lines[0], ","))
}

func ToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}

func ToInts(sli []string) []int {
	ints := make([]int, len(sli))
	for i, v := range sli {
		ints[i] = ToInt(v)
	}
	return ints
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
