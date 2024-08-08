package main

import (
	"testing"
)

func TestNewMartix4x4(t *testing.T) {

	type check struct {
		x     int
		y     int
		value float64
	}

	tests := []struct {
		name     string
		elements [16]float64
		checks   []check
	}{
		{
			name:     "Create a 4x4 Matrix",
			elements: [16]float64{1, 2, 3, 4, 5.5, 6.5, 7.5, 8.5, 9, 10, 11, 12, 13.5, 14.5, 15.5, 16.5},
			checks:   []check{{0, 0, 1}, {1, 0, 5.5}, {1, 2, 7.5}, {2, 2, 11}, {3, 0, 13.5}, {3, 2, 15.5}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewMatrix4x4(tt.elements)

			for _, c := range tt.checks {
				if !AreFloatsEqual(got.Get(c.x, c.y), c.value) {
					t.Errorf("Matrix4x4 at (%d, %d)returned %v, wanted %v", c.x, c.y, got.Get(c.x, c.y), c.value)
				}
			}

		})
	}
}

func TestMartix4x4Equal(t *testing.T) {

	tests := []struct {
		name    string
		matrix1 Matrix4x4
		matrix2 Matrix4x4
		want    bool
	}{
		{
			name:    "Matrix4x4 Should be equal",
			matrix1: NewMatrix4x4([16]float64{1, 2, 3, 4, 5.5, 6.5, 7.5, 8.5, 9, 10, 11, 12, 13.5, 14.5, 15.5, 16.5}),
			matrix2: NewMatrix4x4([16]float64{1, 2, 3, 4, 5.5, 6.5, 7.5, 8.5, 9, 10, 11, 12, 13.5, 14.5, 15.5, 16.5}),
			want:    true,
		},
		{
			name:    "Matrix4x4 Should not be equal",
			matrix1: NewMatrix4x4([16]float64{1, 2, 3, 4, 5.5, 6.5, 7.5, 8.6, 9, 10, 11, 12, 13.5, 14.5, 15.5, 16.5}),
			matrix2: NewMatrix4x4([16]float64{1, 2, 3, 4, 5.5, 6.5, 7.5, 8.5, 9, 10, 11, 12, 13.5, 14.5, 15.5, 16.5}),
			want:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.matrix1.Equal(tt.matrix2)
			if got != tt.want {
				t.Errorf("Matrix4x4 returned %v, wanted %v", got, tt.want)
			}

		})
	}
}

func TestMartix3x3(t *testing.T) {

	type check struct {
		x     int
		y     int
		value float64
	}

	tests := []struct {
		name     string
		elements [9]float64
		checks   []check
	}{
		{
			name:     "Create a 3x3 Matrix",
			elements: [9]float64{-3, 5, 0, 1, -2, -7, 0, 1, 1},
			checks:   []check{{0, 0, -3}, {1, 1, -2}, {2, 2, 1}, {2, 0, 0}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewMatrix3x3(tt.elements)

			for _, c := range tt.checks {
				if !AreFloatsEqual(got.Get(c.x, c.y), c.value) {
					t.Errorf("Matrix3x3 at (%d, %d)returned %v, wanted %v", c.x, c.y, got.Get(c.x, c.y), c.value)
				}
			}

		})
	}
}

func TestMartix3x3Equal(t *testing.T) {

	tests := []struct {
		name    string
		matrix1 Matrix3x3
		matrix2 Matrix3x3
		want    bool
	}{
		{
			name:    "Matrix3x3 Should be equal",
			matrix1: NewMatrix3x3([9]float64{1, 2, 3, 4, 5.5, 6.5, 7.5, 8.5, 9}),
			matrix2: NewMatrix3x3([9]float64{1, 2, 3, 4, 5.5, 6.5, 7.5, 8.5, 9}),
			want:    true,
		},
		{
			name:    "Matrix3x3 Should not be equal",
			matrix1: NewMatrix3x3([9]float64{1, 2, 3, 4, 5.5, 6.5, 7.5, 8.6, 9}),
			matrix2: NewMatrix3x3([9]float64{1, 2, 3, 4, 5.5, 6.7, 7.5, 8.5, 10}),
			want:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.matrix1.Equal(tt.matrix2)
			if got != tt.want {
				t.Errorf("Matrix3x3 returned %v, wanted %v", got, tt.want)
			}

		})
	}
}

