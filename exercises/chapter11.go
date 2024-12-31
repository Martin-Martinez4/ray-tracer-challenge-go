package exercises

import (
	"fmt"
	"math"

	mat "github.com/Martin-Martinez4/ray-tracer-challenge-go/materials"
	pm "github.com/Martin-Martinez4/ray-tracer-challenge-go/primitive_math"
	"github.com/Martin-Martinez4/ray-tracer-challenge-go/shapes"
	"github.com/Martin-Martinez4/ray-tracer-challenge-go/world"
)

func Ch11A() string {
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
	middleSphere.Material.Reflective = 0.5

	middPat := mat.NewRing(mat.WHITE, mat.NewColor(0, 1, 0))
	middPat.SetTransform(pm.Scale(.025, .025, .025))

	middleSphere.GetMaterial().Pattern = middPat

	rightSphere := shapes.NewSphere()
	rightSphere.SetTransforms([]*pm.Matrix4x4{pm.Scale(0.5, 0.5, 0.5), pm.Translate(1.5, 0.5, -0.5), pm.RotationAlongY(50 * (math.Pi / 180))})
	rightSphere.Material.Color = mat.NewColor(0.5, 1, 0.1)
	rightSphere.Material.Diffuse = 0.7
	rightSphere.Material.Specular = 0.3
	rightPattern := mat.NewStripe(mat.WHITE, mat.BLACK)
	rightPattern.SetTransform(pm.Scale(.07, .07, .07))
	rightSphere.Material.Pattern = rightPattern

	leftSphere := shapes.NewSphere()
	leftSphere.SetTransforms([]*pm.Matrix4x4{pm.Scale(0.33, 0.33, 0.33), pm.Translate(-1.5, 0.33, -0.75)})
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

func Ch11B() string {
	// Floor

	floor := shapes.NewPlane()
	floor.SetTransform(pm.Translate(0, -1, 0))
	floor.GetMaterial().Transparency = 0.5
	floor.GetMaterial().RefractiveIndex = 1
	floor.GetMaterial().Color = mat.NewColor(0.2, 0.2, 0.2)

	ball := shapes.NewSphere()
	ball.GetMaterial().Ambient = 0
	ball.GetMaterial().Specular = 0.9
	ball.GetMaterial().Transparency = 0.9
	ball.GetMaterial().Reflective = 0.9
	ball.GetMaterial().RefractiveIndex = 1.5
	ball.GetMaterial().Color = mat.NewColor(1, 1, 1)

	glassSphere := shapes.NewGlassSphere()

	// Light Source
	w := world.NewDefaultWorld()
	w.Shapes = []shapes.Shape{floor, glassSphere}
	w.Light = world.NewLight([3]float64{-10, 10, -10}, [3]float64{1, 1, 1})

	camera := world.NewCamera(800, 400, math.Pi/3)
	camera.Transform = pm.ViewTransformation(pm.Point(0, 1.5, -5), pm.Point(0, 1, 0), pm.Vector(0, 1, 0))

	canvas := world.Render(camera, w)

	return canvas.Newppm()

}

func Ch11C() string {
	floor := shapes.NewPlane()
	floor.SetTransform(pm.Translate(0, -10, 0))
	floor.Material.Pattern = mat.NewChecker(mat.BLACK, mat.WHITE)
	floor.Material.Pattern.SetTransform(pm.Translate(0, 0, 0))
	floor.Material.Specular = 0

	bigger := shapes.NewSphere()
	bigger.Material.Diffuse = 0.1
	bigger.Material.Shininess = 300
	bigger.Material.Reflective = 1
	bigger.Material.Transparency = 1
	bigger.Material.RefractiveIndex = 1.52
	bigger.Material.Color = mat.NewColor(0, 0, 0.1)

	smaller := shapes.NewSphere()
	smaller.SetTransform(pm.Scale(0.5, 0.5, 0.5))
	smaller.Material.Diffuse = 0.1
	smaller.Material.Shininess = 300
	smaller.Material.Reflective = 1
	smaller.Material.Transparency = 1
	smaller.Material.RefractiveIndex = 1
	smaller.Material.Color = mat.NewColor(0, 0, 0.1)

	light := world.NewLight([3]float64{2, 10, -5}, [3]float64{0.9, 0.9, 0.9})
	w := world.NewWorld(&[]shapes.Shape{floor, bigger, smaller}, &light)

	camera := world.NewCamera(512, 512, math.Pi/3)
	from := pm.Point(0, 2.5, 0)
	to := pm.Point(0, 0, 0)
	up := pm.Vector(1, 0, 0)
	camera.Transform = pm.ViewTransformation(from, to, up)

	canvas := world.Render(camera, w)

	return canvas.Newppm()
}

func Ch11D() string {
	floor := shapes.NewPlane()
	floor.SetTransforms([]*pm.Matrix4x4{pm.RotationAlongX(1.5708), pm.Translate(0, 0, 10)})
	floor.Material.Pattern = mat.NewChecker(mat.BLACK, mat.WHITE)
	floor.Material.Specular = 0
	floor.Material.Ambient = 0.8
	floor.Material.Diffuse = 0.2

	bigger := shapes.NewSphere()
	bigger.Material.Diffuse = 0
	bigger.Material.Ambient = 0
	bigger.Material.Shininess = 300
	bigger.Material.Reflective = .9
	bigger.Material.Transparency = .9
	bigger.Material.Specular = .9
	bigger.Material.RefractiveIndex = 1.52
	bigger.Material.Color = mat.NewColor(1, 1, 1)

	smaller := shapes.NewSphere()
	smaller.SetTransform(pm.Scale(0.5, 0.5, 0.5))
	smaller.Material.Diffuse = 0
	smaller.Material.Ambient = 0
	smaller.Material.Specular = 0
	smaller.Material.Shininess = 300
	smaller.Material.Reflective = .9
	smaller.Material.Transparency = .9
	smaller.Material.RefractiveIndex = 1.000034
	smaller.Material.Color = mat.NewColor(1, 1, 1)

	light := world.NewLight([3]float64{2, 10, -5}, [3]float64{0.9, 0.9, 0.9})
	w := world.NewWorld(&[]shapes.Shape{floor, bigger, smaller}, &light)

	camera := world.NewCamera(300, 300, 0.45)
	from := pm.Point(0, 0, -5)
	to := pm.Point(0, 0, 0)
	up := pm.Vector(0, 1, 0)
	camera.Transform = pm.ViewTransformation(from, to, up)

	canvas := world.Render(camera, w)

	fmt.Printf("\n%f\n", camera.HalfHeight)

	return canvas.Newppm()
}

func Ch11D3() string {
	theWorld := world.NewDefaultWorld()

	floor := shapes.NewPlane()
	floor.GetMaterial().Transparency = 0.5
	floor.GetMaterial().RefractiveIndex = 1.5
	floor.SetTransform(pm.Translate(0, -1, 0))

	ball := shapes.NewSphere()
	ball.GetMaterial().Color = mat.NewColor(1, 0, 0)
	ball.GetMaterial().Ambient = 0.5
	ball.SetTransform(pm.Translate(0, -2.5, -0.5))

	theWorld.Shapes = append(theWorld.Shapes, floor, ball)

	// light := world.NewLight([3]float64{2, 10, -5}, [3]float64{0.9, 0.9, 0.9})

	// theWorld.Light = light

	// camera := world.NewCamera(500, 500, 0.45)
	// from := pm.Point(0, 0, -5)
	// to := pm.Point(0, 0, 0)
	// up := pm.Vector(0, 1, 0)
	camera := world.NewCamera(400, 400, math.Pi/3)
	camera.Transform = pm.ViewTransformation(pm.Point(0, 1.5, -5), pm.Point(0, 1, 0), pm.Vector(0, 1, 0))

	canvas := world.Render(camera, theWorld)

	fmt.Printf("\n%f\n", camera.HalfHeight)

	return canvas.Newppm()
}
