package main

import (
	"fmt"
	"testing"
)

func TestRayCylinderIntersect(T *testing.T) {

	cylinder := NewCylinder()
	truncedCylinder := NewCylinder()
	truncedCylinder.Minimum = 1
	truncedCylinder.Maximum = 2

	tests := []struct {
		name      string
		origin    Tuple
		direction Tuple
		cylinder  *Cylinder
		want      []Intersection
	}{
		{
			name:      "the ray misses the cylinder 1",
			cylinder:  cylinder,
			origin:    Point(1, 0, 0),
			direction: Vector(0, 1, 0),
			want:      []Intersection{},
		},
		{
			name:      "the ray misses the cylinder 2",
			cylinder:  cylinder,
			origin:    Point(0, 0, 0),
			direction: Vector(0, 1, 0),
			want:      []Intersection{},
		},
		{
			name:      "the ray misses the cylinder 3",
			cylinder:  cylinder,
			origin:    Point(0, 0, -5),
			direction: Vector(1, 1, 1),
			want:      []Intersection{},
		},
		{
			name:      "the ray intersects the cylinder 1",
			cylinder:  cylinder,
			origin:    Point(1, 0, -5),
			direction: Vector(0, 0, 1),
			want:      []Intersection{NewIntersection(5, cylinder), NewIntersection(5, cylinder)},
		},
		{
			name:      "the ray intersects the cylinder 1",
			cylinder:  cylinder,
			origin:    Point(0, 0, -5),
			direction: Vector(0, 0, 1),
			want:      []Intersection{NewIntersection(4, cylinder), NewIntersection(6, cylinder)},
		},
		{
			name:      "the ray intersects the cylinder 1",
			cylinder:  cylinder,
			origin:    Point(0.5, 0, -5),
			direction: Vector(0.1, 1, 1),
			want:      []Intersection{NewIntersection(6.80798, cylinder), NewIntersection(7.08872, cylinder)},
		},
		{
			name:      "the ray intersects misses the truncated cylinder 1",
			cylinder:  truncedCylinder,
			origin:    Point(0, 1.5, 0),
			direction: Vector(0.1, 1, 0),
			want:      []Intersection{},
		},
		{
			name:      "the ray intersects misses the truncated cylinder 2",
			cylinder:  truncedCylinder,
			origin:    Point(0, 3, -5),
			direction: Vector(0, 0, 1),
			want:      []Intersection{},
		},
		{
			name:      "the ray intersects misses the truncated cylinder 3",
			cylinder:  truncedCylinder,
			origin:    Point(0, 2, -5),
			direction: Vector(0, 0, 1),
			want:      []Intersection{},
		},
		{
			name:      "the ray intersects hits the truncated cylinder 3",
			cylinder:  truncedCylinder,
			origin:    Point(0, 1.5, -2),
			direction: Vector(0, 0, 1),
			want:      []Intersection{NewIntersection(1, truncedCylinder), NewIntersection(3, truncedCylinder)},
		},
	}

	for i, tt := range tests {
		T.Run(fmt.Sprintf("%d: %s", i, tt.name), func(t *testing.T) {

			normed := Normalize(tt.direction)

			ray := NewRay([3]float64{tt.origin.x, tt.origin.y, tt.origin.z}, [3]float64{normed.x, normed.y, normed.z})

			got := tt.cylinder.LocalIntersect(ray)

			if !got.Equal(Intersections{intersections: tt.want}) {
				t.Errorf("%d: \nwant: %v \ngot: %v \ndo not match", i, tt.want, got)
			}

		})
	}
}

func TestRayCylinderLocalNormal(T *testing.T) {

	cylinder := NewCylinder()

	tests := []struct {
		name     string
		point    Tuple
		cylinder *Cylinder
		want     Tuple
	}{
		{
			name:     "cylinder local normal 1",
			cylinder: cylinder,
			point:    Point(1, 0, 0),
			want:     Vector(1, 0, 0),
		},
		{
			name:     "cylinder local normal 2",
			cylinder: cylinder,
			point:    Point(0, 5, -1),
			want:     Vector(0, 0, -1),
		},
		{
			name:     "cylinder local normal 3",
			cylinder: cylinder,
			point:    Point(0, -1, 1),
			want:     Vector(0, 0, 1),
		},
		{
			name:     "cylinder local normal 3",
			cylinder: cylinder,
			point:    Point(-1, 10, 0),
			want:     Vector(-1, 0, 0),
		},
	}

	for i, tt := range tests {
		T.Run(fmt.Sprintf("%d: %s", i, tt.name), func(t *testing.T) {

			got := tt.cylinder.LocalNormalAt(tt.point, nil, nil)

			if !got.Equal(tt.want) {
				t.Errorf("%d: \nwant: %s \ngot: %s \ndo not match", i, tt.want.Print(), got.Print())
			}

		})
	}
}

func TestRayCylinderCapIntersect(T *testing.T) {

	truncedCylinder := NewCylinder()
	truncedCylinder.Minimum = 1
	truncedCylinder.Maximum = 2
	truncedCylinder.Closed = true

	tests := []struct {
		name      string
		origin    Tuple
		direction Tuple
		cylinder  *Cylinder
		want      []Intersection
	}{
		{
			name:      "the ray intersects the truncated cylinder cap 1",
			cylinder:  truncedCylinder,
			origin:    Point(0, 3, 0),
			direction: Vector(0, -1, 0),
			want:      []Intersection{NewIntersection(1, truncedCylinder), NewIntersection(2, truncedCylinder)},
		},
		{
			name:      "the ray intersects the truncated cylinder cap 2",
			cylinder:  truncedCylinder,
			origin:    Point(0, 3, -2),
			direction: Vector(0, -1, 2),
			want:      []Intersection{NewIntersection(2.23606797749979, truncedCylinder), NewIntersection(3.3541019662496843, truncedCylinder)},
		},
		{
			name:      "the ray intersects the truncated cylinder cap 3",
			cylinder:  truncedCylinder,
			origin:    Point(0, 4, -2),
			direction: Vector(0, -1, 1),
			want:      []Intersection{NewIntersection(2.8284271247461903, truncedCylinder), NewIntersection(4.242640687119286, truncedCylinder)},
		},
		{
			name:      "the ray intersects the truncated cylinder cap 4",
			cylinder:  truncedCylinder,
			origin:    Point(0, -1, -2),
			direction: Vector(0, 1, 1),
			want:      []Intersection{NewIntersection(2.8284271247461903, truncedCylinder), NewIntersection(4.242640687119286, truncedCylinder)},
		},
	}

	for i, tt := range tests {
		T.Run(fmt.Sprintf("%d: %s", i, tt.name), func(t *testing.T) {

			normed := Normalize(tt.direction)

			ray := NewRay([3]float64{tt.origin.x, tt.origin.y, tt.origin.z}, [3]float64{normed.x, normed.y, normed.z})

			got := tt.cylinder.LocalIntersect(ray)

			if !got.Equal(Intersections{intersections: tt.want}) {
				t.Errorf("\ntest %d: \nwant: %v \ngot: %v \ndo not match", i, tt.want, got)
			}

		})
	}
}
