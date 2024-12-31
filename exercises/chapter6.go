package exercises

import (
	mat "github.com/Martin-Martinez4/ray-tracer-challenge-go/materials"
	pm "github.com/Martin-Martinez4/ray-tracer-challenge-go/primitive_math"
	"github.com/Martin-Martinez4/ray-tracer-challenge-go/shapes"
	"github.com/Martin-Martinez4/ray-tracer-challenge-go/world"
)

func Ch6() string {

	rayOrigin := pm.Point(0, 0, -5)
	wallZ := 10.0
	wallSize := 7.0
	half := wallSize / 2

	canvasPixels := 100.0

	// shadowColor := NewColor(1, 0, 0)

	pixelSize := wallSize / canvasPixels

	canvas := world.NewCanvas(int32(canvasPixels), int32(canvasPixels))
	// shadowColor := NewColor(1, 0, 0)
	sphere := shapes.NewSphere()
	sphere.Material.Color = mat.NewColor(0.2, 0.2, 1)

	// sphere.Scale(1, 2, 1)

	light := world.NewLight([3]float64{-10, 10, -10}, [3]float64{1, 1, 1})

	for y := 0; y < int(canvasPixels); y++ {
		worldY := half - pixelSize*float64(y)

		for x := 0; x < int(canvasPixels); x++ {
			worldX := -half + pixelSize*float64(x)

			position := pm.Point(worldX, worldY, wallZ)

			substractedPosition := position.Subtract(rayOrigin)
			normalized := pm.Normalize(substractedPosition)

			ray := shapes.NewRay(
				[3]float64{rayOrigin.X, rayOrigin.Y, rayOrigin.Z},
				[3]float64{normalized.X, normalized.Y, normalized.Z},
			)

			// intersect
			xs := shapes.RaySphereInteresect(ray, sphere)
			if xs != nil {
				intersection, found := shapes.Hit(xs.Intersections)
				if found {

					point := shapes.Position(ray, intersection.T)
					sphere1, ok := intersection.S.(*shapes.Sphere)
					if !ok {
						panic("Not a Sphere")
					}
					normal := sphere1.NormalAt(point)

					eye := ray.Direction.SMultiply(-1)

					color := world.EffectiveLighting(sphere1.Material, sphere1, light, point, eye, normal, false)

					canvas.ColorPixel(int32(x), int32(y), color)
				}
			}
		}
	}

	return canvas.Newppm()

}