func TestMartix2x2(t *testing.T) {

	type check struct {
		x     int
		y     int
		value float64
	}

	tests := []struct {
		name     string
		elements [4]float64
		checks   []check
	}{
		{
			name:     "Create a 2x2 Matrix",
			elements: [4]float64{-3, 5, 1, -2},
			checks:   []check{{0, 0, -3}, {1, 0, 1}, {1, 1, -2}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewMatrix2x2(tt.elements)

			for _, c := range tt.checks {
				if !AreFloatsEqual(got.Get(c.x, c.y), c.value) {
					t.Errorf("Matrix2x2 at (%d, %d)returned %v, wanted %v", c.x, c.y, got.Get(c.x, c.y), c.value)
				}
			}

		})
	}
}

func TestMartix2x2Equal(t *testing.T) {

	tests := []struct {
		name    string
		matrix1 Matrix2x2
		matrix2 Matrix2x2
		want    bool
	}{
		{
			name:    "Matrix2x2 Should be equal",
			matrix1: NewMatrix2x2([4]float64{1, 2, 3, 4}),
			matrix2: NewMatrix2x2([4]float64{1, 2, 3, 4}),
			want:    true,
		},
		{
			name:    "Matrix2x2 Should not be equal",
			matrix1: NewMatrix2x2([4]float64{1, 2, 3, 4}),
			matrix2: NewMatrix2x2([4]float64{1, 2.2, 3, 4}),
			want:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.matrix1.Equal(tt.matrix2)
			if got != tt.want {
				t.Errorf("Matrix3x3 returned %v, wanted %v", got, tt.want)
			}

		})
	}
}

func TestMartix4x4Multiply(t *testing.T) {

	tests := []struct {
		name    string
		matrix1 Matrix4x4
		matrix2 Matrix4x4
		result  Matrix4x4
		want    bool
	}{
		{
			name:    "Matrix4x4 Should be equal",
			matrix1: NewMatrix4x4([16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2}),
			matrix2: NewMatrix4x4([16]float64{-2, 1, 2, 3, 3, 2, 1, -1, 4, 3, 6, 5, 1, 2, 7, 8}),
			result:  NewMatrix4x4([16]float64{20, 22, 50, 48, 44, 54, 114, 108, 40, 58, 110, 102, 16, 26, 46, 42}),
			want:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.matrix1.Multiply(tt.matrix2)
			if got.Equal(tt.result) != tt.want {
				t.Errorf("Matrix4x4 returned %v, wanted %v", got, tt.want)
			}

		})
	}
}

func TestMartix4x4TupleMultiply(t *testing.T) {

	tests := []struct {
		name   string
		matrix Matrix4x4
		tuple  Tuple
		want   Tuple
	}{
		{
			name:   "Matrix4x4 multiplied by a Vector4 should return a Vector4",
			matrix: NewMatrix4x4([16]float64{1, 2, 3, 4, 2, 4, 4, 2, 8, 6, 4, 1, 0, 0, 0, 1}),
			tuple:  Tuple{1, 2, 3, 1},
			want:   Tuple{18, 24, 33, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.matrix.TupleMultiply(tt.tuple)
			equal := got.Equal(tt.want)
			if !equal {
				t.Errorf("Matrix4x4 Vector Multiplication returned %v, wanted %v", false, true)
			}

		})
	}
}

func TestTransposeMatrix4x4(t *testing.T) {

	tests := []struct {
		name       string
		matrix     Matrix4x4
		transposed Matrix4x4
	}{
		{
			name:       "Matrix4x4 multiplied by a Vector4 should return a Vector4",
			matrix:     NewMatrix4x4([16]float64{0, 9, 3, 0, 9, 8, 0, 8, 1, 8, 5, 3, 0, 0, 5, 8}),
			transposed: NewMatrix4x4([16]float64{0, 9, 1, 0, 9, 8, 8, 0, 3, 0, 5, 5, 0, 8, 3, 8}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.matrix.Transpose()

			if !got.Equal(tt.transposed) {
				t.Errorf("Matrix4x4 Vector Transpose test failed")
			}

		})
	}

}

