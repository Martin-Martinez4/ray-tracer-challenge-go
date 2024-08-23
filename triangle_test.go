package main

import "testing"

func TestNormalOnATriangle(t *testing.T) {
	tri := NewTriangle(Point(0, 1, 0), Point(-1, 0, 0), Point(1, 0, 0))

	n1 := tri.LocalNormalAt(Point(0, 0.5, 0))
	n2 := tri.LocalNormalAt(Point(-0.5, 0.75, 0))
	n3 := tri.LocalNormalAt(Point(0.5, 0.25, 0))

	if !(n1.Equal(n2) && n2.Equal(n3) && n2.Equal(tri.Normal)) {
		t.Errorf("Normal on triangle test failed,\nwanted: %s\ngot: \nn1: %s\nn2: %s\nn3: %s\n", tri.Normal.Print(), n1.Print(), n2.Print(), n3.Print())
	}

}

func TestTriangleIntersection(t *testing.T) {

	triangle := NewTriangle(Point(0, 1, 0), Point(-1, 0, 0), Point(1, 0, 0))

	tests := []struct {
		name string
		ray  Ray
		want Intersections
	}{
		{
			name: "a ray misses p1-p3 edge",
			ray:  NewRay([3]float64{-1, 1, -2}, [3]float64{0, 0, 1}),
			want: Intersections{intersections: []Intersection{}},
		},
		{
			name: "a ray misses p1-p2 edge",
			ray:  NewRay([3]float64{1, 1, -2}, [3]float64{0, 0, 1}),
			want: Intersections{intersections: []Intersection{}},
		},
		{
			name: "a ray misses p2-p3 edge",
			ray:  NewRay([3]float64{0, -1, -2}, [3]float64{0, 0, 1}),
			want: Intersections{intersections: []Intersection{}},
		},
		{
			name: "a ray strikes a triangle",
			ray:  NewRay([3]float64{0, 0.5, -2}, [3]float64{0, 0, 1}),
			want: Intersections{intersections: []Intersection{{2, triangle}}},
		},
	}

	for i, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			got := triangle.LocalIntersect(tt.ray)

			if !got.Equal(tt.want) {
				t.Errorf("\n%d %s failed:\nwanted: %v\ngot: %v\n", i, tt.name, tt.want, got)
			}
		})

	}

}
