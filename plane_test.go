package main

import "testing"

func TestPlaneLocalNormal(t *testing.T) {

	plane := NewPlane()

	tests := []struct {
		name  string
		point Tuple
		want  Tuple
	}{
		{
			name:  "the normal of a plane is constant everywhere",
			point: Point(0, 0, 0),
			want:  Vector(0, 1, 0),
		},
		{
			name:  "the normal of a plane is constant everywhere",
			point: Point(10, 0, -10),
			want:  Vector(0, 1, 0),
		},
		{
			name:  "the normal of a plane is constant everywhere",
			point: Point(-5, 0, 150),
			want:  Vector(0, 1, 0),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := plane.LocalNormalAt(tt.point)

			if !got.Equal(tt.want) {
				t.Errorf("%s did not pass: \nGot: %s \nWanted: %s", tt.name, got.Print(), tt.want.Print())
			}

		})
	}
}

func TestPlaneLIntersect(t *testing.T) {

	plane := NewPlane()

	tests := []struct {
		name string
		ray  Ray
		want Intersections
	}{
		{
			name: "intersect with a ray parallel to the plane",
			ray:  NewRay([3]float64{0, 10, 0}, [3]float64{0, 0, 1}),
			want: Intersections{intersections: nil},
		},
		{
			name: "a ray intersecting a plane from above",
			ray:  NewRay([3]float64{0, 1, 0}, [3]float64{0, -1, 0}),
			want: Intersections{intersections: []Intersection{{T: 1, S: plane}}},
		},
		{
			name: "a ray intersecting a plane from below",
			ray:  NewRay([3]float64{0, -1, 0}, [3]float64{0, 1, 0}),
			want: Intersections{intersections: []Intersection{{T: 1, S: plane}}},
		},
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := plane.LocalIntersect(tt.ray)

			if len(got.intersections) != len(tt.want.intersections) {
				t.Errorf("%d: %s did not pass, not the same length\ngot: %v\nwanted: %v", i, tt.name, got.intersections, tt.want.intersections)

			}

			if !got.Equal(tt.want) {
				t.Errorf("%d: %s did not pass, values did not match", i, tt.name)
			}

		})
	}
}
