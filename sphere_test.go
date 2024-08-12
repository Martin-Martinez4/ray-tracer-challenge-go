package main

import (
	"math"
	"testing"
)

func TestNormalAt(t *testing.T) {

	tests := []struct {
		name       string
		sphere     *Sphere
		transforms []string
		point      Tuple
		args       [][]float64
		want       Tuple
	}{
		{
			name:       "compute the normal at at point on the x-axis",
			sphere:     NewSphere(),
			transforms: []string{},
			point:      Point(1, 0, 0),
			args:       [][]float64{},
			want:       Vector(1, 0, 0),
		},
		{
			name:       "compute the normal at at point on the y-axis",
			sphere:     NewSphere(),
			transforms: []string{},
			point:      Point(0, 1, 0),
			args:       [][]float64{},
			want:       Vector(0, 1, 0),
		},
		{
			name:       "compute the normal at at point on the z-axis",
			sphere:     NewSphere(),
			transforms: []string{},
			point:      Point(0, 0, 1),
			args:       [][]float64{},
			want:       Vector(0, 0, 1),
		},
		{
			name:       "compute the normal at at point on the z-axis",
			sphere:     NewSphere(),
			transforms: []string{},
			point:      Point(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3),
			args:       [][]float64{},
			want:       Vector(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3),
		},
		{
			name:       "compute the normal at at point on the z-axis",
			sphere:     NewSphere(),
			transforms: []string{},
			point:      NewSphere().NormalAt(Point(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3)),
			args:       [][]float64{},
			want:       Vector(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3),
		},
		{
			name:       "compute the normal on a translated sphere",
			sphere:     NewSphere(),
			transforms: []string{"translate"},
			point:      Point(0, 1.70711, -0.70711),
			args:       [][]float64{{0, 1, 0}},
			want:       Vector(0, 0.70711, -0.70711),
		},
		{
			name:       "compute the normal on a scaled and rotated sphere",
			sphere:     NewSphere(),
			transforms: []string{"rotateZ", "scale"},
			point:      Point(0, math.Sqrt(2)/2, -math.Sqrt(2)/2),
			args:       [][]float64{{math.Pi / 5}, {1, 0.5, 1}},
			want:       Vector(0, 0.97014, -0.24254),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			for i := 0; i < len(tt.transforms); i++ {

				if tt.transforms[i] == "scale" {

					tt.sphere.Scale(tt.args[i][0], tt.args[i][1], tt.args[i][2])

				} else if tt.transforms[i] == "translate" {
					tt.sphere.Translate(tt.args[i][0], tt.args[i][1], tt.args[i][2])

				} else if tt.transforms[i] == "rotateZ" {
					tt.sphere.RotationAlongZ(tt.args[i][0])
				} else {
					t.Errorf("%s is not a vaild transformation option", tt.transforms[i])
					return
				}
			}

			got := tt.sphere.NormalAt(tt.point)

			if !got.Equal(tt.want) {
				t.Errorf("%s did not pass,\ngot:\n%s\nwanted:\n%s", tt.name, got.Print(), tt.want.Print())
			}

		})
	}
}
