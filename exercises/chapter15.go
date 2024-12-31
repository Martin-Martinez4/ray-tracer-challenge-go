package exercises

import (
	"fmt"
	"math"

	mat "github.com/Martin-Martinez4/ray-tracer-challenge-go/materials"
	pm "github.com/Martin-Martinez4/ray-tracer-challenge-go/primitive_math"
	"github.com/Martin-Martinez4/ray-tracer-challenge-go/shapes"
	"github.com/Martin-Martinez4/ray-tracer-challenge-go/world"
)

func Ch15A() string {
	// Floor
	floor := shapes.NewPlane()
	floor.SetTransforms([]*pm.Matrix4x4{pm.Translate(0, 0, 5)})
	floorPat := mat.NewRing(mat.WHITE, mat.NewColor(1, 1, 0))
	// floorPat.SetTransform(pm.Scale(10, 10, 10))
	floor.GetMaterial().Pattern = floorPat

	// Spheres
	tri1 := shapes.NewTriangle(pm.Point(0, 0, 0), pm.Point(2, 2, 0), pm.Point(2, 0, 0))
	tri1.SetTransforms([]*pm.Matrix4x4{pm.Scale(1.1, 1.1, 1.1)})
	tri1.Material.Color = mat.NewColor(0.1, 1, 0.5)
	tri1.Material.Diffuse = 0.7
	tri1.Material.Specular = 0.3
	tri1.Material.Reflective = 0.5

	tri2 := shapes.NewTriangle(pm.Point(0, 0, 0), pm.Point(2, 2, 0), pm.Point(0, 2, 0))
	tri2.SetTransforms([]*pm.Matrix4x4{pm.Scale(1.1, 1.1, 1.1)})
	tri2.Material.Color = mat.NewColor(1, 0.5, 0.5)
	tri2.Material.Diffuse = 0.7
	tri2.Material.Specular = 0.3
	tri2.Material.Reflective = 0.5

	group := shapes.NewGroup()
	group.AddChild(tri1)
	group.AddChild(tri2)

	// middPat := mat.NewRing(mat.WHITE, mat.NewColor(0, 1, 0))
	// middPat.SetTransform(pm.Scale(.025, .025, .025))

	// middleSphere.GetMaterial().Pattern = middPat

	// rightSphere := NewSphere()
	// rightSphere.Transforms = rightSphere.Transforms.pm.Scale(0.5, 0.5, 0.5)
	// rightSphere.Transforms = rightSphere.Transforms.pm.Translate(1.5, 0.5, -0.5)
	// rightSphere.SetTransform(RotationAlongY(50 * (math.Pi / 180)))
	// rightSphere.Material.Color = mat.NewColor(0.5, 1, 0.1)
	// rightSphere.Material.Diffuse = 0.7
	// rightSphere.Material.Specular = 0.3
	// rightPattern := NewStripe(mat.WHITE, BLACK)
	// rightPattern.SetTransform(pm.Scale(.07, .07, .07))
	// rightSphere.Material.Pattern = rightPattern

	// leftSphere := NewSphere()
	// leftSphere.Transforms = leftSphere.Transforms.pm.Scale(0.33, 0.33, 0.33)
	// leftSphere.Transforms = leftSphere.Transforms.pm.Translate(-1.5, 0.33, -0.75)
	// leftSphere.Material.Color = mat.NewColor(1, 0.8, 0.1)
	// leftSphere.Material.Diffuse = 0.7
	// leftSphere.Material.Specular = 0.3

	// // Light Source
	w := world.NewDefaultWorld()
	w.Shapes = []shapes.Shape{floor, group}
	w.Light = world.NewLight([3]float64{-10, 10, -10}, [3]float64{1, 1, 1})

	camera := world.NewCamera(400, 400, math.Pi/3)
	camera.Transform = pm.ViewTransformation(pm.Point(0, 1.5, -5), pm.Point(0, 1, 0), pm.Vector(0, 1, 0))

	canvas := world.Render(camera, w)

	return canvas.Newppm()

}

func Ch15B() string {
	// Floor
	// floor := NewPlane()
	// floor.SetTransforms([]*pm.Matrix4x4{pm.Translate(0, 0, 5)})
	// floorPat := mat.NewRing(mat.WHITE, mat.NewColor(1, 1, 0))
	// // floorPat.SetTransform(pm.Scale(10, 10, 10))
	// floor.GetMaterial().Pattern = floorPat

	teapot := world.ParseObjFile("obj-files/fanTris.obj")
	fmt.Println("obj parsed")

	// // Light Source
	w := world.NewDefaultWorld()
	w.Shapes = []shapes.Shape{teapot.ParserToGroup()}
	w.Light = world.NewLight([3]float64{-10, 10, -10}, [3]float64{1, 1, 1})

	camera := world.NewCamera(400, 400, math.Pi/3)
	camera.Transform = pm.ViewTransformation(pm.Point(0, 1.5, -5), pm.Point(0, 1, 0), pm.Vector(0, 1, 0))

	canvas := world.Render(camera, w)

	return canvas.Newppm()

}