func TestMatrix2x2Determinate(t *testing.T) {
	tests := []struct {
		name   string
		matrix Matrix2x2
		want   float64
	}{
		{
			name:   "Matrix2x2 Determinate",
			matrix: NewMatrix2x2([4]float64{1, 5, -3, 2}),
			want:   17,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.matrix.Determinate()

			if !AreFloatsEqual(got, tt.want) {
				t.Errorf("Matrix2x2 Determinate test failed, wanted %f but got %f", tt.want, got)
			}

		})
	}
}

func TestMartix4x4Submatrix(t *testing.T) {
	tests := []struct {
		name   string
		args   [2]int32
		matrix Matrix4x4
		want   Matrix3x3
	}{
		{
			name:   "Matrix4x4 Submatrix 0,0",
			args:   [2]int32{0, 0},
			matrix: NewMatrix4x4([16]float64{0, 9, 3, 0, 9, 8, 0, 8, 1, 8, 5, 3, 0, 0, 5, 8}),
			want:   NewMatrix3x3([9]float64{8, 0, 8, 8, 5, 3, 0, 5, 8}),
		},
		{
			name:   "Matrix4x4 Submatrix 1,0",
			args:   [2]int32{1, 0},
			matrix: NewMatrix4x4([16]float64{0, 9, 3, 0, 9, 8, 0, 8, 1, 8, 5, 3, 0, 0, 5, 8}),
			want:   NewMatrix3x3([9]float64{9, 3, 0, 8, 5, 3, 0, 5, 8}),
		},
		{
			name:   "Matrix4x4 Submatrix 3,3",
			args:   [2]int32{3, 3},
			matrix: NewMatrix4x4([16]float64{0, 9, 3, 0, 9, 8, 0, 8, 1, 8, 5, 3, 0, 0, 5, 8}),
			want:   NewMatrix3x3([9]float64{0, 9, 3, 9, 8, 0, 1, 8, 5}),
		},
		{
			name:   "Matrix4x4 Submatrix 2,2",
			args:   [2]int32{2, 2},
			matrix: NewMatrix4x4([16]float64{0, 9, 3, 0, 9, 8, 0, 8, 1, 8, 5, 3, 0, 0, 5, 8}),
			want:   NewMatrix3x3([9]float64{0, 9, 0, 9, 8, 8, 0, 0, 8}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.matrix.Submatrix(tt.args[0], tt.args[1])

			if !got.Equal(tt.want) {
				t.Errorf("Matrix2x2 Determinate test failed, \n original \n%s \nwanted \n%s \nbut got \n%s\n", tt.matrix.Print(), tt.want.Print(), got.Print())
			}

		})
	}
}

func TestMartix3x3Submatrix(t *testing.T) {
	tests := []struct {
		name   string
		args   [2]int32
		matrix Matrix3x3
		want   Matrix2x2
	}{
		{
			name:   "Matrix4x4 Submatrix 0,0",
			args:   [2]int32{0, 0},
			matrix: NewMatrix3x3([9]float64{0, 9, 3, 0, 9, 8, 0, 8, 1}),
			want:   NewMatrix2x2([4]float64{9, 8, 8, 1}),
		},
		{
			name:   "Matrix4x4 Submatrix 1,0",
			args:   [2]int32{1, 0},
			matrix: NewMatrix3x3([9]float64{0, 9, 3, 0, 9, 8, 0, 8, 1}),
			want:   NewMatrix2x2([4]float64{9, 3, 8, 1}),
		},
		{
			name:   "Matrix4x4 Submatrix 2,2",
			args:   [2]int32{2, 2},
			matrix: NewMatrix3x3([9]float64{0, 9, 3, 0, 9, 8, 0, 8, 1}),
			want:   NewMatrix2x2([4]float64{0, 9, 0, 9}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.matrix.Submatrix(tt.args[0], tt.args[1])

			if !got.Equal(tt.want) {
				t.Errorf("Matrix2x2 Determinate test failed, \n original \n%s \nwanted \n%s \nbut got \n%s\n", tt.matrix.Print(), tt.want.Print(), got.Print())
			}

		})
	}
}

