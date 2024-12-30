package world

import (
	"math"
	"testing"

	pm "github.com/Martin-Martinez4/ray-tracer-challenge-go/primitive_math"
	"github.com/Martin-Martinez4/ray-tracer-challenge-go/shapes"
)

func TestReflectionV(t *testing.T) {

	tests := []struct {
		name         string
		shape        shapes.Shape
		ray          shapes.Ray
		intersection shapes.Intersection
		want         pm.Tuple
	}{
		{
			name:         "precomputing the reflection vector",
			shape:        shapes.NewPlane(),
			ray:          shapes.NewRay([3]float64{0, 1, -1}, [3]float64{0, -math.Sqrt(2) / 2, math.Sqrt(2) / 2}),
			intersection: shapes.NewIntersection(math.Sqrt(2), nil),
			want:         pm.Vector(0, math.Sqrt(2)/2, math.Sqrt(2)/2),
		},
	}

	for i, tt := range tests {

		tt.intersection.S = tt.shape

		t.Run(tt.name, func(t *testing.T) {
			got := shapes.PrepareComputations(tt.ray, tt.shape, tt.intersection).ReflectV

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
		ray          shapes.Ray
		intersection shapes.Intersection
		want         Color
	}{
		{
			name:         "the reflected color for a nonreflective material shoul be black",
			world:        NewDefaultWorld(),
			ray:          shapes.NewRay([3]float64{0, 1, -1}, [3]float64{0, -math.Sqrt(2) / 2, math.Sqrt(2) / 2}),
			intersection: shapes.NewIntersection(1, nil),
			want:         NewColor(0, 0, 0),
		},
	}

	for i, tt := range tests {

		shape := tt.world.Shapes[1]
		shape.GetMaterial().Ambient = 1

		t.Run(tt.name, func(t *testing.T) {
			preComp := shapes.PrepareComputations(tt.ray, shape, tt.intersection)
			got := RelfectedColor(tt.world, preComp, 1)

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
		shape        shapes.Shape
		ray          shapes.Ray
		intersection shapes.Intersection
		want         Color
	}{
		{
			name:         "the reflected color for a relfective material",
			world:        NewDefaultWorld(),
			shape:        shapes.NewPlane(),
			ray:          shapes.NewRay([3]float64{0, 0, -3}, [3]float64{0, -math.Sqrt(2) / 2, math.Sqrt(2) / 2}),
			intersection: shapes.NewIntersection(math.Sqrt(2), nil),
			want:         NewColor(0.19032, 0.2379, 0.14277),
		},
	}

	for i, tt := range tests {

		tt.shape.SetTransform(pm.Translate(0, -1, 0))
		tt.shape.GetMaterial().Reflective = 0.5

		tt.intersection.S = tt.shape

		t.Run(tt.name, func(t *testing.T) {
			preComp := shapes.PrepareComputations(tt.ray, tt.shape, tt.intersection)
			got := RelfectedColor(tt.world, preComp, 1)

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
		shape        shapes.Shape
		ray          shapes.Ray
		intersection shapes.Intersection
		want         Color
	}{
		{
			name:         "the reflected color for a relfective material",
			world:        NewDefaultWorld(),
			shape:        shapes.NewPlane(),
			ray:          shapes.NewRay([3]float64{0, 0, -3}, [3]float64{0, -math.Sqrt(2) / 2, math.Sqrt(2) / 2}),
			intersection: shapes.NewIntersection(math.Sqrt(2), nil),
			want:         NewColor(0.87677, 0.92436, 0.82918),
		},
	}

	for i, tt := range tests {

		tt.shape.SetTransform(pm.Translate(0, -1, 0))
		tt.shape.GetMaterial().Reflective = 0.5

		tt.intersection.S = tt.shape

		t.Run(tt.name, func(t *testing.T) {
			preComp := shapes.PrepareComputations(tt.ray, tt.shape, tt.intersection)
			got := ShadeHit(tt.world, preComp, 1)

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
		shape        shapes.Shape
		ray          shapes.Ray
		intersection shapes.Intersection
		want         Color
	}{
		{
			name:         "the reflected color for a relfective material",
			world:        NewDefaultWorld(),
			shape:        shapes.NewPlane(),
			ray:          shapes.NewRay([3]float64{0, 0, -3}, [3]float64{0, -math.Sqrt(2) / 2, math.Sqrt(2) / 2}),
			intersection: shapes.NewIntersection(math.Sqrt(2), nil),
			want:         BLACK,
		},
	}

	for i, tt := range tests {

		tt.shape.SetTransform(pm.Translate(0, -1, 0))
		tt.shape.GetMaterial().Reflective = 0.5

		tt.intersection.S = tt.shape

		t.Run(tt.name, func(t *testing.T) {
			preComp := shapes.PrepareComputations(tt.ray, tt.shape, tt.intersection)
			got := RelfectedColor(tt.world, preComp, 0)

			if !got.Equal(tt.want) {
				t.Errorf("%d failed\nwanted:\n%s\ngot:\n%s", i, tt.want.Print(), got.Print())
			}
		})
	}

}

func TestTwoRefelctiveShapes(t *testing.T) {
	lowerPanel := shapes.NewPlane()
	lowerPanel.GetMaterial().Reflective = 1
	lowerPanel.SetTransform(pm.Translate(0, -1, 0))

	upperPanel := shapes.NewPlane()
	upperPanel.GetMaterial().Reflective = 1
	upperPanel.SetTransform(pm.Translate(0, 1, 0))

	light := NewLight([3]float64{0, 0, 0}, [3]float64{1, 1, 1})

	world := NewWorld(&[]shapes.Shape{lowerPanel, upperPanel}, &light)

	ray := shapes.NewRay([3]float64{0, 0, 0}, [3]float64{0, 1, 0})

	ColorAt(ray, world, 4)

}
