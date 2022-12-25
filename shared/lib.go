package shared

import (
	"sort"
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

func BinToInt(binary string) int {
	i, _ := strconv.ParseInt(binary, 2, 64)
	return int(i)
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

func Max(vars ...int) int {
	max := vars[0]
	for _, v := range vars {
		if max < v {
			max = v
		}
	}
	return max
}

func Min(vars ...int) int {
	min := vars[0]
	for _, v := range vars {
		if min > v {
			min = v
		}
	}
	return min
}

func Mean(vars []int) int {
	sum := 0
	for _, v := range vars {
		sum += v
	}
	return sum / len(vars)
}

func Median(vars []int) int {
	l := len(vars)
	sort.Ints(vars)
	if l%2 == 0 {
		return (vars[l/2-1] + vars[l/2]) / 2
	}
	return (vars[l/2])
}

func TriangularNumber(n int) int {
	return n * (n + 1) / 2
}

func IntDict(ints []int) map[int]int {
	m := make(map[int]int)
	for _, v := range ints {
		m[v]++
	}
	return m
}

func ReverseStr(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// integer b^n
func Pow(b, n int) int {
	r := 1
	for i := 0; i < n; i++ {
		r *= b
	}
	return r
}
