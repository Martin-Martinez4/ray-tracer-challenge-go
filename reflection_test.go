package main

import (
	"math"
	"testing"
)

func TestReflectionV(t *testing.T) {

	tests := []struct {
		name         string
		shape        Shape
		ray          Ray
		intersection Intersection
		want         Tuple
	}{
		{
			name:         "precomputing the reflection vector",
			shape:        NewPlane(),
			ray:          NewRay([3]float64{0, 1, -1}, [3]float64{0, -math.Sqrt(2) / 2, math.Sqrt(2) / 2}),
			intersection: Intersection{math.Sqrt(2), nil},
			want:         Vector(0, math.Sqrt(2)/2, math.Sqrt(2)/2),
		},
	}

	for i, tt := range tests {

		tt.intersection.S = tt.shape

		t.Run(tt.name, func(t *testing.T) {
			got := PrepareComputations(tt.ray, tt.shape, tt.intersection).ReflectV

			if !got.Equal(tt.want) {
				t.Errorf("%d failed\nwanted:\n%s\ngot:\n%s", i, tt.want.Print(), got.Print())
			}
		})
	}

}

func TestNonreflective(t *testing.T) {

	tests := []struct {
		name         string
		world        World
		ray          Ray
		intersection Intersection
		want         Color
	}{
		{
			name:         "the reflected color for a nonreflective material shoul be black",
			world:        NewDefaultWorld(),
			ray:          NewRay([3]float64{0, 1, -1}, [3]float64{0, -math.Sqrt(2) / 2, math.Sqrt(2) / 2}),
			intersection: Intersection{1, nil},
			want:         NewColor(0, 0, 0),
		},
	}

	for i, tt := range tests {

		shape := tt.world.Shapes[1]
		shape.GetMaterial().Ambient = 1

		t.Run(tt.name, func(t *testing.T) {
			preComp := PrepareComputations(tt.ray, shape, tt.intersection)
			got := RelfectedColor(&tt.world, &preComp, 1)

			if !got.Equal(tt.want) {
				t.Errorf("%d failed\nwanted:\n%s\ngot:\n%s", i, tt.want.Print(), got.Print())
			}
		})
	}

}

func TestReflective(t *testing.T) {

	tests := []struct {
		name         string
		world        World
		shape        Shape
		ray          Ray
		intersection Intersection
		want         Color
	}{
		{
			name:         "the reflected color for a relfective material",
			world:        NewDefaultWorld(),
			shape:        NewPlane(),
			ray:          NewRay([3]float64{0, 0, -3}, [3]float64{0, -math.Sqrt(2) / 2, math.Sqrt(2) / 2}),
			intersection: Intersection{math.Sqrt(2), nil},
			want:         NewColor(0.19032, 0.2379, 0.14277),
		},
	}

	for i, tt := range tests {

		tt.shape.SetTransform(Translate(0, -1, 0))
		tt.shape.GetMaterial().Reflective = 0.5

		tt.intersection.S = tt.shape

		t.Run(tt.name, func(t *testing.T) {
			preComp := PrepareComputations(tt.ray, tt.shape, tt.intersection)
			got := RelfectedColor(&tt.world, &preComp, 1)

			if !got.Equal(tt.want) {
				t.Errorf("%d failed\nwanted:\n%s\ngot:\n%s", i, tt.want.Print(), got.Print())
			}
		})
	}

}

func TestReflectiveShadeHit(t *testing.T) {

	tests := []struct {
		name         string
		world        World
		shape        Shape
		ray          Ray
		intersection Intersection
		want         Color
	}{
		{
			name:         "the reflected color for a relfective material",
			world:        NewDefaultWorld(),
			shape:        NewPlane(),
			ray:          NewRay([3]float64{0, 0, -3}, [3]float64{0, -math.Sqrt(2) / 2, math.Sqrt(2) / 2}),
			intersection: Intersection{math.Sqrt(2), nil},
			want:         NewColor(0.87677, 0.92436, 0.82918),
		},
	}

	for i, tt := range tests {

		tt.shape.SetTransform(Translate(0, -1, 0))
		tt.shape.GetMaterial().Reflective = 0.5

		tt.intersection.S = tt.shape

		t.Run(tt.name, func(t *testing.T) {
			preComp := PrepareComputations(tt.ray, tt.shape, tt.intersection)
			got := ShadeHit(&tt.world, &preComp, 1)

			if !got.Equal(tt.want) {
				t.Errorf("%d failed\nwanted:\n%s\ngot:\n%s", i, tt.want.Print(), got.Print())
			}
		})
	}
}

func TestReflectiveZeroAllowed(t *testing.T) {

	tests := []struct {
		name         string
		world        World
		shape        Shape
		ray          Ray
		intersection Intersection
		want         Color
	}{
		{
			name:         "the reflected color for a relfective material",
			world:        NewDefaultWorld(),
			shape:        NewPlane(),
			ray:          NewRay([3]float64{0, 0, -3}, [3]float64{0, -math.Sqrt(2) / 2, math.Sqrt(2) / 2}),
			intersection: Intersection{math.Sqrt(2), nil},
			want:         BLACK,
		},
	}

	for i, tt := range tests {

		tt.shape.SetTransform(Translate(0, -1, 0))
		tt.shape.GetMaterial().Reflective = 0.5

		tt.intersection.S = tt.shape

		t.Run(tt.name, func(t *testing.T) {
			preComp := PrepareComputations(tt.ray, tt.shape, tt.intersection)
			got := RelfectedColor(&tt.world, &preComp, 0)

			if !got.Equal(tt.want) {
				t.Errorf("%d failed\nwanted:\n%s\ngot:\n%s", i, tt.want.Print(), got.Print())
			}
		})
	}

}

func TestTwoRefelctiveShapes(t *testing.T) {
	lowerPanel := NewPlane()
	lowerPanel.GetMaterial().Reflective = 1
	lowerPanel.SetTransform(Translate(0, -1, 0))

	upperPanel := NewPlane()
	upperPanel.GetMaterial().Reflective = 1
	upperPanel.SetTransform(Translate(0, 1, 0))

	light := NewLight([3]float64{0, 0, 0}, [3]float64{1, 1, 1})

	world := NewWorld(&[]Shape{lowerPanel, upperPanel}, &light)

	ray := NewRay([3]float64{0, 0, 0}, [3]float64{0, 1, 0})

	ColorAt(&ray, &world, 4)

}