func TestCofactor3x3(t *testing.T) {
	tests := []struct {
		name   string
		args   [2]int32
		matrix Matrix3x3
		want   float64
	}{
		{
			name:   "cofactor of a 3x3 matrix",
			args:   [2]int32{0, 0},
			matrix: NewMatrix3x3([9]float64{3, 5, 0, 2, -1, -7, 6, -1, 5}),
			want:   -12,
		},
		{
			name:   "cofactor of a 3x3 matrix",
			args:   [2]int32{0, 1},
			matrix: NewMatrix3x3([9]float64{3, 5, 0, 2, -1, -7, 6, -1, 5}),
			want:   -52,
		},
		{
			name:   "cofactor of a 3x3 matrix",
			args:   [2]int32{0, 0},
			matrix: NewMatrix3x3([9]float64{8, 51, 10, 1, .5, 7, .25, 4, 10}),
			want:   -23,
		},
		{
			name:   "cofactor of a 3x3 matrix",
			args:   [2]int32{1, 2},
			matrix: NewMatrix3x3([9]float64{8, 51, 10, 1, .5, 7, .25, 4, 10}),
			want:   -19.25,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.matrix.Cofactor(tt.args[0], tt.args[1])

			if !AreFloatsEqual(got, tt.want) {
				t.Errorf("%s test failed, \n original \n%s \nwanted \n%f \nbut got \n%f\n", tt.name, tt.matrix.Print(), tt.want, got)
			}

		})
	}
}

func TestCofactor4x4(t *testing.T) {
	tests := []struct {
		name   string
		args   [2]int32
		matrix Matrix4x4
		want   float64
	}{
		{
			name:   "cofactor of a 4x4 matrix",
			args:   [2]int32{0, 0},
			matrix: NewMatrix4x4([16]float64{10, 1, 20, 2, 0.10, 10, 0.20, 20, 8, 51, 10, 1, 0.5, 7, 0.25, 4}),
			want:   -786.9,
		},
		{
			name:   "cofactor of a 4x4 matrix",
			args:   [2]int32{0, 1},
			matrix: NewMatrix4x4([16]float64{10, 1, 20, 2, 0.10, 10, 0.20, 20, 8, 51, 10, 1, 0.5, 7, 0.25, 4}),
			want:   62.325,
		},
		{
			name:   "cofactor of a 4x4 matrix",
			args:   [2]int32{0, 0},
			matrix: NewMatrix4x4([16]float64{7, -7, 8, -8, 9, -9, 10, -10, 11, -11, 12, -12, 13, -13, 14, -14}),
			want:   0,
		},
		{
			name:   "cofactor of a 4x4 matrix",
			args:   [2]int32{3, 3},
			matrix: NewMatrix4x4([16]float64{46, 75, 20, 32, 87, 60, 48, 11, 51, 82, 54, 55, 16, 12, 43, 64}),
			want:   -119286,
		},
		{
			name:   "cofactor of a 4x4 matrix",
			args:   [2]int32{3, 3},
			matrix: NewMatrix4x4([16]float64{10, 1, 10, 1, 0.10, 10, 0.2, 20, 8, 51, 10, 1, 0.5, 7, 0.25, 4}),
			want:   149.6,
		},
		{
			name:   "cofactor of a 4x4 matrix",
			args:   [2]int32{3, 1},
			matrix: NewMatrix4x4([16]float64{10, 1, 10, 1, 0.10, 10, 0.2, 20, 8, 51, 10, 1, 0.5, 7, 0.25, 4}),
			want:   -399.6,
		},
		{
			name:   "cofactor of a 4x4 matrix",
			args:   [2]int32{0, 0},
			matrix: NewMatrix4x4([16]float64{10, 1, 10, 1, 0.10, 10, 0.2, 20, 8, 51, 10, 1, 0.5, 7, 0.25, 4}),
			want:   -786.9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.matrix.Cofactor(tt.args[0], tt.args[1])

			if !AreFloatsEqual(got, tt.want) {
				t.Errorf("%s test failed, \n original \n%s \nwanted \n%f \nbut got \n%f\n", tt.name, tt.matrix.Print(), tt.want, got)
			}

		})
	}
}

