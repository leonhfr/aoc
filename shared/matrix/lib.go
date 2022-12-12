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

type Direction int

const (
	Clockwise Direction = iota
	CounterClockwise
)

func NewMatrix(m, n int) Matrix {
	matrix := make(Matrix, m)
	for i := 0; i < m; i++ {
		matrix[i] = make([]int, n)
	}
	return matrix
}

func NewMatrixWithDefault(m, n, v int) Matrix {
	matrix := make(Matrix, m)
	for i := 0; i < m; i++ {
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			matrix[i][j] = v
		}
	}
	return matrix
}

func IntMatrix(str string) Matrix {
	lines := sh.Lines(str)
	matrix := NewMatrix(len(lines), len(lines[0]))
	for i, line := range lines {
		matrix[i] = sh.ToInts(strings.Split(line, ""))
	}
	return matrix
}

func SharpMatrix(str string) Matrix {
	lines := sh.Lines(str)
	matrix := NewMatrix(len(lines), len(lines[0]))
	for i, line := range lines {
		for j, r := range line {
			if r == '#' {
				matrix.Set(i+1, j+1, 1)
			} else {
				matrix.Set(i+1, j+1, 0)
			}
		}
	}
	return matrix
}

func (m Matrix) Duplicate() Matrix {
	matrix := NewMatrix(m.M(), m.N())
	for i := range m {
		row := make([]int, m.N())
		copy(row, m[i])
		matrix[i] = row
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

func (m Matrix) Row(i int) []int {
	row := make([]int, m.N())
	copy(row, m[i-1])
	return row
}

func (m Matrix) Col(j int) []int {
	col := make([]int, m.M())
	for i, row := range m {
		col[i] = row[j-1]
	}
	return col
}

func (m Matrix) Swap(i1, j1, i2, j2 int) {
	m[i1-1][j1-1], m[i2-1][j2-1] = m[i2-1][j2-1], m[i1-1][j1-1]
}

func (m Matrix) Transpose() {
	if m.M() != m.N() {
		panic("Not a square matrix.")
	}

	for i := 1; i <= m.M(); i++ {
		for j := i; j <= m.N(); j++ {
			m.Swap(i, j, j, i)
		}
	}
}

func (m Matrix) Rotate(d Direction) {
	m.Transpose()
	switch d {
	case Clockwise:
		m.VerticalFlip()
	case CounterClockwise:
		m.HorizontalFlip()
	}
}

func (m Matrix) HorizontalFlip() {
	for i := 1; i <= m.M()/2; i++ {
		for j := 1; j <= m.N(); j++ {
			m.Swap(i, j, m.M()+1-i, j)
		}
	}
}

func (m Matrix) VerticalFlip() {
	for i := 1; i <= m.M(); i++ {
		for j := 1; j <= m.N()/2; j++ {
			m.Swap(i, j, i, m.N()+1-j)
		}
	}
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
