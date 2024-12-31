package shapes

import (
	"math"

	pm "github.com/Martin-Martinez4/ray-tracer-challenge-go/primitive_math"
)

type Sphere struct {
	*PrimeShape
	Parent Shape
	Bounds *BoundingBox
}

func NewGlassSphere() *Sphere {
	sphere := NewSphere()
	sphere.Material.Transparency = 1
	sphere.Material.RefractiveIndex = 1.5

	return sphere
}

func NewSphere() *Sphere {

	return &Sphere{
		PrimeShape: CreateDefaultPrimeShape(),
		Parent:     nil,
		Bounds:     nil,
	}
}

func (sphere *Sphere) SetTransform(mat44 *pm.Matrix4x4) pm.Matrix4x4 {
	sphere.Transforms = mat44.Multiply(sphere.Transforms)
	return sphere.Transforms
}

func (sphere *Sphere) SetTransforms(mat44 []*pm.Matrix4x4) {

	for _, transform := range mat44 {

		sphere.SetTransform(transform)
	}
}

func (sphere *Sphere) LocalIntersect(ray Ray) Intersections {

	inters := Intersections{}

	sphereToRay := ray.origin.Subtract(pm.Point(0, 0, 0))

	a := pm.Dot(ray.Direction, ray.Direction)
	b := 2 * pm.Dot(ray.Direction, sphereToRay)
	c := pm.Dot(sphereToRay, sphereToRay) - 1

	discriminant := (b * b) - (4*a)*c

	if discriminant < 0 {

	} else {
		d1 := (-b - math.Sqrt(discriminant)) / (2 * a)
		d2 := (-b + math.Sqrt(discriminant)) / (2 * a)

		if !pm.AreFloatsEqual(d1, d2) {

			inters.Add(NewIntersection(d1, sphere))
			inters.Add(NewIntersection(d2, sphere))

		} else {

			inters.Add(NewIntersection(d1, sphere))

		}
	}

	return inters
}

func (sphere *Sphere) Intersect(ray *Ray) Intersections {

	sphere.SetSavedRay(ray.Transform(sphere.Transforms.Inverse()))
	return sphere.LocalIntersect(sphere.GetSavedRay())

}

func (sphere *Sphere) Translate(x, y, z float64) {
	newTransform := sphere.GetTransforms().Translate(x, y, z)
	sphere.Transforms = newTransform
}

func (sphere *Sphere) Scale(x, y, z float64) {
	newTransform := sphere.GetTransforms().Scale(x, y, z)
	sphere.Transforms = newTransform
}

func (sphere *Sphere) RotationAlongZ(rads float64) {
	newTransform := sphere.GetTransforms().RotationAlongZ(rads)
	sphere.Transforms = newTransform
}

func (sphere *Sphere) LocalNormalAt(localPoint pm.Tuple, hitPoint *pm.Tuple, intersection *Intersection) pm.Tuple {
	// return localPoint.Subtract(pm.Point(0, 0, 0))

	invTransf := sphere.GetTransforms().Inverse()
	objectPoint := invTransf.TupleMultiply(localPoint)
	objectNormal := objectPoint.Subtract(pm.Point(0, 0, 0))

	invTransfTransposed := invTransf.Transpose()
	worldNormal := invTransfTransposed.TupleMultiply(objectNormal)
	worldNormal.W = 0
	return pm.Normalize(worldNormal)
}

func (sphere *Sphere) NormalAt(worldPoint pm.Tuple) pm.Tuple {
	invTransf := sphere.GetTransforms().Inverse()
	objectPoint := invTransf.TupleMultiply(worldPoint)

	objectNormal := objectPoint.Subtract(pm.Point(0, 0, 0))

	invTransfTransposed := invTransf.Transpose()
	worldNormal := invTransfTransposed.TupleMultiply(objectNormal)
	worldNormal.W = 0
	return pm.Normalize(worldNormal)
}

func (sphere *Sphere) GetParent() Shape {
	return sphere.Parent
}

func (sphere *Sphere) SetParent(shape Shape) {
	sphere.Parent = shape
}

func (sphere *Sphere) BoundingBox() *BoundingBox {
	if sphere.Bounds == nil {
		sphere.Bounds = &BoundingBox{
			Minimum: pm.Point(-1, -1, -1),
			Maximum: pm.Point(1, 1, 1),
		}
	}

	return sphere.Bounds
}