func TestCofactorMatrix(t *testing.T) {
	tests := []struct {
		name   string
		matrix Matrix4x4
		want   Matrix4x4
	}{
		{
			name:   "cofactor of a 4x4 matrix",
			matrix: NewMatrix4x4([16]float64{10, 1, 10, 1, 0.10, 10, 0.2, 20, 8, 51, 10, 1, 0.5, 7, 0.25, 4}),
			want:   NewMatrix4x4([16]float64{-786.9, 62.325, 314.7, -30.375, 1987.5, 79.5, -1969, -264.5, 996.9, -53.925, -994.7, 31.925, -9990, -399.6, 10015, 149.6}),
		},
		{
			name:   "cofactor of a 4x4 matrix",
			matrix: NewMatrix4x4([16]float64{7, -7, 8, -8, 9, -9, 10, -10, 11, -11, 12, -12, 13, -13, 14, -14}),
			want:   NewMatrix4x4([16]float64{}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.matrix.cofactorMatrix()

			if !got.Equal(tt.want) {
				t.Errorf("%s test failed, \n original \n%s \nwanted \n%s \nbut got \n%s\n", tt.name, tt.matrix.Print(), tt.want.Print(), got.Print())
			}

		})
	}
}

func TestMinor4x4(t *testing.T) {
	tests := []struct {
		name   string
		args   [2]int32
		matrix Matrix4x4
		want   float64
	}{
		{

			name:   "minor of a 3x3 matrix",
			args:   [2]int32{0, 0},
			matrix: NewMatrix4x4([16]float64{10, 1, 20, 2, 0.10, 10, 0.20, 20, 8, 51, 10, 1, 0.5, 7, 0.25, 4}),
			want:   -786.9,
		},
		{
			/*
				10, 1, 20, 2,
				0.10, 10, 0.20, 20,
				8, 51, 10, 1,
				0.5, 7, 0.25, 4
			*/
			name:   "minor of a 3x3 matrix",
			args:   [2]int32{0, 1},
			matrix: NewMatrix4x4([16]float64{10, 1, 20, 2, 0.10, 10, 0.20, 20, 8, 51, 10, 1, 0.5, 7, 0.25, 4}),
			want:   -62.325,
		},
		{
			/*
				8,	51,		10, 	1,
				.5,	7,		.25, 	4,
				10,	10,		1, 		20,
				2,	0.10,	10, 	0.20
			*/
			name:   "minor of a 3x3 matrix",
			args:   [2]int32{0, 0},
			matrix: NewMatrix4x4([16]float64{8, 51, 10, 1, .5, 7, .25, 4, 10, 10, 1, 20, 2, 0.10, 10, 0.20}),
			want:   -999.00,
		},
		{
			/*
				8,	51,		10, 	1,

				10,	10,		1, 		20,
				2,	0.10,	10, 	0.20
			*/
			name:   "minor of a 3x3 matrix",
			args:   [2]int32{1, 3},
			matrix: NewMatrix4x4([16]float64{8, 51, 10, 1, .5, 7, .25, 4, 10, 10, 1, 20, 2, 0.10, 10, 0.20}),
			want:   -4388.8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.matrix.Minor(tt.args[0], tt.args[1])

			if !AreFloatsEqual(got, tt.want) {
				t.Errorf("%s test failed, \n original \n%s \nwanted \n%f \nbut got \n%f\n", tt.name, tt.matrix.Print(), tt.want, got)
			}

		})
	}
}

