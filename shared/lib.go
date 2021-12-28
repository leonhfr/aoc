package shared

import (
	"strconv"
	"strings"
)

func Lines(str string) []string {
	return strings.Split(strings.TrimRight(str, "\n"), "\n")
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
