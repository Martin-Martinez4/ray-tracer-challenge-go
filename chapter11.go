package main

import (
	"fmt"
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
	floor.GetMaterial().Color = NewColor(0.2, 0.2, 0.2)

	ball := NewSphere()
	ball.GetMaterial().Ambient = 0
	ball.GetMaterial().Specular = 0.9
	ball.GetMaterial().Transparency = 0.9
	ball.GetMaterial().Reflective = 0.9
	ball.GetMaterial().RefractiveIndex = 1.5
	ball.GetMaterial().Color = NewColor(1, 1, 1)

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
	floor := NewPlane()
	floor.SetTransform(Translate(0, -10, 0))
	floor.Material.Pattern = NewChecker(BLACK, WHITE)
	floor.Material.Pattern.SetTransform(Translate(0, 0, 0))
	floor.Material.Specular = 0

	bigger := NewSphere()
	bigger.Material.Diffuse = 0.1
	bigger.Material.Shininess = 300
	bigger.Material.Reflective = 1
	bigger.Material.Transparency = 1
	bigger.Material.RefractiveIndex = 1.52
	bigger.Material.Color = NewColor(0, 0, 0.1)

	smaller := NewSphere()
	smaller.SetTransform(Scale(0.5, 0.5, 0.5))
	smaller.Material.Diffuse = 0.1
	smaller.Material.Shininess = 300
	smaller.Material.Reflective = 1
	smaller.Material.Transparency = 1
	smaller.Material.RefractiveIndex = 1
	smaller.Material.Color = NewColor(0, 0, 0.1)

	light := NewLight([3]float64{2, 10, -5}, [3]float64{0.9, 0.9, 0.9})
	world := NewWorld(&[]Shape{floor, bigger, smaller}, &light)

	camera := NewCamera(512, 512, math.Pi/3)
	from := Point(0, 2.5, 0)
	to := Point(0, 0, 0)
	up := Vector(1, 0, 0)
	camera.Transform = ViewTransformation(from, to, up)

	canvas := Render(camera, world)

	return canvas.Newppm()
}

func ch11D() string {
	floor := NewPlane()
	floor.SetTransforms([]*Matrix4x4{RotationAlongX(1.5708), Translate(0, 0, 10)})
	floor.Material.Pattern = NewChecker(BLACK, WHITE)
	floor.Material.Specular = 0
	floor.Material.Ambient = 0.8
	floor.Material.Diffuse = 0.2

	bigger := NewSphere()
	bigger.Material.Diffuse = 0
	bigger.Material.Ambient = 0
	bigger.Material.Shininess = 300
	bigger.Material.Reflective = .9
	bigger.Material.Transparency = .9
	bigger.Material.Specular = .9
	bigger.Material.RefractiveIndex = 1.52
	bigger.Material.Color = NewColor(1, 1, 1)

	smaller := NewSphere()
	smaller.SetTransform(Scale(0.5, 0.5, 0.5))
	smaller.Material.Diffuse = 0
	smaller.Material.Ambient = 0
	smaller.Material.Specular = 0
	smaller.Material.Shininess = 300
	smaller.Material.Reflective = .9
	smaller.Material.Transparency = .9
	smaller.Material.RefractiveIndex = 1.000034
	smaller.Material.Color = NewColor(1, 1, 1)

	light := NewLight([3]float64{2, 10, -5}, [3]float64{0.9, 0.9, 0.9})
	world := NewWorld(&[]Shape{floor, bigger, smaller}, &light)

	camera := NewCamera(300, 300, 0.45)
	from := Point(0, 0, -5)
	to := Point(0, 0, 0)
	up := Vector(0, 1, 0)
	camera.Transform = ViewTransformation(from, to, up)

	canvas := Render(camera, world)

	fmt.Printf("\n%f\n", camera.HalfHeight)

	return canvas.Newppm()
}

func ch11D3() string {
	theWorld := NewDefaultWorld()

	floor := NewPlane()
	floor.GetMaterial().Transparency = 0.5
	floor.GetMaterial().RefractiveIndex = 1.5
	floor.SetTransform(Translate(0, -1, 0))

	ball := NewSphere()
	ball.GetMaterial().Color = NewColor(1, 0, 0)
	ball.GetMaterial().Ambient = 0.5
	ball.SetTransform(Translate(0, -2.5, -0.5))

	theWorld.Shapes = append(theWorld.Shapes, floor, ball)

	// light := NewLight([3]float64{2, 10, -5}, [3]float64{0.9, 0.9, 0.9})

	// theWorld.Light = light

	// camera := NewCamera(500, 500, 0.45)
	// from := Point(0, 0, -5)
	// to := Point(0, 0, 0)
	// up := Vector(0, 1, 0)
	camera := NewCamera(400, 400, math.Pi/3)
	camera.Transform = ViewTransformation(Point(0, 1.5, -5), Point(0, 1, 0), Vector(0, 1, 0))

	canvas := Render(camera, theWorld)

	fmt.Printf("\n%f\n", camera.HalfHeight)

	return canvas.Newppm()
}