func TestMinor3x3(t *testing.T) {
	tests := []struct {
		name   string
		args   [2]int32
		matrix Matrix3x3
		want   float64
	}{
		{

			name:   "minor of a 3x3 matrix",
			args:   [2]int32{1, 1},
			matrix: NewMatrix3x3([9]float64{3, 5, 0, 2, -1, -7, 6, -1, 5}),
			want:   15,
		},
		{

			name:   "minor of a 3x3 matrix",
			args:   [2]int32{0, 0},
			matrix: NewMatrix3x3([9]float64{3, 5, 0, 2, -1, -7, 6, -1, 5}),
			want:   -12,
		},
		{

			name:   "minor of a 3x3 matrix",
			args:   [2]int32{0, 0},
			matrix: NewMatrix3x3([9]float64{7, -7, 8, -8, 9, -9, 10, -10, 11}),
			want:   9,
		},
		{
			name:   "minor of a 3x3 matrix",
			args:   [2]int32{2, 2},
			matrix: NewMatrix3x3([9]float64{46, 75, 20, 32, 87, 60, 48, 11, 51}),
			want:   1602,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.matrix.Minor(tt.args[0], tt.args[1])

			if !AreFloatsEqual(got, tt.want) {
				t.Errorf("%s test failed, \n original \n%s \nwanted \n%f \nbut got \n%f\n", tt.name, tt.matrix.Print(), tt.want, got)
			}

		})
	}
}

func TestDeterminate4x4(t *testing.T) {
	tests := []struct {
		name   string
		matrix Matrix4x4
		want   float64
	}{
		{

			name:   "minor of a 4x44 matrix",
			matrix: NewMatrix4x4([16]float64{10, 1, 20, 2, 0.10, 10, 0.20, 20, 8, 51, 10, 1, 0.5, 7, 0.25, 4}),
			want:   -1573.425,
		},
		{

			name: "minor of a 4x44 matrix",

			matrix: NewMatrix4x4([16]float64{1, 1, 2, 2, 1, 1, 2, 2, 4, 5, 1, 1, 5, 7, 5, 4}),
			want:   0,
		},
		{

			name:   "minor of a 4x44 matrix",
			matrix: NewMatrix4x4([16]float64{35, 14, 93, 62, 85, 84, 56, 84, 84, 62, 96, 92, 49, 46, 22, 28}),
			want:   656256,
		},
		{
			name:   "minor of a 4x44 matrix",
			matrix: NewMatrix4x4([16]float64{50.620, 50.02, 95.730, 89.200, 45.260, 59.970, 39.790, 45.030, 62.870, 28.610, 48.710, 80.710, 40.830, 63.250, 63.280, 30.350}),
			want:   3112830.36700,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.matrix.Determinate()

			if !AreFloatsEqual(got, tt.want) {
				t.Errorf("%s test failed, \n original \n%s \nwanted \n%f \nbut got \n%f\n", tt.name, tt.matrix.Print(), tt.want, got)
			}

		})
	}
}

func TestDeterminate3x3(t *testing.T) {
	tests := []struct {
		name   string
		matrix Matrix3x3
		want   float64
	}{
		{

			name:   "Determinate of a 4x4 matrix",
			matrix: NewMatrix3x3([9]float64{10, 1, 20, 2, 0.10, 10, 0.20, 20, 8}),
			want:   -1206.4,
		},
		{

			name: "Determinate of a 4x4 matrix",

			matrix: NewMatrix3x3([9]float64{1, 1, 2, 2, 1, 1, 2, 2, 4}),
			want:   0,
		},
		{

			name:   "Determinate of a 4x4 matrix",
			matrix: NewMatrix3x3([9]float64{35, 14, 93, 62, 85, 84, 56, 84, 84}),
			want:   37547.99999999999998,
		},
		{
			name:   "Determinate of a 4x4 matrix",
			matrix: NewMatrix3x3([9]float64{50.620, 50.02, 95.730, 89.200, 45.260, 59.970, 39.790, 45.030, 62.870}),
			want:   58304.640086,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.matrix.Determinate()

			if !AreFloatsEqual(got, tt.want) {
				t.Errorf("%s test failed, \n original \n%s \nwanted \n%f \nbut got \n%f\n", tt.name, tt.matrix.Print(), tt.want, got)
			}

		})
	}
}

