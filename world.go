package main

type World struct {
	Spheres []Sphere
	Light   Light
}

func NewDefaultWorld() World {
	sphere1 := NewSphere()
	sphere1.Material.Color = NewColor(0.8, 1.0, 0.6)
	sphere1.Material.Diffuse = 0.7
	sphere1.Material.Specular = 0.2

	sphere2 := NewSphere()
	sphere2.Transforms.Scale(0.5, 0.5, 0.5)

	light := NewLight([3]float64{-10, 10, -10}, [3]float64{1, 1, 1})

	return World{Spheres: []Sphere{sphere1, sphere2}, Light: light}

}
