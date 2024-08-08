package main

import (
	"fmt"
	"strings"
)

type Matrix4x4 [][]float64
type Matrix2x2 [][]float64
type Matrix3x3 [][]float64

func NewMatrix4x4(elements [16]float64) Matrix4x4 {

	m44 := [][]float64{{0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}

	for y := int32(0); y < 4; y++ {
		for x := int32(0); x < 4; x++ {
			m44[y][x] = elements[(y*4)+x]
		}
	}

	return m44
}

func IdentitiyMatrix4x4() Matrix4x4 {
	return NewMatrix4x4([16]float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1})
}

func IdentitiyMatrix3x3() Matrix3x3 {
	return NewMatrix3x3([9]float64{1, 0, 0, 0, 0, 1, 0, 0, 0})
}

func (m44 Matrix4x4) ScalarMultiply(scalar float64) Matrix4x4 {

	mat44 := NewMatrix4x4([16]float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})

	for y := int32(0); y < 4; y++ {
		for x := int32(0); x < 4; x++ {
			mat44[x][y] = m44[x][y] * scalar
		}
	}

	return mat44
}

func (m33 Matrix3x3) ScalarMultiply(scalar float64) Matrix3x3 {

	mat33 := IdentitiyMatrix3x3()

	for y := int32(0); y < 3; y++ {
		for x := int32(0); x < 3; x++ {
			mat33[y][x] = m33[y][x] * scalar
		}
	}

	return mat33
}

func NewMatrix3x3(elements [9]float64) Matrix3x3 {

	m33 := [][]float64{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}

	for y := int32(0); y < 3; y++ {
		for x := int32(0); x < 3; x++ {
			m33[y][x] = elements[(y*3)+x]
		}
	}

	return m33
}

func NewMatrix2x2(elements [4]float64) Matrix2x2 {

	m22 := [][]float64{{0, 0}, {0, 0}}

	for y := int32(0); y < 2; y++ {
		m22[y] = make([]float64, 2)
		for x := int32(0); x < 2; x++ {
			m22[y][x] = elements[(y*2)+x]
		}
	}

	return m22
}

func (m44 Matrix4x4) Get(x, y int) float64 {
	return m44[x][y]
}

func (m33 Matrix3x3) Get(x, y int) float64 {

	return m33[x][y]
}

func (m22 Matrix2x2) Get(x, y int) float64 {
	return m22[x][y]
}

func (m44 Matrix4x4) Equal(other Matrix4x4) bool {

	if m44 == nil && other == nil {
		return true
	} else if m44 == nil || other == nil {
		return false
	}

	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			if !AreFloatsEqual(m44[x][y], other[x][y]) {
				return false
			}
		}
	}
	return true
}

func (m33 Matrix3x3) Equal(other Matrix3x3) bool {

	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			if !AreFloatsEqual(m33[x][y], other[x][y]) {
				return false
			}
		}
	}
	return true
}

func (m22 Matrix2x2) Equal(other Matrix2x2) bool {
	for y := 0; y < 2; y++ {
		for x := 0; x < 2; x++ {
			if !AreFloatsEqual(m22[x][y], other[x][y]) {
				return false
			}
		}
	}
	return true
}

