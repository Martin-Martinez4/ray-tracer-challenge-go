package main

import (
	"fmt"
	"log"
	"os"
)

func CastShadow() string {

	rayOrigin := Point(0, 0, -5)
	wallZ := 10.0
	wallSize := 7.0
	half := wallSize / 2

	canvasPixels := 100.0

	pixelSize := wallSize / canvasPixels

	canvas := NewCanvas(int32(canvasPixels), int32(canvasPixels))
	shadowColor := NewColor(1, 0, 0)
	sphere := NewSphere()

	for y := 0; y < int(canvasPixels); y++ {
		worldY := half - pixelSize*float64(y)

		for x := 0; x < int(canvasPixels); x++ {
			worldX := -half + pixelSize*float64(x)

			position := Point(worldX, worldY, wallZ)

			normalized := Normalize(rayOrigin.Subtract(position))

			ray := NewRay(
				[3]float64{rayOrigin.x, rayOrigin.y, rayOrigin.z},
				[3]float64{normalized.x, normalized.y, normalized.z},
			)

			// intersect
			xs := RaySphereInteresect(ray, &sphere)
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
