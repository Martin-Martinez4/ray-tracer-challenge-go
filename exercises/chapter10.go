package exercises

import (
	"math"

	mat "github.com/Martin-Martinez4/ray-tracer-challenge-go/materials"
	pm "github.com/Martin-Martinez4/ray-tracer-challenge-go/primitive_math"
	"github.com/Martin-Martinez4/ray-tracer-challenge-go/shapes"
	"github.com/Martin-Martinez4/ray-tracer-challenge-go/world"
)

// Same as ch7 but with shadows
func ch10() string {
	// Floor
	floor := shapes.NewPlane()
	floor.SetTransforms([]*pm.Matrix4x4{pm.Translate(0, 0, 5)})
	floorPat := mat.NewRing(mat.WHITE, mat.NewColor(1, 1, 0))
	// floorPat.SetTransform(pm.Scale(10, 10, 10))
	floor.GetMaterial().Pattern = floorPat

	// Spheres
	middleSphere := shapes.NewSphere()
	middleSphere.SetTransforms([]*pm.Matrix4x4{pm.Scale(1.1, 1.1, 1.1), pm.Translate(-0.5, 1, -1.1), pm.RotationAlongX(90 * (math.Pi / 180))})
	middleSphere.Material.Color = mat.NewColor(0.1, 1, 0.5)
	middleSphere.Material.Diffuse = 0.7
	middleSphere.Material.Specular = 0.3

	middPat := mat.NewRing(mat.WHITE, mat.NewColor(0, 1, 0))
	middPat.SetTransform(pm.Scale(.025, .025, .025))

	middleSphere.GetMaterial().Pattern = middPat

	rightSphere := shapes.NewSphere()
	rightSphere.Transforms = rightSphere.Transforms.Scale(0.5, 0.5, 0.5)
	rightSphere.Transforms = rightSphere.Transforms.Translate(1.5, 0.5, -0.5)
	rightSphere.SetTransform(pm.RotationAlongY(50 * (math.Pi / 180)))
	rightSphere.Material.Color = mat.NewColor(0.5, 1, 0.1)
	rightSphere.Material.Diffuse = 0.7
	rightSphere.Material.Specular = 0.3
	rightPattern := mat.NewStripe(mat.WHITE, mat.BLACK)
	rightPattern.SetTransform(pm.Scale(.07, .07, .07))
	rightSphere.Material.Pattern = rightPattern

	leftSphere := shapes.NewSphere()
	leftSphere.Transforms = leftSphere.Transforms.Scale(0.33, 0.33, 0.33)
	leftSphere.Transforms = leftSphere.Transforms.Translate(-1.5, 0.33, -0.75)
	leftSphere.Material.Color = mat.NewColor(1, 0.8, 0.1)
	leftSphere.Material.Diffuse = 0.7
	leftSphere.Material.Specular = 0.3

	// Light Source
	w := world.NewDefaultWorld()
	w.Shapes = []shapes.Shape{floor, middleSphere, leftSphere, rightSphere}
	w.Light = world.NewLight([3]float64{-10, 10, -10}, [3]float64{1, 1, 1})

	camera := world.NewCamera(800, 400, math.Pi/3)
	camera.Transform = pm.ViewTransformation(pm.Point(0, 1.5, -5), pm.Point(0, 1, 0), pm.Vector(0, 1, 0))

	canvas := world.Render(camera, w)

	return canvas.Newppm()

}
