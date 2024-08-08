package main

import (
	"github.com/google/uuid"
)

type Sphere struct {
	id         uuid.UUID
	Material   Material
	Transforms Matrix4x4
}

func (sphere Sphere) GetTransforms() Matrix4x4 {
	return sphere.Transforms
}

func NewSphere() Sphere {
	id, err := uuid.NewUUID()
	identityMatix := NewMatrix4x4([16]float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1})

	if err != nil {
		panic("not able to cerate unique id for sphere")
	}
	return Sphere{
		id:         id,
		Transforms: identityMatix,
		Material:   DefaultMaterial(),
	}
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

func (sphere Sphere) NormalAt(worldPoint Tuple) Tuple {
	invTransf := sphere.GetTransforms().Inverse()
	objectPoint := invTransf.TupleMultiply(worldPoint)

	objectNormal := objectPoint.Subtract(Point(0, 0, 0))

	invTransfTransposed := invTransf.Transpose()
	worldNormal := invTransfTransposed.TupleMultiply(objectNormal)
	worldNormal.w = 0
	return Normalize(worldNormal)
}
