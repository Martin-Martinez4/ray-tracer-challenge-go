package main

import "math"

type Camera struct {
	HSize       float64
	VSize       float64
	FieldOfView float64
	Transform   Matrix4x4
	HalfWidth   float64
	HalfHeight  float64
	PixelSize   float64
}

func PixelSize(camera *Camera) {
	halfView := math.Tan(camera.FieldOfView / 2)
	aspect := camera.HSize / camera.VSize

	if aspect >= 1 {
		camera.HalfWidth = halfView
		camera.HalfHeight = halfView / float64(aspect)
	} else {
		camera.HalfWidth = halfView * float64(aspect)
		camera.HalfHeight = halfView
	}

	camera.PixelSize = (camera.HalfWidth * 2) / float64(camera.HSize)

}

func NewCamera(hsize float64, vsize float64, fieldOfView float64) Camera {

	tempCamera := Camera{HSize: hsize, VSize: vsize, FieldOfView: fieldOfView, Transform: IdentitiyMatrix4x4()}
	PixelSize(&tempCamera)
	return tempCamera

}

func RayForPixel(camera Camera, px, py float64) Ray {
	xoffset := (px + 0.5) * camera.PixelSize
	yoffset := (py + 0.5) * camera.PixelSize

	worldX := camera.HalfWidth - xoffset
	worldY := camera.HalfHeight - yoffset

	cameraTransformInv := camera.Transform.Inverse()
	pixel := cameraTransformInv.TupleMultiply(Point(worldX, worldY, -1))
	origin := cameraTransformInv.TupleMultiply(Point(0, 0, 0))
	direction := Normalize(pixel.Subtract(origin))

	return NewRay([3]float64{origin.x, origin.y, origin.z}, [3]float64{direction.x, direction.y, direction.z})
}
