package main

import (
	"math"
	"testing"
)

func TestTupleTranslate(t *testing.T) {
	tests := []struct {
		name  string
		tuple Tuple
		args  [3]float64
		want  Tuple
	}{
		{
			name:  "Translating a vector should return the same vector",
			tuple: Vector(-3, 4, 5),
			args:  [3]float64{5, -3, 2},
			want:  Vector(-3, 4, 5),
		},
		{
			name:  "Translating a point should return the transformed point",
			tuple: Point(-3, 4, 5),
			args:  [3]float64{5, -3, 2},
			want:  Point(2, 1, 7),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := tt.tuple.Translate(tt.args[0], tt.args[1], tt.args[2])

			if !got.Equal(tt.want) || IsPoint(tt.tuple) != IsPoint(got) || IsVector(tt.tuple) != IsVector(got) {
				t.Errorf("Tuple did not match: %s\noriginal is:\n%s", got.Print(), tt.tuple.Print())
			}

		})
	}
}

func TestTupleInverseTranslate(t *testing.T) {
	tests := []struct {
		name  string
		tuple Tuple
		args  [3]float64
		want  Tuple
	}{
		{
			name:  "Translating a vector should return the same vector",
			tuple: Vector(-3, 4, 5),
			args:  [3]float64{5, -3, 2},
			want:  Vector(-3, 4, 5),
		},
		{
			name:  "Translating a point should return the transformed point",
			tuple: Point(-3, 4, 5),
			args:  [3]float64{5, -3, 2},
			want:  Point(-8, 7, 3),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := tt.tuple.TranslateInverse(tt.args[0], tt.args[1], tt.args[2])

			if !got.Equal(tt.want) || IsPoint(tt.tuple) != IsPoint(got) || IsVector(tt.tuple) != IsVector(got) {
				t.Errorf("Tuple did not match: %s\noriginal is:\n%s", got.Print(), tt.tuple.Print())
			}

		})
	}
}

func TestTupleScale(t *testing.T) {
	tests := []struct {
		name  string
		tuple Tuple
		args  [3]float64
		want  Tuple
	}{
		{
			name:  "Scaling a vector should return the transformed vector",
			tuple: Vector(-3, 4, 5),
			args:  [3]float64{5, -3, 2},
			want:  Vector(-15, -12, 10),
		},
		{
			name:  "Scaling a point should return the transformed point",
			tuple: Point(2, 3, 4),
			args:  [3]float64{-4, 6, 8},
			want:  Point(-8, 18, 32),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := tt.tuple.Scale(tt.args[0], tt.args[1], tt.args[2])

			if !got.Equal(tt.want) || IsPoint(tt.tuple) != IsPoint(got) || IsVector(tt.tuple) != IsVector(got) {
				t.Errorf("Tuple did not match: %s\noriginal is:\n%s", got.Print(), tt.tuple.Print())
			}

		})
	}
}

func TestTupleScaleInverse(t *testing.T) {
	tests := []struct {
		name  string
		tuple Tuple
		args  [3]float64
		want  Tuple
	}{
		{
			name:  "Scaling a vector should return the transformed vector",
			tuple: Vector(-15, -12, 10),
			args:  [3]float64{5, -3, 2},
			want:  Vector(-3, 4, 5),
		},
		{
			name:  "Scaling a point should return the transformed point",
			tuple: Point(-8, 18, 32),
			args:  [3]float64{-4, 6, 8},
			want:  Point(2, 3, 4),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := tt.tuple.ScaleInverse(tt.args[0], tt.args[1], tt.args[2])

			if !got.Equal(tt.want) || IsPoint(tt.tuple) != IsPoint(got) || IsVector(tt.tuple) != IsVector(got) {
				t.Errorf("Tuple did not match: %s\noriginal is:\n%s", got.Print(), tt.tuple.Print())
			}

		})
	}
}

func TestReflectX(t *testing.T) {
	tests := []struct {
		name  string
		tuple Tuple
		args  [3]float64
		want  Tuple
	}{
		{
			name:  "Scaling a vector should return the transformed vector",
			tuple: Vector(-15, -12, 10),
			want:  Vector(15, -12, 10),
		},
		{
			name:  "Scaling a point should return the transformed point",
			tuple: Point(-8, 18, 32),
			want:  Point(8, 18, 32),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := tt.tuple.ReflectX()

			if !got.Equal(tt.want) || IsPoint(tt.tuple) != IsPoint(got) || IsVector(tt.tuple) != IsVector(got) {
				t.Errorf("Tuple did not match: %s\noriginal is:\n%s", got.Print(), tt.tuple.Print())
			}

		})
	}
}

func TestReflectY(t *testing.T) {
	tests := []struct {
		name  string
		tuple Tuple
		args  [3]float64
		want  Tuple
	}{
		{
			name:  "Scaling a vector should return the transformed vector",
			tuple: Vector(15, -12, 10),
			want:  Vector(15, 12, 10),
		},
		{
			name:  "Scaling a point should return the transformed point",
			tuple: Point(8, 18, 32),
			want:  Point(8, -18, 32),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := tt.tuple.ReflectY()

			if !got.Equal(tt.want) || IsPoint(tt.tuple) != IsPoint(got) || IsVector(tt.tuple) != IsVector(got) {
				t.Errorf("Tuple did not match: %s\noriginal is:\n%s", got.Print(), tt.tuple.Print())
			}

		})
	}
}

