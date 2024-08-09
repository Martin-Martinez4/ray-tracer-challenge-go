package main

import (
	"fmt"
	"testing"
)

func TestCamera(T *testing.T) {

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
