package main

func ch6() string {

	rayOrigin := Point(0, 0, -5)
	wallZ := 10.0
	wallSize := 7.0
	half := wallSize / 2

	canvasPixels := 100.0

	// shadowColor := NewColor(1, 0, 0)

	pixelSize := wallSize / canvasPixels

	canvas := NewCanvas(int32(canvasPixels), int32(canvasPixels))
	// shadowColor := NewColor(1, 0, 0)
	sphere := NewSphere()
	sphere.Material.Color = NewColor(0.2, 0.2, 1)

	// sphere.Scale(1, 2, 1)

	light := NewLight([3]float64{-10, 10, -10}, [3]float64{1, 1, 1})

	for y := 0; y < int(canvasPixels); y++ {
		worldY := half - pixelSize*float64(y)

		for x := 0; x < int(canvasPixels); x++ {
			worldX := -half + pixelSize*float64(x)

			position := Point(worldX, worldY, wallZ)

			substractedPosition := position.Subtract(rayOrigin)
			normalized := Normalize(substractedPosition)

			ray := NewRay(
				[3]float64{rayOrigin.x, rayOrigin.y, rayOrigin.z},
				[3]float64{normalized.x, normalized.y, normalized.z},
			)

			// intersect
			xs := RaySphereInteresect(ray, sphere)
			if xs != nil {
				intersection, found := Hit(xs)
				if found {

					point := Position(ray, intersection.T)
					sphere1, ok := intersection.S.(*Sphere)
					if !ok {
						panic("Not a Sphere")
					}
					normal := sphere1.NormalAt(point)

					eye := ray.direction.SMultiply(-1)

					color := EffectiveLighting(sphere1.Material, sphere1, light, point, eye, normal, false)

					canvas.ColorPixel(int32(x), int32(y), color)
				}
			}
		}
	}

	return canvas.Newppm()

}
