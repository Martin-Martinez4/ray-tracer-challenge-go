package shapes

import (
	"fmt"
	"math"
	"testing"

	pm "github.com/Martin-Martinez4/ray-tracer-challenge-go/primitive_math"
)

func TestRayConeIntersect(T *testing.T) {

	cone := NewCone()

	tests := []struct {
		name      string
		origin    pm.Tuple
		direction pm.Tuple
		cone      *Cone
		want      []Intersection
	}{
		{
			name:      "the ray hits the cone 1",
			cone:      cone,
			origin:    pm.Point(0, 0, -5),
			direction: pm.Vector(0, 0, 1),
			want:      []Intersection{NewIntersection(5, cone), NewIntersection(5, cone)},
		},
		{
			name:      "the ray hits the cone 2",
			cone:      cone,
			origin:    pm.Point(0, 0, -5),
			direction: pm.Vector(1, 1, 1),
			want:      []Intersection{NewIntersection(8.66025, cone), NewIntersection(8.66025, cone)},
		},
		{
			name:      "the ray hits the cone 3",
			cone:      cone,
			origin:    pm.Point(1, 1, -5),
			direction: pm.Vector(-0.5, -1, 1),
			want:      []Intersection{NewIntersection(4.55006, cone), NewIntersection(49.44994, cone)},
		},
		{
			name:      "the ray hits the cone once 1",
			cone:      cone,
			origin:    pm.Point(0, 0, -1),
			direction: pm.Vector(0, 1, 1),
			want:      []Intersection{NewIntersection(0.35355, cone)},
		},
	}

	for i, tt := range tests {
		T.Run(fmt.Sprintf("%d: %s", i, tt.name), func(t *testing.T) {

			normed := pm.Normalize(tt.direction)

			ray := NewRay([3]float64{tt.origin.X, tt.origin.Y, tt.origin.Z}, [3]float64{normed.X, normed.Y, normed.Z})

			got := tt.cone.LocalIntersect(ray)

			if !got.Equal(Intersections{intersections: tt.want}) {
				t.Errorf("%d: \nwant: %v \ngot: %v \ndo not match", i, tt.want, got)
			}

		})
	}
}

func TestRayConeCapIntersect(T *testing.T) {

	cone := NewCone()
	cone.Maximum = 0.5
	cone.Minimum = -0.5
	cone.Closed = true

	tests := []struct {
		name      string
		origin    pm.Tuple
		direction pm.Tuple
		cone      *Cone
		want      []Intersection
	}{
		{
			name:      "the ray intersects the truncated cone cap 1",
			cone:      cone,
			origin:    pm.Point(0, 0, -5),
			direction: pm.Vector(0, 1, 0),
			want:      []Intersection{},
		},
		{
			name:      "the ray intersects the truncated cone cap 2",
			cone:      cone,
			origin:    pm.Point(0, 0, -0.25),
			direction: pm.Vector(0, 1, 1),
			want:      []Intersection{NewIntersection(0.08838834764831845, cone), NewIntersection(0.7071067811865476, cone)},
		},
		{
			name:      "the ray intersects the truncated cone cap 3",
			cone:      cone,
			origin:    pm.Point(0, 0, -0.25),
			direction: pm.Vector(0, 1, 0),
			want:      []Intersection{NewIntersection(-0.5, cone), NewIntersection(-0.25, cone), NewIntersection(0.25, cone), NewIntersection(0.5, cone)},
		},
	}

	for i, tt := range tests {
		T.Run(fmt.Sprintf("%d: %s", i, tt.name), func(t *testing.T) {

			normed := pm.Normalize(tt.direction)

			ray := NewRay([3]float64{tt.origin.X, tt.origin.Y, tt.origin.Z}, [3]float64{normed.X, normed.Y, normed.Z})

			got := tt.cone.LocalIntersect(ray)

			if !got.Equal(Intersections{intersections: tt.want}) {
				t.Errorf("\ntest %d: \nwant: %v \ngot: %v \ndo not match", i, tt.want, got)
			}

		})
	}
}

func TestRayConeLocalNormal(T *testing.T) {

	cone := NewCone()

	tests := []struct {
		name  string
		point pm.Tuple
		cone  *Cone
		want  pm.Tuple
	}{
		{
			name:  "cone local normal 1",
			cone:  cone,
			point: pm.Point(0, 0, 0),
			want:  pm.Vector(0, 0, 0),
		},
		{
			name:  "cone local normal 2",
			cone:  cone,
			point: pm.Point(1, 1, 1),
			want:  pm.Vector(1, -math.Sqrt(2), 1),
		},
		{
			name:  "cone local normal 3",
			cone:  cone,
			point: pm.Point(-1, -1, 0),
			want:  pm.Vector(-1, 1, 0),
		},
	}

	for i, tt := range tests {
		T.Run(fmt.Sprintf("%d: %s", i, tt.name), func(t *testing.T) {

			got := tt.cone.LocalNormalAt(tt.point, nil, nil)

			if !got.Equal(tt.want) {
				t.Errorf("%d: \nwant: %s \ngot: %s \ndo not match", i, tt.want.Print(), got.Print())
			}

		})
	}
}
