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

func TestShadeHit(T *testing.T) {

	theWorld := NewDefaultWorld()
	otherLight := (NewLight([3]float64{0, 0.25, 0}, [3]float64{1, 1, 1}))
	theOtherWorld := NewWorld(nil, &otherLight)

	tests := []struct {
		name         string
		ray          Ray
		sphere       Sphere
		world        World
		intersection Intersection
		want         Color
	}{
		{
			name:         "shading an intersection",
			ray:          NewRay([3]float64{0, 0, -5}, [3]float64{0, 0, 1}),
			sphere:       theWorld.Spheres[0],
			world:        theWorld,
			intersection: Intersection{4, &theWorld.Spheres[0]},
			want:         NewColor(0.38066, 0.47583, 0.2855),
		},
		{
			name:         "shading an intersection from the inside",
			ray:          NewRay([3]float64{0, 0, 0}, [3]float64{0, 0, 1}),
			sphere:       theOtherWorld.Spheres[1],
			world:        theOtherWorld,
			intersection: Intersection{0.5, &theOtherWorld.Spheres[1]},
			want:         NewColor(0.90498, 0.90498, 0.90498),
		},
	}

	for i, tt := range tests {
		T.Run(fmt.Sprintf("%d: %s", i, tt.name), func(t *testing.T) {

			comps := PrepareComputations(tt.ray, &tt.sphere, tt.intersection)

			got := ShadeHit(&tt.world, &comps)

			if !got.Equal(tt.want) {
				t.Errorf("%d: \nwant: %v \ngot: %v \ndo not match", i, tt.want, got)
			}

		})
	}
}

func TestColorAt(T *testing.T) {

	theWorld := NewDefaultWorld()

	tests := []struct {
		name  string
		ray   Ray
		world World
		want  Color
	}{
		{
			name:  "the color when a ray misses should be black (0,0,0)",
			ray:   NewRay([3]float64{0, 0, -5}, [3]float64{0, 1, 0}),
			world: theWorld,
			want:  NewColor(0, 0, 0),
		},
		{
			name:  "the color when a ray hits",
			ray:   NewRay([3]float64{0, 0, -5}, [3]float64{0, 0, 1}),
			world: theWorld,
			want:  NewColor(0.38066, 0.47583, 0.2855),
		},
	}

	for i, tt := range tests {
		T.Run(fmt.Sprintf("%d: %s", i, tt.name), func(t *testing.T) {

			got := ColorAt(&tt.ray, &tt.world)

			if !got.Equal(tt.want) {
				t.Errorf("%d: \nwant: %v \ngot: %v \ndo not match", i, tt.want, got)
			}

		})
	}
}

func TestColorAtInner(T *testing.T) {

	theWorld := NewDefaultWorld()

	outer := theWorld.Spheres[0]
	outer.Material.Ambient = 1
	theWorld.Spheres[0] = outer

	inner := theWorld.Spheres[1]
	inner.Material.Ambient = 1
	theWorld.Spheres[0] = inner

	tests := []struct {
		name  string
		ray   Ray
		world World
		want  Color
	}{
		{
			name:  "the color with an intersection behind the ray",
			ray:   NewRay([3]float64{0, 0, 0.75}, [3]float64{0, 0, -1}),
			world: theWorld,
			want:  inner.Material.Color,
		},
	}

	for i, tt := range tests {
		T.Run(fmt.Sprintf("%d: %s", i, tt.name), func(t *testing.T) {

			got := ColorAt(&tt.ray, &tt.world)

			if !got.Equal(tt.want) {
				t.Errorf("%d: \nwant: %v \ngot: %v \ndo not match", i, tt.want, got)
			}

		})
	}
}
