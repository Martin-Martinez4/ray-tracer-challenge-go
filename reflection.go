package main

func RelfectedColor(world World, comps Computations, reflectionsLeft int) Color {

	if comps.Object.GetMaterial().Reflective < Epsilon || reflectionsLeft <= 0 {
		return NewColor(0, 0, 0)
	} else {
		reflectRay := NewRay([3]float64{comps.OverPoint.x, comps.OverPoint.y, comps.OverPoint.z}, [3]float64{comps.ReflectV.x, comps.ReflectV.y, comps.ReflectV.z})
		color := ColorAt(reflectRay, world, reflectionsLeft-1)

		return color.SMultiply(comps.Object.GetMaterial().Reflective)
	}
}
