package shapes

import (
	"testing"

	pm "github.com/Martin-Martinez4/ray-tracer-challenge-go/primitive_math"
)

func TestUVIntersect(t *testing.T) {

	name := "an intersection with a smooth triangle stores u/v"

	sTri := NewSmoothTriangle(pm.Point(0, 1, 0), pm.Point(-1, 0, 0), pm.Point(1, 0, 0), pm.Vector(0, 1, 0), pm.Vector(-1, 0, 0), pm.Vector(1, 0, 0))

	ray := NewRay([3]float64{-0.2, 0.3, -2}, [3]float64{0, 0, 1})

	t.Run(name, func(t *testing.T) {
		got := sTri.LocalIntersect(ray)

		U := (*got)[0].U
		V := (*got)[0].V
		if !pm.AreFloatsEqual(*U, 0.45) || !pm.AreFloatsEqual(*V, 0.25) {
			t.Errorf("got U: %f V: %f, wanted U: 0.45 V: 0.25", *U, *V)

		}
	})

}

func TestSmoothTriNormalAt(t *testing.T) {

	name := "a smooth triangle uses u/v to interpolate the normal"

	sTri := NewSmoothTriangle(pm.Point(0, 1, 0), pm.Point(-1, 0, 0), pm.Point(1, 0, 0), pm.Vector(0, 1, 0), pm.Vector(-1, 0, 0), pm.Vector(1, 0, 0))

	f1 := 0.45
	f2 := 0.25

	t.Run(name, func(t *testing.T) {
		i := NewIntersectionWithUV(1, sTri, &f1, &f2)
		point := pm.Point(0, 0, 0)
		n := sTri.NormalAt(pm.Point(0, 0, 0), &point, &i)
		want := pm.Vector(-0.5547, 0.83205, 0)
		if !n.Equal(want) {
			t.Errorf("\nwanted: %s got: %s\n", want.Print(), n.Print())

		}
	})
}

func TestSmoothTriCompsNormalV(t *testing.T) {

	name := "preparing the normal on a smooth triangle"

	sTri := NewSmoothTriangle(pm.Point(0, 1, 0), pm.Point(-1, 0, 0), pm.Point(1, 0, 0), pm.Vector(0, 1, 0), pm.Vector(-1, 0, 0), pm.Vector(1, 0, 0))

	f1 := 0.45
	f2 := 0.25

	t.Run(name, func(t *testing.T) {
		i := NewIntersectionWithUV(1, sTri, &f1, &f2)
		ray := NewRay([3]float64{-0.2, 0.3, -2}, [3]float64{0, 0, 1})

		xs := Intersections{}
		xs.Add(i)

		comps := PrepareComputationsWithHit(i, ray, xs)

		want := pm.Vector(-0.5547, 0.83205, 0)
		if !comps.Normalv.Equal(want) {
			t.Errorf("\nwanted: %s got: %s\n", want.Print(), comps.Normalv.Print())

		}
	})

}
