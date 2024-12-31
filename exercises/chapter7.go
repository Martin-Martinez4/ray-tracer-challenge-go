package exercises

import (
	"math"

	mat "github.com/Martin-Martinez4/ray-tracer-challenge-go/materials"
	pm "github.com/Martin-Martinez4/ray-tracer-challenge-go/primitive_math"
	"github.com/Martin-Martinez4/ray-tracer-challenge-go/shapes"
	"github.com/Martin-Martinez4/ray-tracer-challenge-go/world"
)

func ch7() string {
	// Floor
	floor := shapes.NewSphere()
	floor.Transforms = floor.Transforms.Scale(10, 0.01, 10)
	floor.Material.Color = mat.NewColor(1, 0.9, 0.9)
	floor.Material.Specular = 0

	leftWall := shapes.NewSphere()
	leftWall.Transforms = leftWall.Transforms.Scale(10, 0.01, 10)
	leftWall.Transforms = leftWall.Transforms.RotationAlongX(math.Pi / 2)
	leftWall.Transforms = leftWall.Transforms.RotationAlongY(-math.Pi / 4)
	leftWall.Transforms = leftWall.Transforms.Translate(0, 0, 5)

	rightWall := shapes.NewSphere()
	rightWall.Transforms = rightWall.Transforms.Scale(10, 0.01, 10)
	rightWall.Transforms = rightWall.Transforms.RotationAlongX(math.Pi / 2)
	rightWall.Transforms = rightWall.Transforms.RotationAlongY(math.Pi / 4)
	rightWall.Transforms = rightWall.Transforms.Translate(0, 0, 5)

	// Spheres
	middleSphere := shapes.NewSphere()
	middleSphere.Transforms = middleSphere.Transforms.Translate(-0.5, 1, 0.5)
	middleSphere.Material.Color = mat.NewColor(0.1, 1, 0.5)
	middleSphere.Material.Diffuse = 0.7
	middleSphere.Material.Specular = 0.3

	rightSphere := shapes.NewSphere()
	rightSphere.Transforms = rightSphere.Transforms.Scale(0.5, 0.5, 0.5)
	rightSphere.Transforms = rightSphere.Transforms.Translate(1.5, 0.5, -0.5)
	rightSphere.Material.Color = mat.NewColor(0.5, 1, 0.1)
	rightSphere.Material.Diffuse = 0.7
	rightSphere.Material.Specular = 0.3

	leftSphere := shapes.NewSphere()
	leftSphere.Transforms = leftSphere.Transforms.Scale(0.33, 0.33, 0.33)
	leftSphere.Transforms = leftSphere.Transforms.Translate(-1.5, 0.33, -0.75)
	leftSphere.Material.Color = mat.NewColor(1, 0.8, 0.1)
	leftSphere.Material.Diffuse = 0.7
	leftSphere.Material.Specular = 0.3

	// Light Source
	w := world.NewDefaultWorld()
	w.Shapes = []shapes.Shape{floor, leftWall, rightWall, middleSphere, leftSphere, rightSphere}
	w.Light = world.NewLight([3]float64{-10, 10, -10}, [3]float64{1, 1, 1})

	camera := world.NewCamera(100, 50, math.Pi/3)
	camera.Transform = pm.ViewTransformation(pm.Point(0, 1.5, -5), pm.Point(0, 1, 0), pm.Vector(0, 1, 0))

	canvas := world.Render(camera, w)

	return canvas.Newppm()

}
