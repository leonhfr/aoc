package matrix

import (
	"fmt"
	"strings"

	sh "github.com/leonhfr/aoc/shared"
)

type Matrix [][]int

type Coordinates struct {
	I, J int
}

func NewMatrix(m, n int) Matrix {
	matrix := make(Matrix, m)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	return matrix
}

func IntMatrix(str string) Matrix {
	lines := sh.Lines(str)
	matrix := NewMatrix(len(lines), len(lines[0]))
	for y, line := range lines {
		matrix[y] = sh.ToInts(strings.Split(line, ""))
	}
	return matrix
}

func (m Matrix) Duplicate() Matrix {
	matrix := NewMatrix(m.M(), m.N())
	for y := range m {
		row := make([]int, m.N())
		copy(row, m[y])
		matrix[y] = row
	}
	return matrix
}

func (m Matrix) M() int {
	return len(m)
}

func (m Matrix) N() int {
	return len(m[0])
}

func (m Matrix) Inside(i, j int) bool {
	return 1 <= i && i <= m.M() && 1 <= j && j <= m.N()
}

func (m Matrix) Get(i, j int) int {
	return m[i-1][j-1]
}

func (m Matrix) Set(i, j, v int) {
	m[i-1][j-1] = v
}

func (m Matrix) Inc(i, j, v int) {
	m[i-1][j-1] += v
}

func (m Matrix) Coordinates() []Coordinates {
	var coords []Coordinates
	for i := 0; i < m.M(); i++ {
		for j := 0; j < m.N(); j++ {
			coords = append(coords, Coordinates{i + 1, j + 1})
		}
	}
	return coords
}

func (m *Matrix) String() string {
	tmp := *m
	matrix := make([]string, m.M())
	for i := 0; i < m.M(); i++ {
		row := make([]string, m.N())
		for j := 0; j < m.N(); j++ {
			if 0 < tmp[i][j] && tmp[i][j] <= 9 {
				row[j] = fmt.Sprint(tmp[i][j])
			} else {
				row[j] = "."
			}
		}
		matrix[i] = strings.Join(row, "")
	}
	return strings.Join(matrix, "\n")
}
