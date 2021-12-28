package shared

import "strings"

type Matrix [][]int

func NewMatrix(m, n int) Matrix {
	matrix := make(Matrix, m)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	return matrix
}

func IntMatrix(str string) Matrix {
	lines := Lines(str)
	matrix := NewMatrix(len(lines), len(lines[0]))
	for y, line := range lines {
		matrix[y] = ToInts(strings.Split(line, ""))
	}
	return matrix
}

func (m Matrix) M() int {
	return len(m)
}

func (m Matrix) N() int {
	return len(m[0])
}

func (m Matrix) Get(i, j int) int {
	return m[i][j]
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

type Coordinates struct {
	I, J int
}
