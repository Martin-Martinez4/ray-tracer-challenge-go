package main

import (
	"math"
)

// Same as ch7 but with shadows
func ch10() string {
	// Floor
	floor := NewPlane()
	floor.SetTransforms([]*Matrix4x4{Translate(0, 0, 5)})
	floorPat := NewRing(WHITE, NewColor(1, 1, 0))
	// floorPat.SetTransform(Scale(10, 10, 10))
	floor.GetMaterial().Pattern = floorPat

	// Spheres
	middleSphere := NewSphere()
	middleSphere.SetTransforms([]*Matrix4x4{Scale(1.1, 1.1, 1.1), Translate(-0.5, 1, -1.1), RotationAlongX(90 * (math.Pi / 180))})
	middleSphere.Material.Color = NewColor(0.1, 1, 0.5)
	middleSphere.Material.Diffuse = 0.7
	middleSphere.Material.Specular = 0.3

	middPat := NewRing(WHITE, NewColor(0, 1, 0))
	middPat.SetTransform(Scale(.025, .025, .025))

	middleSphere.GetMaterial().Pattern = middPat

	rightSphere := NewSphere()
	rightSphere.Transforms = rightSphere.Transforms.Scale(0.5, 0.5, 0.5)
	rightSphere.Transforms = rightSphere.Transforms.Translate(1.5, 0.5, -0.5)
	rightSphere.SetTransform(RotationAlongY(50 * (math.Pi / 180)))
	rightSphere.Material.Color = NewColor(0.5, 1, 0.1)
	rightSphere.Material.Diffuse = 0.7
	rightSphere.Material.Specular = 0.3
	rightPattern := NewStripe(WHITE, BLACK)
	rightPattern.SetTransform(Scale(.07, .07, .07))
	rightSphere.Material.Pattern = rightPattern

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

	camera := NewCamera(800, 400, math.Pi/3)
	camera.Transform = ViewTransformation(Point(0, 1.5, -5), Point(0, 1, 0), Vector(0, 1, 0))

	canvas := Render(camera, world)

	return canvas.Newppm()

}
