package main

import (
	"fmt"
	"math"
	"testing"
)

func TestRayConeIntersect(T *testing.T) {

	cone := NewCone()

	tests := []struct {
		name      string
		origin    Tuple
		direction Tuple
		cone      *Cone
		want      []Intersection
	}{
		{
			name:      "the ray hits the cone 1",
			cone:      cone,
			origin:    Point(0, 0, -5),
			direction: Vector(0, 0, 1),
			want:      []Intersection{{5, cone}, {5, cone}},
		},
		{
			name:      "the ray hits the cone 2",
			cone:      cone,
			origin:    Point(0, 0, -5),
			direction: Vector(1, 1, 1),
			want:      []Intersection{{8.66025, cone}, {8.66025, cone}},
		},
		{
			name:      "the ray hits the cone 3",
			cone:      cone,
			origin:    Point(1, 1, -5),
			direction: Vector(-0.5, -1, 1),
			want:      []Intersection{{4.55006, cone}, {49.44994, cone}},
		},
		{
			name:      "the ray hits the cone once 1",
			cone:      cone,
			origin:    Point(0, 0, -1),
			direction: Vector(0, 1, 1),
			want:      []Intersection{{0.35355, cone}},
		},
	}

	for i, tt := range tests {
		T.Run(fmt.Sprintf("%d: %s", i, tt.name), func(t *testing.T) {

			normed := Normalize(tt.direction)

			ray := NewRay([3]float64{tt.origin.x, tt.origin.y, tt.origin.z}, [3]float64{normed.x, normed.y, normed.z})

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
		origin    Tuple
		direction Tuple
		cone      *Cone
		want      []Intersection
	}{
		{
			name:      "the ray intersects the truncated cone cap 1",
			cone:      cone,
			origin:    Point(0, 0, -5),
			direction: Vector(0, 1, 0),
			want:      []Intersection{},
		},
		{
			name:      "the ray intersects the truncated cone cap 2",
			cone:      cone,
			origin:    Point(0, 0, -0.25),
			direction: Vector(0, 1, 1),
			want:      []Intersection{{0.08838834764831845, cone}, {0.7071067811865476, cone}},
		},
		{
			name:      "the ray intersects the truncated cone cap 3",
			cone:      cone,
			origin:    Point(0, 0, -0.25),
			direction: Vector(0, 1, 0),
			want:      []Intersection{{-0.5, cone}, {-0.25, cone}, {0.25, cone}, {0.5, cone}},
		},
	}

	for i, tt := range tests {
		T.Run(fmt.Sprintf("%d: %s", i, tt.name), func(t *testing.T) {

			normed := Normalize(tt.direction)

			ray := NewRay([3]float64{tt.origin.x, tt.origin.y, tt.origin.z}, [3]float64{normed.x, normed.y, normed.z})

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
		point Tuple
		cone  *Cone
		want  Tuple
	}{
		{
			name:  "cone local normal 1",
			cone:  cone,
			point: Point(0, 0, 0),
			want:  Vector(0, 0, 0),
		},
		{
			name:  "cone local normal 2",
			cone:  cone,
			point: Point(1, 1, 1),
			want:  Vector(1, -math.Sqrt(2), 1),
		},
		{
			name:  "cone local normal 3",
			cone:  cone,
			point: Point(-1, -1, 0),
			want:  Vector(-1, 1, 0),
		},
	}

	for i, tt := range tests {
		T.Run(fmt.Sprintf("%d: %s", i, tt.name), func(t *testing.T) {

			got := tt.cone.LocalNormalAt(tt.point)

			if !got.Equal(tt.want) {
				t.Errorf("%d: \nwant: %s \ngot: %s \ndo not match", i, tt.want.Print(), got.Print())
			}

		})
	}
}
