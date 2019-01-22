package matrix

type Matrix interface {
	Rows() int
	Cols() int
	At(i, j int) int
	Set(i, j, v int)
}

//
// simpleMatrix
//

type simpleMatrix [][]int

func newSimpleMatrix(r, c int) Matrix {
	mat := make([][]int, r)
	for i := 0; i < r; i++ {
		mat[i] = make([]int, c)
	}
	return simpleMatrix(mat)
}

func (m simpleMatrix) Rows() int       { return len(m) }
func (m simpleMatrix) Cols() int       { return len(m[0]) }
func (m simpleMatrix) At(i, j int) int { return m[i][j] }
func (m simpleMatrix) Set(i, j, v int) { m[i][j] = v }

//
// denseMatrix
//

type denseMatrix [][]int

func newDenseMatrix(r, c int) Matrix {
	mem := make([]int, r*c)
	mat := make([][]int, r)
	for i := 0; i < r; i++ {
		mat[i], mem = mem[:c], mem[c:]
	}
	return denseMatrix(mat)
}

func (m denseMatrix) Rows() int       { return len(m) }
func (m denseMatrix) Cols() int       { return len(m[0]) }
func (m denseMatrix) At(i, j int) int { return m[i][j] }
func (m denseMatrix) Set(i, j, v int) { m[i][j] = v }

//
// flatMatrix
//

type flatMatrix struct {
	data []int
	rows int
	cols int
}

func newFlatMatrix(r, c int) Matrix {
	return &flatMatrix{
		data: make([]int, r*c),
		rows: r,
		cols: c,
	}
}

func (m *flatMatrix) Rows() int       { return m.rows }
func (m *flatMatrix) Cols() int       { return m.cols }
func (m *flatMatrix) At(i, j int) int { return m.data[i*m.cols+j] }
func (m *flatMatrix) Set(i, j, v int) { m.data[i*m.cols+j] = v }
