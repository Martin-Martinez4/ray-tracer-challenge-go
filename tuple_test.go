package main

import (
	"math"
	"testing"
)

/*
Scenerio a tuple with w = 1.0 is a point
given (4.3, -4.2, 3.1, 1.0)
	a.x should equal 4.3
	a.y should equal -4.2
	a.z should equal 3.1
	a should be a point
	a should not be a vector

Idea: Create tuple function that takes in a list of floats and returns the concrete tuple
*/

func TestIsPoint(t *testing.T) {
	tests := []struct {
		name string
		args Tuple
		want bool
	}{
		{
			name: "should be a point",
			args: Tuple{4.3, -4.2, 3.1, 1},
			want: true,
		},
		{
			name: "should be a point",
			args: Tuple{5.0, -0.2, 6, 1},
			want: true,
		},
		{
			name: "should be false",
			args: Tuple{5.0, -0.2, 6, 0},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsPoint(tt.args)

			if got != tt.want {
				t.Errorf("IsPoint returned %v, wanted %v", got, tt.want)

			}
		})
	}
}

func TestIsVector(t *testing.T) {
	tests := []struct {
		name string
		args Tuple
		want bool
	}{
		{
			name: "should be a point",
			args: Tuple{4.3, -4.2, 3.1, 1},
			want: false,
		},
		{
			name: "should be a point",
			args: Tuple{5.0, -0.2, 6, 1},
			want: false,
		},
		{
			name: "should be false",
			args: Tuple{5.0, -0.2, 6, 0},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsVector(tt.args)

			if got != tt.want {
				t.Errorf("IsPoint returned %v, wanted %v", got, tt.want)

			}
		})
	}
}

func TestPoint(t *testing.T) {
	tests := []struct {
		name string
		args []float64
		want Tuple
	}{
		{
			name: "should be a point",
			args: []float64{4.3, -4.2, 3.1},
			want: Tuple{4.3, -4.2, 3.1, 1},
		},
		{
			name: "should be a point",
			args: []float64{5.0, -0.2, 6, 1},
			want: Tuple{5.0, -0.2, 6, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Point(tt.args[0], tt.args[1], tt.args[2])

			if got.x != tt.want.x || got.y != tt.want.y || got.z != tt.want.z || !IsPoint(got) {
				t.Errorf("Point returned %v, wanted %v", got, tt.want)

			}
		})
	}
}

func TestVector(t *testing.T) {
	tests := []struct {
		name string
		args []float64
		want Tuple
	}{
		{
			name: "should be a vector",
			args: []float64{4.3, -4.2, 3.1},
			want: Tuple{4.3, -4.2, 3.1, 0},
		},
		{
			name: "should be a vector",
			args: []float64{5.0, -0.2, 6, 1},
			want: Tuple{5.0, -0.2, 6, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Vector(tt.args[0], tt.args[1], tt.args[2])

			if got.x != tt.want.x || got.y != tt.want.y || got.z != tt.want.z || !IsVector(got) {
				t.Errorf("Point returned %v, wanted %v", got, tt.want)

			}
		})
	}
}

func TestTupleEqual(t *testing.T) {
	tests := []struct {
		name string
		args []Tuple
		want bool
	}{
		{
			name: "points should be equal",
			args: []Tuple{Point(1, 2, 3), Point(1, 2, 3)},
			want: true,
		},
		{
			name: "points should not be equal",
			args: []Tuple{Point(1, 2, 3), Point(1, 2, 4)},
			want: false,
		},
		{
			name: "vector and point should not be equal",
			args: []Tuple{Point(1, 2, 3), Vector(1, 2, 3)},
			want: false,
		},
		{
			name: "vectors should be equal",
			args: []Tuple{Vector(1, 2, 3), Vector(1, 2, 3)},
			want: true,
		},
		{
			name: "vectors should be equal",
			args: []Tuple{Vector(1, 2, 3), Vector(1, 2, 4)},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args[0].Equal(tt.args[1])

			if got != tt.want {
				t.Errorf("Equal check returned %v, wanted %v", got, tt.want)

			}
		})
	}

}

