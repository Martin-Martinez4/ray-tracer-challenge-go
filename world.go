package main

import (
	"math"
)

type World struct {
	Shapes []Shape
	Light  Light
}

// Returns a Default world with:
//
// with a default sphere and a sphere scaled by 0.5
//
// and a Light with position : -10, 10,-10 and intensity: 1,1,1
func NewDefaultWorld() World {
	return NewWorld(nil, nil)
}

/*
[]Spheres
Light
*/
func NewWorld(shapes *[]Shape, light *Light) World {

	if shapes == nil {
		sphere1 := NewSphere()
		sphere1.Material.Color = NewColor(0.8, 1.0, 0.6)
		sphere1.Material.Diffuse = 0.7
		sphere1.Material.Specular = 0.2

		sphere2 := NewSphere()

		// Need to make better
		sphere2.Transforms = sphere2.Transforms.Scale(0.5, 0.5, 0.5)

		shapes = &[]Shape{sphere1, sphere2}

	}

	if light == nil {
		newLight := NewLight([3]float64{-10, 10, -10}, [3]float64{1, 1, 1})

		light = &newLight
	}

	return World{Shapes: *shapes, Light: *light}
}

func RayWorldIntersect(ray Ray, world World) Intersections {

	inters := Intersections{intersections: []Intersection{}}

	for i := 0; i < len(world.Shapes); i++ {

		inters.RayShapeInteresect(ray, world.Shapes[i])
	}

	return inters
}

func ShadeHit(world World, comps Computations, reflectionsLeft int) Color {

	shadowed := IsShadowed(world, comps.OverPoint)

	surface := EffectiveLighting(*comps.Object.GetMaterial(), comps.Object, world.Light, comps.OverPoint, comps.Eyev, comps.Normalv, shadowed)
	reflectedColor := RelfectedColor(world, comps, reflectionsLeft)
	refractedColor := RefreactedColor(world, comps, reflectionsLeft)

	material := comps.Object.GetMaterial()
	if material.Reflective > 0 && material.Transparency > 0 {
		reflectance := Schlick(&comps)

		reflec := reflectedColor.SMultiply(reflectance)
		refrac := refractedColor.SMultiply(1 - reflectance)
		return surface.Add(reflec).Add(refrac)
	}

	return surface.Add(reflectedColor).Add(refractedColor)
}

func ColorAt(ray Ray, world World, reflectionsLeft int) Color {
	inters := RayWorldIntersect(ray, world)

	// intersection := inters.Hit()
	intersection, hit := Hit(inters.intersections)

	if !hit {
		return NewColor(0, 0, 0)
	}

	// comps := PrepareComputations(*ray, intersection.S, *intersection)
	comps := PrepareComputationsWithHit(intersection, ray, inters.intersections)

	return ShadeHit(world, *comps, reflectionsLeft)

}

func RefreactedColor(world World, comps Computations, reflectionsLeft int) Color {
	if comps.Object.GetMaterial().Transparency == 0 || reflectionsLeft <= 0 {
		return BLACK
	}

	nRatio := comps.N1 / comps.N2
	cosI := Dot(comps.Eyev, comps.Normalv)
	sin2T := (nRatio * nRatio) * (1 - (cosI * cosI))

	if sin2T > 1 {
		return BLACK
	}

	cosT := math.Sqrt(1.0 - sin2T)

	d1 := comps.Normalv.SMultiply(nRatio*cosI - cosT)
	d2 := comps.Eyev.SMultiply(nRatio)
	direction := d1.Subtract(d2)

	// fmt.Printf("\nLeft: %d\n", reflectionsLeft)
	// fmt.Printf("\nn1: %f\n", comps.N1)
	// fmt.Printf("n2: %f\n", comps.N2)
	// fmt.Printf("EyeV: %s\n", comps.Eyev.Print())
	// fmt.Printf("NormalV: %s\n", comps.Normalv.Print())
	// fmt.Printf("CosI: %f\n", cosI)
	// fmt.Printf("Sin2T: %f\n", sin2T)
	// fmt.Printf("CosT: %f\n", cosI)
	// fmt.Printf("refract direction: %s\n", direction.Print())

	refractRay := NewRay([3]float64{comps.UnderPoint.x, comps.UnderPoint.y, comps.UnderPoint.z}, [3]float64{direction.x, direction.y, direction.z})

	return ColorAt(refractRay, world, reflectionsLeft-1).SMultiply(comps.Object.GetMaterial().Transparency)
}

func Render(camera Camera, world World) Canvas {
	image := NewCanvas(int32(camera.HSize), int32(camera.VSize))

	for y := 0; y < int(camera.VSize)-1; y++ {
		for x := 0; x < int(camera.HSize)-1; x++ {
			ray := RayForPixel(camera, float64(x), float64(y))
			color := ColorAt(ray, world, 4)
			image.ColorPixel(int32(x), int32(y), color)
		}
	}

	return image
}

func IsShadowed(world World, point Tuple) bool {
	v := world.Light.Position.Subtract(point)
	distance := v.Magnitude()
	direction := Normalize(v)

	ray := NewRay([3]float64{point.x, point.y, point.z}, [3]float64{direction.x, direction.y, direction.z})

	intersections := RayWorldIntersect(ray, world)

	intersection, hit := Hit(intersections.intersections)
	if hit && intersection.T < distance {
		return true
	} else {
		return false
	}

}

func Schlick(comps *Computations) float64 {
	cos := Dot(comps.Eyev, comps.Normalv)

	if comps.N1 > comps.N2 {
		n := comps.N1 / comps.N2
		sin2t := (n * n) * (1 - (cos * cos))
		if sin2t > 1 {
			return 1
		}

		cosT := math.Sqrt(1 - sin2t)

		cos = cosT
	}

	r0 := math.Pow(((comps.N1 - comps.N2) / (comps.N1 + comps.N2)), 2)
	return r0 + (1-r0)*math.Pow((1-cos), 5)
}
