package shapes

import (
	"testing"

	pm "github.com/Martin-Martinez4/ray-tracer-challenge-go/primitive_math"
)

func TestNewRay(t *testing.T) {
	tests := []struct {
		name      string
		origin    [3]float64
		direction [3]float64
		want      Ray
	}{
		{
			name:      "Create a Ray with the correct origin and direction",
			origin:    [3]float64{0, 1, 0},
			direction: [3]float64{0, 0, 1},
			want:      Ray{origin: pm.Point(0, 1, 0), direction: pm.Vector(0, 0, 1)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := NewRay(tt.origin, tt.direction)

			if !got.Equal(tt.want) {
				t.Errorf("%s did not pass: \nGot: %s \nWanted: %s", tt.name, got.Print(), tt.want.Print())
			}

		})
	}
}

func TestPosition(t *testing.T) {
	tests := []struct {
		name   string
		ray    Ray
		tvalue float64
		want   pm.Tuple
	}{
		{
			name:   "Create a point from a distance",
			ray:    NewRay([3]float64{2, 3, 4}, [3]float64{1, 0, 0}),
			tvalue: 0,
			want:   pm.Point(2, 3, 4),
		},
		{
			name:   "Create a point from a distance",
			ray:    NewRay([3]float64{2, 3, 4}, [3]float64{1, 0, 0}),
			tvalue: 1,
			want:   pm.Point(3, 3, 4),
		},
		{
			name:   "Create a point from a distance",
			ray:    NewRay([3]float64{2, 3, 4}, [3]float64{1, 0, 0}),
			tvalue: 10,
			want:   pm.Point(12, 3, 4),
		},
		{
			name:   "Create a point from a distance",
			ray:    NewRay([3]float64{2, 3, 4}, [3]float64{1, 0, 0}),
			tvalue: 2.5,
			want:   pm.Point(4.5, 3, 4),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := tt.ray.Position(tt.tvalue)

			if !got.Equal(tt.want) {
				t.Errorf("%s did not pass: \nGot: %s \nWanted: %s", tt.name, got.Print(), tt.want.Print())
			}

		})
	}
}