func TestAddTuple(t *testing.T) {
	tests := []struct {
		name string
		args []Tuple
		want Tuple
	}{
		{
			name: "points can be added",
			args: []Tuple{Point(1, 2, 3), Point(1, 2, 3)},
			want: Tuple{2, 4, 6, 2},
		},
		{
			name: "points should not be equal",
			args: []Tuple{Point(1, 2, 3), Vector(1, 2, 4)},
			want: Tuple{2, 4, 7, 1},
		},
		{
			name: "Vector and Vector should be added",
			args: []Tuple{Vector(1, 2, 3), Vector(1, 2, 3)},
			want: Tuple{2, 4, 6, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args[0].Add(tt.args[1])

			if !got.Equal(tt.want) {
				t.Errorf("Equal check returned %v, wanted %v", got, tt.want)

			}
		})
	}

}

func TestSubtractTuple(t *testing.T) {
	tests := []struct {
		name string
		args []Tuple
		want Tuple
	}{
		{
			name: "points can be added",
			args: []Tuple{Point(1, 2, 3), Point(1, 2, 3)},
			want: Tuple{0, 0, 0, 0},
		},
		{
			name: "points should not be equal",
			args: []Tuple{Point(1, 2, 3), Vector(1, 2, 4)},
			want: Tuple{0, 0, -1, 1},
		},
		{
			name: "Vector and Vector should be added",
			args: []Tuple{Vector(1, 2, 3), Vector(2, 4, 3)},
			want: Tuple{-1, -2, 0, 0},
		},
		{
			name: "Vector and Vector should be added",
			args: []Tuple{Vector(0, 0, 0), Vector(2, 4, 3)},
			want: Tuple{-2, -4, -3, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args[0].Subtract(tt.args[1])

			if !got.Equal(tt.want) {
				t.Errorf("Equal check returned %v, wanted %v", got, tt.want)

			}
		})
	}
}

func TestNegate(t *testing.T) {
	tests := []struct {
		name string
		args []Tuple
		want Tuple
	}{
		{
			name: "negate a point",
			args: []Tuple{Point(1, 2, 3)},
			want: Tuple{-1, -2, -3, -1},
		},
		{
			name: "negate a vector",
			args: []Tuple{Vector(1, 2, 4)},
			want: Tuple{-1, -2, -4, 0},
		},
		{
			name: "negate a negative value vector",
			args: []Tuple{Vector(-2, -4, -3)},
			want: Tuple{2, 4, 3, 0},
		},
		{
			name: "negate a mixed value point",
			args: []Tuple{Point(-2, 4, -3)},
			want: Tuple{2, -4, 3, -1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args[0].Negate()

			if !got.Equal(tt.want) {
				t.Errorf("Equal check returned %v, wanted %v", got, tt.want)

			}
		})
	}

}

// SMultiply means scalar multiply
func TestSMultiply(t *testing.T) {
	tests := []struct {
		name   string
		tuple  Tuple
		scalar float64
		want   Tuple
	}{
		{
			name:   "scalar multiply a point",
			tuple:  Point(1, 2, 3),
			scalar: 4,
			want:   Tuple{4, 8, 12, 4},
		},
		{
			name:   "scalar multiply a vector",
			tuple:  Vector(1, 2, 3),
			scalar: 4,
			want:   Tuple{4, 8, 12, 0},
		},
		{
			name:   "scalar multiply a Tuple",
			tuple:  Tuple{0, 4, 5, 8},
			scalar: 2,
			want:   Tuple{0, 8, 10, 16},
		},
		{
			name:   "scalar multiply a tuple by a negative",
			tuple:  Point(1, 2, 3),
			scalar: -10,
			want:   Tuple{-10, -20, -30, -10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.tuple.SMultiply(tt.scalar)

			if !got.Equal(tt.want) {
				t.Errorf("Equal check returned %v, wanted %v", got, tt.want)

			}
		})
	}

}