func TestCofactorTranspose(t *testing.T) {
	tests := []struct {
		name   string
		matrix Matrix4x4
		want   Matrix4x4
	}{
		{

			name:   "Inverse of a 4x44 matrix",
			matrix: NewMatrix4x4([16]float64{10, 1, 10, 1, 0.10, 10, 0.2, 20, 8, 51, 10, 1, 0.5, 7, 0.25, 4}),
			want:   NewMatrix4x4([16]float64{-786.9, 1987.5, 996.9, -9990, 62.325, 79.5, -53.925, -399.6, 314.7, -1969, -994.7, 10015, -30.375, -264.5, 31.925, 149.6}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.matrix.cofactorMatrix().Transpose()

			if !got.Equal(tt.want) {
				t.Errorf("%s test failed, \n original \n%s \nwanted \n%s \nbut got \n%s\n", tt.name, tt.matrix.Print(), tt.want.Print(), got.Print())
			}

		})
	}
}

func TestScalarMultiplyMatrix4x4(t *testing.T) {
	tests := []struct {
		name   string
		matrix Matrix4x4
		scalar float64
		want   Matrix4x4
	}{
		{

			name:   "Inverse of a 4x44 matrix",
			matrix: NewMatrix4x4([16]float64{10, 1, 10, 1, 0.10, 10, 0.2, 20, 8, 51, 10, 1, 0.5, 7, 0.25, 4}),
			scalar: 4,
			want:   NewMatrix4x4([16]float64{40, 4, 40, 4, 0.40, 40, 0.8, 80, 32, 204, 40, 4, 2, 28, 1, 16}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.matrix.ScalarMultiply(tt.scalar)

			if !got.Equal(tt.want) {
				t.Errorf("%s test failed, \n original \n%s \nwanted \n%s \nbut got \n%s\n", tt.name, tt.matrix.Print(), tt.want.Print(), got.Print())
			}

		})
	}
}

func TestInverse4x4(t *testing.T) {
	tests := []struct {
		name   string
		matrix Matrix4x4
		want   Matrix4x4
	}{
		{
			/*
				10, 1, 10, 1,
				0.10, 10, 0.2, 20,
				8, 51, 10, 1,
				0.5, 7, 0.25, 4
			*/

			name:   "Inverse of a 4x44 matrix",
			matrix: NewMatrix4x4([16]float64{10, 1, 10, 1, 0.10, 10, 0.2, 20, 8, 51, 10, 1, 0.5, 7, 0.25, 4}),
			want:   NewMatrix4x4([16]float64{0.167781, -0.423769, -0.212556, 2.130041, -0.013288, -0.016950, 0.011497745226596731377, 0.085201650302235583847, -0.06709949787315700261, 0.4198249485613159773, 0.21208729118026460273, -2.1353716911333567873, 0.0064764767966226372851, 0.056395987249602882696, -0.0068069636784256031385, -0.031897314527563672031}),
		},
		{
			/*
				35, 14, 93, 62,
				85, 84, 56, 84,
				84, 62, 96, 92,
				49, 46, 22, 28
			*/

			name:   "minor of a 4x44 matrix",
			matrix: NewMatrix4x4([16]float64{35, 14, 93, 62, 85, 84, 56, 84, 84, 62, 96, 92, 49, 46, 22, 28}),
			want:   NewMatrix4x4([16]float64{-0.0926712, -0.0795421, 0.139519, -0.0145919, 0.0955846, 0.0711247, -0.148509, 0.0629328, 0.058404, 0.00908792, -0.0630669, 0.0506327, -0.0407463, 0.0152105, 0.049374, -0.0819223}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.matrix.Inverse()

			if !got.Equal(tt.want) {
				t.Errorf("%s test failed, \n original \n%s \nwanted \n%s \nbut got \n%s\n", tt.name, tt.matrix.Print(), tt.want.Print(), got.Print())
			}

		})
	}
}

