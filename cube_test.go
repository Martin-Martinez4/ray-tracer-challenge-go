package main

import (
	"fmt"
	"testing"
)

func TestRayCubeIntersect(T *testing.T) {

	cube := NewCube()

	tests := []struct {
		name string
		ray  Ray
		cube Cube
		want []Intersection
	}{
		{
			name: "the cube intersect +x direction",
			cube: *cube,
			ray:  NewRay([3]float64{5, 0.5, 0}, [3]float64{-1, 0, 0}),
			want: []Intersection{{4, cube}, {6, cube}},
		},
		{
			name: "the cube intersect -x direction",
			cube: *cube,
			ray:  NewRay([3]float64{-5, 0.5, 0}, [3]float64{1, 0, 0}),
			want: []Intersection{{4, cube}, {6, cube}},
		},
		{
			name: "the cube intersect +y direction",
			cube: *cube,
			ray:  NewRay([3]float64{0.5, 5, 0}, [3]float64{0, -1, 0}),
			want: []Intersection{{4, cube}, {6, cube}},
		},
		{
			name: "the cube intersect -y direction",
			cube: *cube,
			ray:  NewRay([3]float64{0.5, -5, 0}, [3]float64{0, 1, 0}),
			want: []Intersection{{4, cube}, {6, cube}},
		},
		{
			name: "the cube intersect +z direction",
			cube: *cube,
			ray:  NewRay([3]float64{0.5, 0, 5}, [3]float64{0, 0, -1}),
			want: []Intersection{{4, cube}, {6, cube}},
		},
		{
			name: "the cube intersect -z direction",
			cube: *cube,
			ray:  NewRay([3]float64{0.5, 0, -5}, [3]float64{0, 0, 1}),
			want: []Intersection{{4, cube}, {6, cube}},
		},
		{
			name: "the cube intersect inside",
			cube: *cube,
			ray:  NewRay([3]float64{0, 0.5, 0}, [3]float64{0, 0, 1}),
			want: []Intersection{{-1, cube}, {1, cube}},
		},
		{
			name: "a ray misses a cube 1",
			cube: *cube,
			ray:  NewRay([3]float64{-2, 0, 0}, [3]float64{0.2673, 0.5345, 0.8018}),
			want: []Intersection{},
		},
		{
			name: "a ray misses a cube 2",
			cube: *cube,
			ray:  NewRay([3]float64{0, -2, 0}, [3]float64{0.8018, 0.2673, 0.5345}),
			want: []Intersection{},
		},
		{
			name: "a ray misses a cube 3",
			cube: *cube,
			ray:  NewRay([3]float64{0, 0, -2}, [3]float64{0.5345, 0.8018, 0.2673}),
			want: []Intersection{},
		},
		{
			name: "a ray misses a cube 4",
			cube: *cube,
			ray:  NewRay([3]float64{2, 2, 0}, [3]float64{-1, 0, 0}),
			want: []Intersection{},
		},
	}

	for i, tt := range tests {
		T.Run(fmt.Sprintf("%d: %s", i, tt.name), func(t *testing.T) {

			got := cube.LocalIntersect(tt.ray)

			if !got.Equal(Intersections{intersections: tt.want}) {
				t.Errorf("%d: \nwant: %v \ngot: %v \ndo not match", i, tt.want, got)
			}

		})
	}
}

func TestRayCubeNormal(T *testing.T) {

	cube := NewCube()

	tests := []struct {
		name  string
		cube  Cube
		point Tuple
		want  Tuple
	}{
		{
			name:  "the cube normal",
			cube:  *cube,
			point: Point(1, 0.5, -0.8),
			want:  Vector(1, 0, 0),
		},
		{
			name:  "the cube normal",
			cube:  *cube,
			point: Point(-1, -0.2, 0.9),
			want:  Vector(-1, 0, 0),
		},
		{
			name:  "the cube normal",
			cube:  *cube,
			point: Point(-0.4, 1, -0.1),
			want:  Vector(0, 1, 0),
		},
		{
			name:  "the cube normal",
			cube:  *cube,
			point: Point(-0.6, 0.3, 1),
			want:  Vector(0, 0, 1),
		},
		{
			name:  "the cube normal",
			cube:  *cube,
			point: Point(1, 1, 1),
			want:  Vector(1, 0, 0),
		},
		{
			name:  "the cube normal",
			cube:  *cube,
			point: Point(-1, -1, -1),
			want:  Vector(-1, 0, 0),
		},
	}

	for i, tt := range tests {
		T.Run(fmt.Sprintf("%d: %s", i, tt.name), func(t *testing.T) {

			got := cube.LocalNormalAt(tt.point)

			if !got.Equal(tt.want) {
				t.Errorf("%d: \nwant: %s \ngot: %s \ndo not match", i, tt.want.Print(), got.Print())
			}

		})
	}
}