func TestReflectZ(t *testing.T) {
	tests := []struct {
		name  string
		tuple Tuple
		args  [3]float64
		want  Tuple
	}{
		{
			name:  "Scaling a vector should return the transformed vector",
			tuple: Vector(15, 12, -10),
			want:  Vector(15, 12, 10),
		},
		{
			name:  "Scaling a point should return the transformed point",
			tuple: Point(8, -18, 32),
			want:  Point(8, -18, -32),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := tt.tuple.ReflectZ()

			if !got.Equal(tt.want) || IsPoint(tt.tuple) != IsPoint(got) || IsVector(tt.tuple) != IsVector(got) {
				t.Errorf("Tuple did not match: %s\noriginal is:\n%s", got.Print(), tt.tuple.Print())
			}

		})
	}
}

func TestRotationAlongX(t *testing.T) {
	tests := []struct {
		name    string
		tuple   Tuple
		radians float64
		want    Tuple
	}{
		{
			name:    "Rotating a point around the X axis should return the desired point",
			tuple:   Point(0, 1, 0),
			radians: math.Pi / 4,
			want:    Point(0, math.Sqrt(2)/2, math.Sqrt(2)/2),
		},
		{
			name:    "Rotating a point around the X axis should return the desired point",
			tuple:   Point(0, 1, 0),
			radians: math.Pi / 2,
			want:    Point(0, 0, 1),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := tt.tuple.RotationAlongX(tt.radians)

			if !got.Equal(tt.want) || IsPoint(tt.tuple) != IsPoint(got) || IsVector(tt.tuple) != IsVector(got) {
				t.Errorf("Tuple did not match: %s\noriginal is:\n%s", got.Print(), tt.tuple.Print())
			}

		})
	}
}

func TestRotationAlongY(t *testing.T) {
	tests := []struct {
		name    string
		tuple   Tuple
		radians float64
		want    Tuple
	}{
		{
			name:    "Rotating a point around the Y axis should return the desired point",
			tuple:   Point(0, 0, 1),
			radians: math.Pi / 4,
			want:    Point(math.Sqrt(2)/2, 0, math.Sqrt(2)/2),
		},
		{
			name:    "Rotating a point around the Y axis should return the desired point",
			tuple:   Point(0, 0, 1),
			radians: math.Pi / 2,
			want:    Point(1, 0, 0),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := tt.tuple.RotationAlongY(tt.radians)

			if !got.Equal(tt.want) || IsPoint(tt.tuple) != IsPoint(got) || IsVector(tt.tuple) != IsVector(got) {
				t.Errorf("\nTuple did not match: %s\noriginal is:\n%s\nwanted: %s\n", got.Print(), tt.tuple.Print(), tt.want.Print())
			}

		})
	}
}

func TestChainingTransforms(t *testing.T) {
	tests := []struct {
		name       string
		tuple      Tuple
		transforms []string
		arguments  [][]float64
		want       Tuple
	}{
		{
			name:       "Transformations should be able to be done in succession",
			tuple:      Point(1, 0, 1),
			transforms: []string{"rotateX", "scale", "translate"},
			arguments:  [][]float64{{math.Pi / 2}, {5, 5, 5}, {10, 5, 7}},
			want:       Point(15, 0, 7),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := tt.tuple

			for i, transform := range tt.transforms {

				switch transform {
				case "rotateX":
					if len(tt.arguments[i]) != 1 {
						t.Errorf("%s #%d needs %d argument but was given %d", tt.transforms[i], i, 1, len(tt.arguments[i]))
					}
					got = got.RotationAlongX(tt.arguments[i][0])
				case "scale":
					if len(tt.arguments[i]) != 3 {
						t.Errorf("%s #%d needs %d argument but was given %d", tt.transforms[i], i, 3, len(tt.arguments[i]))
					}
					got = got.Scale(tt.arguments[i][0], tt.arguments[i][1], tt.arguments[i][2])
				case "translate":
					if len(tt.arguments[i]) != 3 {
						t.Errorf("%s #%d needs %d argument but was given %d", tt.transforms[i], i, 3, len(tt.arguments[i]))
					}
					got = got.Translate(tt.arguments[i][0], tt.arguments[i][1], tt.arguments[i][2])

				default:
					t.Errorf("%s #%d is not a valid transformation", tt.transforms[i], i)
				}
			}

			if !got.Equal(tt.want) || IsPoint(tt.tuple) != IsPoint(got) || IsVector(tt.tuple) != IsVector(got) {
				t.Errorf("Tuple did not match: %s\noriginal is:\n%s", got.Print(), tt.tuple.Print())
			}

		})
	}
}

func TestReflectBy(t *testing.T) {
	tests := []struct {
		name    string
		tuple   Tuple
		reflect Tuple
		want    Tuple
	}{
		{
			name:    "Reflecting a vector should return a vector",
			tuple:   Vector(1, -1, 0),
			reflect: Vector(0, 1, 0),
			want:    Vector(1, 1, 0),
		},
		{
			name:    "Reflecting a vector should return a vector",
			tuple:   Vector(0, -1, 0),
			reflect: Vector(math.Sqrt(2)/2, math.Sqrt(2)/2, 0),
			want:    Vector(1, 0, 0),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := tt.tuple.ReflectBy(tt.reflect)

			if !got.Equal(tt.want) || IsPoint(tt.tuple) != IsPoint(got) || IsVector(tt.tuple) != IsVector(got) {
				t.Errorf("\nTuple did not match: %s\nOriginal is: %s \nWanted: %s", got.Print(), tt.tuple.Print(), tt.want.Print())
			}

		})
	}
}