func TestInverse3x3(t *testing.T) {
	tests := []struct {
		name   string
		matrix Matrix3x3
		want   Matrix3x3
	}{
		{
			/*
				0.110777, -0.106773, -0.00500501
				,0.00333667, -0.0433767, 0.0500501
				,-0.0111111, 0.111111, 0
			*/

			name:   "Inverse of a 3x3 matrix",
			matrix: NewMatrix3x3([9]float64{10, 1, 10, 1, 0.10, 10, 0.2, 20, 8}),
			want:   NewMatrix3x3([9]float64{0.110777, -0.106773, -0.00500501, 0.00333667, -0.0433767, 0.0500501, -0.0111111, 0.111111, 0}),
		},
		{
			/*
				0.00223714, 0.176734, -0.179211
				,-0.0134228, -0.0604027, 0.0752637
				,0.0119314, -0.0574198, 0.0561148
			*/

			name:   "Inverse of a 3x3 matrix",
			matrix: NewMatrix3x3([9]float64{35, 14, 93, 62, 85, 84, 56, 84, 84}),
			want:   NewMatrix3x3([9]float64{0.00223714, 0.176734, -0.179211, -0.0134228, -0.0604027, 0.0752637, 0.0119314, -0.0574198, 0.0561148}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.matrix.Inverse()

			if !got.Equal(tt.want) {
				t.Errorf("%s test failed, \n original \n%s \nwanted \n%s \nbut got \n%s\n", tt.name, tt.matrix.Print(), tt.want.Print(), got.Print())
			}

		})
	}
}

func TestInverse4x4AndBack(t *testing.T) {
	tests := []struct {
		name    string
		matrix  Matrix4x4
		matrix2 Matrix4x4
	}{
		{
			/*
				10, 1, 10, 1,
				0.10, 10, 0.2, 20,
				8, 51, 10, 1,
				0.5, 7, 0.25, 4
			*/

			name: "Inverse and back of a 4x4 matrix",

			matrix:  NewMatrix4x4([16]float64{10, 1, 10, 1, 0.10, 10, 0.2, 20, 8, 51, 10, 1, 0.5, 7, 0.25, 4}),
			matrix2: NewMatrix4x4([16]float64{35, 14, 93, 62, 85, 84, 56, 84, 84, 62, 96, 92, 49, 46, 22, 28}),
		},
		{
			/*
				10, 1, 10, 1,
				0.10, 10, 0.2, 20,
				8, 51, 10, 1,
				0.5, 7, 0.25, 4
			*/

			name:    "Inverse and back of a 4x4 matrix",
			matrix:  NewMatrix4x4([16]float64{35, 14, 93, 62, 85, 84, 56, 84, 84, 62, 96, 92, 49, 46, 22, 28}),
			matrix2: NewMatrix4x4([16]float64{10, 1, 10, 1, 0.10, 10, 0.2, 20, 8, 51, 10, 1, 0.5, 7, 0.25, 4}),
		},
		{
			/*
				35, 14, 93, 62,
				85, 84, 56, 84,
				84, 62, 96, 92,
				49, 46, 22, 28
			*/

			name:    "Inverse and back of a 4x4 matrix",
			matrix:  NewMatrix4x4([16]float64{35, 14, 93, 62, 85, 84, 56, 84, 84, 62, 96, 92, 49, 46, 22, 28}),
			matrix2: NewMatrix4x4([16]float64{-0.0926712, -0.0795421, 0.139519, -0.0145919, 0.0955846, 0.0711247, -0.148509, 0.0629328, 0.058404, 0.00908792, -0.0630669, 0.0506327, -0.0407463, 0.0152105, 0.049374, -0.0819223}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inverse := tt.matrix.Inverse()
			multi := tt.matrix2.Multiply(tt.matrix)
			got := multi.Multiply(inverse)

			if !got.Equal(tt.matrix2) {
				t.Errorf("Matrix did not match: %s\noriginal is:\n%s", got.Print(), tt.matrix.Print())
			}

		})
	}
}
