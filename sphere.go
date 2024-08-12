package main

import (
	"math"

	"github.com/google/uuid"
)

type Sphere struct {
	id         uuid.UUID
	Material   Material
	Transforms Matrix4x4
}

func (sphere *Sphere) GetTransforms() Matrix4x4 {
	return sphere.Transforms
}

func NewSphere() *Sphere {
	id, err := uuid.NewUUID()
	identityMatix := NewMatrix4x4([16]float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1})

	if err != nil {
		panic("not able to cerate unique id for sphere")
	}
	return &Sphere{
		id:         id,
		Transforms: identityMatix,
		Material:   DefaultMaterial(),
	}
}

func (sphere *Sphere) SetTransform(mat44 *Matrix4x4) Matrix4x4 {
	sphere.Transforms = mat44.Multiply(sphere.Transforms)
	return sphere.Transforms
}

func (sphere *Sphere) SetTransforms(mat44 ...*Matrix4x4) {

	for _, transform := range mat44 {

		sphere.SetTransform(transform)
	}
}

func (sphere *Sphere) GetMaterial() *Material {
	return &sphere.Material
}

func (sphere *Sphere) SetMaterial(material Material) {
	sphere.Material = material
}

func (sphere *Sphere) LocalIntersect(ray Ray) Intersections {

	inters := Intersections{}

	sphereToRay := ray.origin.Subtract(Point(0, 0, 0))

	a := Dot(ray.direction, ray.direction)
	b := 2 * Dot(ray.direction, sphereToRay)
	c := Dot(sphereToRay, sphereToRay) - 1

	discriminant := (b * b) - (4*a)*c

	if discriminant < 0 {

	} else {
		d1 := (-b - math.Sqrt(discriminant)) / (2 * a)
		d2 := (-b + math.Sqrt(discriminant)) / (2 * a)

		if !AreFloatsEqual(d1, d2) {

			inters.Add(Intersection{d1, sphere})
			inters.Add(Intersection{d2, sphere})

		} else {

			inters.Add(Intersection{d1, sphere})

		}
	}

	return inters
}

func (sphere *Sphere) Intersect(ray *Ray) Intersections {

	localRay := ray.Transform(sphere.Transforms.Inverse())
	return sphere.LocalIntersect(localRay)

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

func (sphere *Sphere) NormalAt(worldPoint Tuple) Tuple {
	invTransf := sphere.GetTransforms().Inverse()
	objectPoint := invTransf.TupleMultiply(worldPoint)

	objectNormal := objectPoint.Subtract(Point(0, 0, 0))

	invTransfTransposed := invTransf.Transpose()
	worldNormal := invTransfTransposed.TupleMultiply(objectNormal)
	worldNormal.w = 0
	return Normalize(worldNormal)
}
