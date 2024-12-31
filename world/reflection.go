package world

import (
	mat "github.com/Martin-Martinez4/ray-tracer-challenge-go/materials"
	pm "github.com/Martin-Martinez4/ray-tracer-challenge-go/primitive_math"
	"github.com/Martin-Martinez4/ray-tracer-challenge-go/shapes"
)

func RelfectedColor(world World, comps shapes.Computations, reflectionsLeft int) mat.Color {

	if comps.Object.GetMaterial().Reflective < pm.Epsilon || reflectionsLeft <= 0 {
		return mat.NewColor(0, 0, 0)
	} else {
		reflectRay := shapes.NewRay([3]float64{comps.OverPoint.X, comps.OverPoint.Y, comps.OverPoint.Z}, [3]float64{comps.ReflectV.X, comps.ReflectV.Y, comps.ReflectV.Z})
		color := ColorAt(reflectRay, world, reflectionsLeft-1)

		return color.SMultiply(comps.Object.GetMaterial().Reflective)
	}
}
