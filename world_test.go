package main

import (
	"fmt"
	"testing"
)

func TestIntersectWorld(T *testing.T) {

	theWorld := NewDefaultWorld()

	tests := []struct {
		name string
		ray  Ray
		want Intersections
	}{
		{
			name: "a ray intersecting a default world should return an intersection struct with four members",
			ray:  NewRay([3]float64{0, 0, -5}, [3]float64{0, 0, 1}),
			want: Intersections{
				intersections: []Intersection{
					{S: &theWorld.Spheres[0], T: 4},
					{S: &theWorld.Spheres[0], T: 4.5},
					{S: &theWorld.Spheres[1], T: 5.5},
					{S: &theWorld.Spheres[1], T: 6},
				},
			},
		},
	}

	for i, tt := range tests {
		T.Run(fmt.Sprintf("%d: %s", i, tt.name), func(t *testing.T) {

			got := RayWorldIntersect(tt.ray, theWorld)

			if len(tt.want.intersections) != len(got.intersections) {
				t.Errorf("Lengths did not match for test %d: %s", i, tt.name)
			}

			for k := 0; k < len(got.intersections); k++ {
				if !AreFloatsEqual(tt.want.intersections[k].T, got.intersections[k].T) {
					t.Errorf("%d: \nwant: %v \ngot: %v \ndo not match", k, tt.want.intersections, got.intersections)
				}
			}

		})
	}
}
