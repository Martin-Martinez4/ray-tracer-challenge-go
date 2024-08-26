package main

import (
	"fmt"
	"math"
)

func ch15A() string {
	// Floor
	floor := NewPlane()
	floor.SetTransforms([]*Matrix4x4{Translate(0, 0, 5)})
	floorPat := NewRing(WHITE, NewColor(1, 1, 0))
	// floorPat.SetTransform(Scale(10, 10, 10))
	floor.GetMaterial().Pattern = floorPat

	// Spheres
	tri1 := NewTriangle(Point(0, 0, 0), Point(2, 2, 0), Point(2, 0, 0))
	tri1.SetTransforms([]*Matrix4x4{Scale(1.1, 1.1, 1.1)})
	tri1.Material.Color = NewColor(0.1, 1, 0.5)
	tri1.Material.Diffuse = 0.7
	tri1.Material.Specular = 0.3
	tri1.Material.Reflective = 0.5

	tri2 := NewTriangle(Point(0, 0, 0), Point(2, 2, 0), Point(0, 2, 0))
	tri2.SetTransforms([]*Matrix4x4{Scale(1.1, 1.1, 1.1)})
	tri2.Material.Color = NewColor(1, 0.5, 0.5)
	tri2.Material.Diffuse = 0.7
	tri2.Material.Specular = 0.3
	tri2.Material.Reflective = 0.5

	group := NewGroup()
	group.AddChild(tri1)
	group.AddChild(tri2)

	// middPat := NewRing(WHITE, NewColor(0, 1, 0))
	// middPat.SetTransform(Scale(.025, .025, .025))

	// middleSphere.GetMaterial().Pattern = middPat

	// rightSphere := NewSphere()
	// rightSphere.Transforms = rightSphere.Transforms.Scale(0.5, 0.5, 0.5)
	// rightSphere.Transforms = rightSphere.Transforms.Translate(1.5, 0.5, -0.5)
	// rightSphere.SetTransform(RotationAlongY(50 * (math.Pi / 180)))
	// rightSphere.Material.Color = NewColor(0.5, 1, 0.1)
	// rightSphere.Material.Diffuse = 0.7
	// rightSphere.Material.Specular = 0.3
	// rightPattern := NewStripe(WHITE, BLACK)
	// rightPattern.SetTransform(Scale(.07, .07, .07))
	// rightSphere.Material.Pattern = rightPattern

	// leftSphere := NewSphere()
	// leftSphere.Transforms = leftSphere.Transforms.Scale(0.33, 0.33, 0.33)
	// leftSphere.Transforms = leftSphere.Transforms.Translate(-1.5, 0.33, -0.75)
	// leftSphere.Material.Color = NewColor(1, 0.8, 0.1)
	// leftSphere.Material.Diffuse = 0.7
	// leftSphere.Material.Specular = 0.3

	// // Light Source
	world := NewDefaultWorld()
	world.Shapes = []Shape{floor, group}
	world.Light = NewLight([3]float64{-10, 10, -10}, [3]float64{1, 1, 1})

	camera := NewCamera(400, 400, math.Pi/3)
	camera.Transform = ViewTransformation(Point(0, 1.5, -5), Point(0, 1, 0), Vector(0, 1, 0))

	canvas := Render(camera, world)

	return canvas.Newppm()

}

func ch15B() string {
	// Floor
	// floor := NewPlane()
	// floor.SetTransforms([]*Matrix4x4{Translate(0, 0, 5)})
	// floorPat := NewRing(WHITE, NewColor(1, 1, 0))
	// // floorPat.SetTransform(Scale(10, 10, 10))
	// floor.GetMaterial().Pattern = floorPat

	teapot := ParseObjFile("obj-files/fanTris.obj")
	fmt.Println("obj parsed")

	// // Light Source
	world := NewDefaultWorld()
	world.Shapes = []Shape{teapot.parserToGroup()}
	world.Light = NewLight([3]float64{-10, 10, -10}, [3]float64{1, 1, 1})

	camera := NewCamera(400, 400, math.Pi/3)
	camera.Transform = ViewTransformation(Point(0, 1.5, -5), Point(0, 1, 0), Vector(0, 1, 0))

	canvas := Render(camera, world)

	return canvas.Newppm()

}
