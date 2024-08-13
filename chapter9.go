package main

import "math"

// Same as ch7 but with shadows
func ch9() string {
	// Floor
	floor := NewPlane()
	floor.SetTransforms([]Matrix4x4{RotationAlongX(math.Pi / 2), Translate(0, 0, 5)})

	// Spheres
	middleSphere := NewSphere()
	middleSphere.Transforms = middleSphere.Transforms.Translate(-0.5, 1, 0.5)
	middleSphere.Material.Color = NewColor(0.1, 1, 0.5)
	middleSphere.Material.Diffuse = 0.7
	middleSphere.Material.Specular = 0.3

	rightSphere := NewSphere()
	rightSphere.Transforms = rightSphere.Transforms.Scale(0.5, 0.5, 0.5)
	rightSphere.Transforms = rightSphere.Transforms.Translate(1.5, 0.5, -0.5)
	rightSphere.Material.Color = NewColor(0.5, 1, 0.1)
	rightSphere.Material.Diffuse = 0.7
	rightSphere.Material.Specular = 0.3

	leftSphere := NewSphere()
	leftSphere.Transforms = leftSphere.Transforms.Scale(0.33, 0.33, 0.33)
	leftSphere.Transforms = leftSphere.Transforms.Translate(-1.5, 0.33, -0.75)
	leftSphere.Material.Color = NewColor(1, 0.8, 0.1)
	leftSphere.Material.Diffuse = 0.7
	leftSphere.Material.Specular = 0.3

	// Light Source
	world := NewDefaultWorld()
	world.Shapes = []Shape{floor, middleSphere, leftSphere, rightSphere}
	world.Light = NewLight([3]float64{-10, 10, -10}, [3]float64{1, 1, 1})

	camera := NewCamera(400, 225, math.Pi/3)
	camera.Transform = ViewTransformation(Point(0, 1.5, -5), Point(0, 1, 0), Vector(0, 1, 0))

	canvas := Render(camera, world)

	return canvas.Newppm()

}
