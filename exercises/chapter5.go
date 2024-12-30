package exercises

import (
	"fmt"
	"log"
	"os"

	mat "github.com/Martin-Martinez4/ray-tracer-challenge-go/materials"
	pm "github.com/Martin-Martinez4/ray-tracer-challenge-go/primitive_math"
	"github.com/Martin-Martinez4/ray-tracer-challenge-go/shapes"
)

func ch5() string {

	rayOrigin := pm.Point(0, 0, -5)
	wallZ := 10.0
	wallSize := 7.0
	half := wallSize / 2

	canvasPixels := 100.0

	pixelSize := wallSize / canvasPixels

	canvas := NewCanvas(int32(canvasPixels), int32(canvasPixels))
	shadowColor := mat.NewColor(1, 0, 0)
	sphere := shapes.NewSphere()

	for y := 0; y < int(canvasPixels); y++ {
		worldY := half - pixelSize*float64(y)

		for x := 0; x < int(canvasPixels); x++ {
			worldX := -half + pixelSize*float64(x)

			position := pm.Point(worldX, worldY, wallZ)

			normalized := pm.Normalize(position.Subtract(rayOrigin))

			ray := shapes.NewRay(
				[3]float64{rayOrigin.X, rayOrigin.Y, rayOrigin.Z},
				[3]float64{normalized.X, normalized.Y, normalized.Z},
			)

			// intersect
			xs := shapes.RaySphereInteresect(ray, sphere)
			if xs != nil {
				canvas.ColorPixel(int32(x), int32(y), shadowColor)
			}
		}
	}

	return canvas.Newppm()

}

func printAnswer(str string) {

	f, err := os.Create("chapter5.ppm")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(str)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("done")
}
