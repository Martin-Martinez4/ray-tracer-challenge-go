package main

import (
	"math"
)

// Same as ch7 but with shadows
func ch11A() string {
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
	middleSphere.Material.Reflective = 0.5

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

func ch11B() string {
	// Floor

	floor := NewPlane()
	floor.SetTransform(Translate(0, -1, 0))
	floor.GetMaterial().Transparency = 0.5
	floor.GetMaterial().RefractiveIndex = 1

	ball := NewSphere()
	ball.GetMaterial().Ambient = 0.5
	ball.GetMaterial().Transparency = 0.1

	glassSphere := NewGlassSphere()

	// Light Source
	world := NewDefaultWorld()
	world.Shapes = []Shape{floor, glassSphere}
	world.Light = NewLight([3]float64{-10, 10, -10}, [3]float64{1, 1, 1})

	camera := NewCamera(800, 400, math.Pi/3)
	camera.Transform = ViewTransformation(Point(0, 1.5, -5), Point(0, 1, 0), Vector(0, 1, 0))

	canvas := Render(camera, world)

	return canvas.Newppm()

}

func ch11C() string {
	camera := NewCamera(300, 300, 0.45)
	from := Point(0, 0, -5)
	to := Point(0, 0, 0)
	up := Point(0, 1, 0)
	camera.Transform = ViewTransformation(from, to, up)

	light := NewLight([3]float64{2, 10, -5}, [3]float64{0.9, 0.9, 0.9})

	// # wall
	plane := NewPlane()
	plane.SetTransforms([]*Matrix4x4{RotationAlongX(1.5708), Translate(0, 0, 10)})
	plane.GetMaterial().Pattern = NewChecker(NewColor(0.15, 0.15, 0.15), NewColor(0.85, 0.85, 0.85))
	plane.GetMaterial().Ambient = 0.8
	plane.GetMaterial().Diffuse = 0.2
	plane.GetMaterial().Specular = 0

	// # glass ball
	gSphere := NewSphere()
	// Should make a way to add a material instead of setting things on at a time
	gSphere.GetMaterial().Color = NewColor(1, 1, 1)
	gSphere.GetMaterial().Ambient = 0
	gSphere.GetMaterial().Diffuse = 0
	gSphere.GetMaterial().Specular = 0.9
	gSphere.GetMaterial().Shininess = 300
	gSphere.GetMaterial().Reflective = 0.9
	gSphere.GetMaterial().Transparency = 0.9
	gSphere.GetMaterial().RefractiveIndex = 1.5

	// # hollow center
	hollow := NewSphere()
	hollow.SetTransform(Scale(0.5, 0.5, 0.5))
	// Should make a way to add a material instead of setting things on at a time
	gSphere.GetMaterial().Color = NewColor(1, 1, 1)
	gSphere.GetMaterial().Ambient = 0
	gSphere.GetMaterial().Diffuse = 0
	gSphere.GetMaterial().Specular = 0.9
	gSphere.GetMaterial().Shininess = 300
	gSphere.GetMaterial().Reflective = 0.9
	gSphere.GetMaterial().Transparency = 0.9
	gSphere.GetMaterial().RefractiveIndex = 1.0000034

	world := NewWorld(&[]Shape{gSphere, plane, hollow}, &light)

	canvas := Render(camera, world)

	return canvas.Newppm()

}
