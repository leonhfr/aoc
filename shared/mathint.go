package shared

type Matrix [][]int

func NewMatrix(m, n int) Matrix {
	matrix := make(Matrix, m)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
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

type Coordinates struct {
	I, J int
}