func (m44 Matrix4x4) Multiply(other Matrix4x4) Matrix4x4 {
	matrix := [][]float64{{0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}

	for y := int32(0); y < 4; y++ {
		for x := int32(0); x < 4; x++ {
			matrix[y][x] = m44[y][0]*other[0][x] + m44[y][1]*other[1][x] + m44[y][2]*other[2][x] + m44[y][3]*other[3][x]
		}
	}

	return matrix
}

func (m44 Matrix4x4) TupleMultiply(tuple Tuple) Tuple {

	return Tuple{
		m44[0][0]*tuple.x + m44[0][1]*tuple.y + m44[0][2]*tuple.z + m44[0][3]*tuple.w,
		m44[1][0]*tuple.x + m44[1][1]*tuple.y + m44[1][2]*tuple.z + m44[1][3]*tuple.w,
		m44[2][0]*tuple.x + m44[2][1]*tuple.y + m44[2][2]*tuple.z + m44[2][3]*tuple.w,
		m44[3][0]*tuple.x + m44[3][1]*tuple.y + m44[3][2]*tuple.z + m44[3][3]*tuple.w,
	}
}

func (m33 Matrix3x3) Multiply(other Matrix3x3) Matrix3x3 {
	matrix := [][]float64{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}

	for y := int32(0); y < 3; y++ {
		for x := int32(0); x < 3; x++ {
			matrix[y][x] = m33[y][0]*other[0][x] + m33[y][1]*other[1][x] + m33[y][2]*other[2][x]
		}
	}

	return matrix
}

func (m22 Matrix2x2) Multiply(other Matrix2x2) Matrix2x2 {
	matrix := [][]float64{{0, 0}, {0, 0}}

	for y := int32(0); y < 2; y++ {
		for x := int32(0); x < 2; x++ {
			matrix[y][x] = m22[y][0]*other[0][x] + m22[y][1]*other[1][x]
		}
	}

	return matrix
}

func (m44 Matrix4x4) Transpose() Matrix4x4 {
	tempm44 := [][]float64{{0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}

	for y := int32(0); y < 4; y++ {
		for x := int32(0); x < 4; x++ {
			tempm44[x][y] = m44[y][x]
		}
	}

	return tempm44
}

func (m33 Matrix3x3) Transpose() Matrix3x3 {
	tempm33 := [][]float64{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}

	for y := int32(0); y < 3; y++ {
		for x := int32(0); x < 3; x++ {
			tempm33[x][y] = m33[y][x]
		}
	}

	return tempm33
}

func (m22 Matrix2x2) Transpose() Matrix2x2 {
	tempm22 := [][]float64{{0, 0}, {0, 0}}

	for y := int32(0); y < 3; y++ {
		for x := int32(0); x < 3; x++ {
			tempm22[x][y] = m22[y][x]
		}
	}

	return tempm22
}

func (m44 Matrix4x4) Print() string {

	var sb strings.Builder

	for y := int32(0); y < 4; y++ {
		if y > 0 {

			sb.WriteString("\n")
		}
		for x := int32(0); x < 4; x++ {
			sb.WriteString(fmt.Sprintf("%f ", m44[y][x]))
		}
	}

	return sb.String()
}

func (m33 Matrix3x3) Print() string {
	var sb strings.Builder

	for y := int32(0); y < 3; y++ {
		if y > 0 {

			sb.WriteString("\n")
		}
		for x := int32(0); x < 3; x++ {
			sb.WriteString(fmt.Sprintf("%f ", m33[y][x]))
		}
	}

	return sb.String()
}

func (m22 Matrix2x2) Print() string {
	var sb strings.Builder

	for y := int32(0); y < 2; y++ {
		if y > 0 {

			sb.WriteString("\n")
		}
		for x := int32(0); x < 2; x++ {
			sb.WriteString(fmt.Sprintf("%f ", m22[y][x]))
		}
	}

	return sb.String()
}

func (m22 Matrix2x2) Determinate() float64 {
	return m22[0][0]*m22[1][1] - m22[0][1]*m22[1][0]
}

func (m44 Matrix4x4) Submatrix(row, column int32) Matrix3x3 {

	tempm33 := [][]float64{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}

	newrow := 0

	for y := int32(0); y < 4; y++ {
		if y == row {
			continue
		} else {

			newColumn := 0

			for x := int32(0); x < 4; x++ {
				if x == column {
					continue
				} else {
					tempm33[newrow][newColumn] = m44[y][x]
					newColumn++
				}
			}
		}

		newrow++
	}

	return tempm33
}

func (m33 Matrix3x3) Submatrix(row, column int32) Matrix2x2 {

	tempmat2x2 := [][]float64{{0, 0}, {0, 0}}

	newrow := 0

	for y := int32(0); y < 3; y++ {
		if y == row {
			continue
		} else {

			newColumn := 0

			for x := int32(0); x < 3; x++ {
				if x == column {
					continue
				} else {
					tempmat2x2[newrow][newColumn] = m33[y][x]
					newColumn++
				}
			}
		}

		newrow++
	}

	return tempmat2x2
}

/*
Find the determinate of the sub matrix
*/
func (m33 Matrix3x3) Minor(row, column int32) float64 {

	m22 := m33.Submatrix(row, column)
	return m22.Determinate()
}

func (m44 Matrix4x4) Minor(row, column int32) float64 {

	m33 := m44.Submatrix(row, column)
	return m33.Determinate()
}

func (m33 Matrix3x3) Cofactor(row, column int32) float64 {
	if (row+column)%2 == 0 {
		return m33.Minor(row, column)
	} else {
		return -1 * m33.Minor(row, column)
	}
}

func (m44 Matrix4x4) Cofactor(row, column int32) float64 {
	if (row+column)%2 == 0 {
		return m44.Minor(row, column)
	} else {
		return -1 * m44.Minor(row, column)
	}
}

func (m33 Matrix3x3) Determinate() float64 {
	return m33.Cofactor(0, 0)*m33.Get(0, 0) + m33.Cofactor(0, 1)*m33.Get(0, 1) + m33.Cofactor(0, 2)*m33.Get(0, 2)
}

func (m44 Matrix4x4) Determinate() float64 {
	return m44.Cofactor(0, 0)*m44.Get(0, 0) + m44.Cofactor(0, 1)*m44.Get(0, 1) + m44.Cofactor(0, 2)*m44.Get(0, 2) + m44.Cofactor(0, 3)*m44.Get(0, 3)
}

func (m44 Matrix4x4) IsInvertible() bool {
	return m44.Determinate() != 0
}

func (m33 Matrix3x3) IsInvertible() bool {
	return m33.Determinate() != 0
}

func (m22 Matrix2x2) IsInvertible() bool {
	return m22.Determinate() != 0
}

func (m44 Matrix4x4) cofactorMatrix() Matrix4x4 {
	cofactorMatrix := NewMatrix4x4([16]float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})

	for row := 0; row < 4; row++ {
		for column := 0; column < 4; column++ {

			// cofactorMatrix[row][column] = m44.Cofactor(int32(row), int32(column)
			cofactorMatrix[row][column] = m44.Cofactor(int32(row), int32(column))
		}
	}

	return cofactorMatrix
}

func (m33 Matrix3x3) cofactorMatrix() Matrix3x3 {
	cofactorMatrix := NewMatrix3x3([9]float64{0, 0, 0, 0, 0, 0, 0, 0, 0})

	for row := 0; row < 3; row++ {
		for column := 0; column < 3; column++ {

			// cofactorMatrix[row][column] = m44.Cofactor(int32(row), int32(column)
			cofactorMatrix[row][column] = m33.Cofactor(int32(row), int32(column))
		}
	}

	return cofactorMatrix
}

func (m44 Matrix4x4) Inverse() Matrix4x4 {
	determinate := m44.Determinate()

	if determinate == 0 {
		return nil
	}

	factor := 1 / determinate

	if m44.IsInvertible() {

		transposeCofactor := m44.cofactorMatrix().Transpose()

		return transposeCofactor.ScalarMultiply(factor)
	} else {
		return nil
	}

}

func (m33 Matrix3x3) Inverse() Matrix3x3 {
	determinate := m33.Determinate()

	if determinate == 0 {
		return nil
	}

	factor := 1 / determinate
	transposeCofactor := m33.cofactorMatrix().Transpose()

	return transposeCofactor.ScalarMultiply(factor)

}