func TestSDivide(t *testing.T) {
	tests := []struct {
		name   string
		tuple  Tuple
		scalar float64
		want   Tuple
	}{
		{
			name:   "scalar divide a point",
			tuple:  Point(1, 2, 3),
			scalar: 4,
			want:   Tuple{0.25, 0.5, 0.75, 0.25},
		},
		{
			name:   "scalar divide a vector",
			tuple:  Vector(1, 2, 3),
			scalar: 4,
			want:   Tuple{0.25, 0.5, .75, 0},
		},
		{
			name:   "scalar divide a Tuple",
			tuple:  Tuple{0, 4, 5, 8},
			scalar: 2,
			want:   Tuple{0, 2, 2.5, 4},
		},
		{
			name:   "scalar divide a tuple by a negative",
			tuple:  Point(1, 2, 3),
			scalar: -10,
			want:   Tuple{-0.1, -0.2, -0.3, -0.1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.tuple.SDivide(tt.scalar)

			if !got.Equal(tt.want) {
				t.Errorf("Equal check returned %v, wanted %v", got, tt.want)

			}
		})
	}

}

func TestMagnitude(t *testing.T) {
	tests := []struct {
		name  string
		tuple Tuple
		want  float64
	}{
		{
			name:  "magnitude should be 1",
			tuple: Vector(1, 0, 0),
			want:  1,
		},
		{
			name:  "magnitude should be 1",
			tuple: Vector(0, 1, 0),
			want:  1,
		},
		{
			name:  "magnitude should be math.sqrt(14)",
			tuple: Vector(1, 2, 3),
			want:  float64(math.Sqrt(14)),
		},
		{
			name:  "magnitude should be math.sqrt(14)",
			tuple: Vector(-1, -2, -3),
			want:  float64(math.Sqrt(14)),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.tuple.Magnitude()

			if !AreFloatsEqual(got, tt.want) {
				t.Errorf("Equal check returned %v, wanted %v", got, tt.want)

			}
		})
	}
}

func TestNormalize(t *testing.T) {
	tests := []struct {
		name  string
		tuple Tuple
		want  Tuple
	}{
		{
			name:  "magnitude should be 1",
			tuple: Vector(1, 0, 0),
			want:  Tuple{1, 0, 0, 0},
		},
		{
			name:  "magnitude should be 1",
			tuple: Vector(0, 1, 0),
			want:  Tuple{0, 1, 0, 0},
		},
		{
			name:  "magnitude should be math.sqrt(14)",
			tuple: Vector(1, 2, 3),
			want:  Tuple{float64(1 / math.Sqrt(14)), float64(2 / math.Sqrt(14)), float64(3 / math.Sqrt(14)), 0},
		},
		{
			name:  "magnitude should be math.sqrt(14)",
			tuple: Vector(-1, -2, -3),
			want:  Tuple{float64(-1 / math.Sqrt(14)), float64(-2 / math.Sqrt(14)), float64(-3 / math.Sqrt(14)), 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Normalize(tt.tuple)

			if !got.Equal(tt.want) {
				t.Errorf("Equal check returned %v, wanted %v", got, tt.want)

			}
		})
	}
}

func TestDot(t *testing.T) {
	tests := []struct {
		name   string
		tuple1 Tuple
		tuple2 Tuple
		want   float64
	}{
		{
			name:   "magnitude should be 1",
			tuple1: Vector(1, 2, 3),
			tuple2: Vector(2, 3, 4),
			want:   20,
		},
		{
			name:   "magnitude should be 1",
			tuple1: Vector(1, 2, 3),
			tuple2: Vector(3, 4, 5),
			want:   26,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Dot(tt.tuple1, tt.tuple2)

			if !AreFloatsEqual(got, tt.want) {
				t.Errorf("Equal check returned %v, wanted %v", got, tt.want)

			}
		})
	}
}

func TestCross(t *testing.T) {
	tests := []struct {
		name   string
		tuple1 Tuple
		tuple2 Tuple
		want   Tuple
	}{
		{
			name:   "magnitude should be 1",
			tuple1: Vector(1, 2, 3),
			tuple2: Vector(2, 3, 4),
			want:   Vector(-1, 2, -1),
		},
		{
			name:   "magnitude should be 1",
			tuple1: Vector(2, 3, 4),
			tuple2: Vector(1, 2, 3),
			want:   Vector(1, -2, 1),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Cross(tt.tuple1, tt.tuple2)

			if !got.Equal(tt.want) {
				t.Errorf("Equal check returned %v, wanted %v", got, tt.want)

			}
		})
	}
}
