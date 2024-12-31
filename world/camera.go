package world

import (
	"math"

	pm "github.com/Martin-Martinez4/ray-tracer-challenge-go/primitive_math"
	"github.com/Martin-Martinez4/ray-tracer-challenge-go/shapes"
)

type Camera struct {
	HSize       float64
	VSize       float64
	FieldOfView float64
	Transform   pm.Matrix4x4
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

	tempCamera := Camera{HSize: hsize, VSize: vsize, FieldOfView: fieldOfView, Transform: pm.IdentitiyMatrix4x4()}
	PixelSize(&tempCamera)
	return tempCamera

}

func RayForPixel(camera Camera, px, py float64) shapes.Ray {
	xoffset := (px + 0.5) * camera.PixelSize
	yoffset := (py + 0.5) * camera.PixelSize

	worldX := camera.HalfWidth - xoffset
	worldY := camera.HalfHeight - yoffset

	cameraTransformInv := camera.Transform.Inverse()
	pixel := cameraTransformInv.TupleMultiply(pm.Point(worldX, worldY, -1))
	origin := cameraTransformInv.TupleMultiply(pm.Point(0, 0, 0))
	direction := pm.Normalize(pixel.Subtract(origin))

	return shapes.NewRay([3]float64{origin.X, origin.Y, origin.Z}, [3]float64{direction.X, direction.Y, direction.Z})
}
