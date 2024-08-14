package main

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

func ShadeHit(world *World, comps *Computations, reflectionsLeft int) Color {

	shadowed := IsShadowed(*world, comps.OverPoint)

	reflectedColor := RelfectedColor(world, comps, reflectionsLeft)
	surface := EffectiveLighting(*comps.Object.GetMaterial(), comps.Object, world.Light, comps.OverPoint, comps.Eyev, comps.Normalv, shadowed)

	return surface.Add(reflectedColor)
}

func ColorAt(ray *Ray, world *World, reflectionsLeft int) Color {
	inters := RayWorldIntersect(*ray, *world)

	intersection := inters.Hit()

	if intersection == nil {
		return NewColor(0, 0, 0)
	}

	comps := PrepareComputations(*ray, intersection.S, *intersection)

	return ShadeHit(world, &comps, reflectionsLeft)

}

func Render(camera Camera, world World) Canvas {
	image := NewCanvas(int32(camera.HSize), int32(camera.VSize))

	for y := 0; y < int(camera.VSize)-1; y++ {
		for x := 0; x < int(camera.HSize)-1; x++ {
			ray := RayForPixel(camera, float64(x), float64(y))
			color := ColorAt(&ray, &world, 4)
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

	intersection, hit := Hit(&intersections)
	if hit && intersection.T < distance {
		return true
	} else {
		return false
	}

}
