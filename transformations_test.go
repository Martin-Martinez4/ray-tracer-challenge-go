package main

import (
	"fmt"
	"testing"
)

func TestViewTransformation(T *testing.T) {

	theWorld := NewDefaultWorld()

	outer := theWorld.Shapes[0]
	outer.GetMaterial().SetAmbient(1)
	theWorld.Shapes[0] = outer

	inner := theWorld.Shapes[1]
	inner.GetMaterial().SetAmbient(1)
	theWorld.Shapes[0] = inner

	tests := []struct {
		name string
		from Tuple
		to   Tuple
		up   Tuple
		want Matrix4x4
	}{
		{
			name: "a view transformation matrix looking in positive z direction",
			from: Point(0, 0, 0),
			to:   Point(0, 0, 1),
			up:   Vector(0, 1, 0),
			want: IdentitiyMatrix4x4().Scale(-1, 1, -1),
		},
		{
			name: "a view transformation moves the world",
			from: Point(0, 0, 8),
			to:   Point(0, 0, 0),
			up:   Vector(0, 1, 0),
			want: IdentitiyMatrix4x4().Translate(0, 0, -8),
		},
		{
			name: "an arbitrary view transformation",
			from: Point(1, 3, 2),
			to:   Point(4, -2, 8),
			up:   Vector(1, 1, 0),
			want: NewMatrix4x4([16]float64{-0.50709, 0.50709, 0.67612, -2.36643, 0.76772, 0.60609, 0.12122, -2.82843, -.35857, .59761, -.71714, 0, 0, 0, 0, 1}),
		},
	}

	for i, tt := range tests {
		T.Run(fmt.Sprintf("%d: %s", i, tt.name), func(t *testing.T) {

			got := ViewTransformation(tt.from, tt.to, tt.up)

			if !got.Equal(tt.want) {
				t.Errorf("%d: \nwant: %v \ngot: %v \ndo not match", i, tt.want, got)
			}

		})
	}
}
