package main

type World struct {
	Spheres []Sphere
	Light   Light
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
func NewWorld(spheres *[]Sphere, light *Light) World {

	if spheres == nil {
		sphere1 := NewSphere()
		sphere1.Material.Color = NewColor(0.8, 1.0, 0.6)
		sphere1.Material.Diffuse = 0.7
		sphere1.Material.Specular = 0.2

		sphere2 := NewSphere()

		// Need to make better
		sphere2.Transforms = sphere2.Transforms.Scale(0.5, 0.5, 0.5)

		spheres = &[]Sphere{sphere1, sphere2}

	}

	if light == nil {
		newLight := NewLight([3]float64{-10, 10, -10}, [3]float64{1, 1, 1})

		light = &newLight
	}

	return World{Spheres: *spheres, Light: *light}
}

func RayWorldIntersect(ray Ray, world World) Intersections {

	inters := Intersections{intersections: []Intersection{}}

	for i := 0; i < len(world.Spheres); i++ {

		inters.RaySphereInteresect(ray, &world.Spheres[i])
	}

	return inters
}

func ShadeHit(world *World, comps *Computations) Color {
	return EffectiveLighting(comps.Object.Material, world.Light, comps.Point, comps.Eyev, comps.Normalv)
}

func ColorAt(ray *Ray, world *World) Color {
	inters := RayWorldIntersect(*ray, *world)

	intersection := inters.Hit()

	if intersection == nil {
		return NewColor(0, 0, 0)
	}

	comps := PrepareComputations(*ray, intersection.S, *intersection)

	return ShadeHit(world, &comps)

}
