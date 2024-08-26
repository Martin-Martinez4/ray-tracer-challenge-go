package main

import (
	"math"

	"github.com/google/uuid"
)

type SmoothTriangle struct {
	id uuid.UUID

	P1     Tuple
	P2     Tuple
	P3     Tuple
	Normal Tuple

	N1 Tuple
	N2 Tuple
	N3 Tuple

	E1 Tuple
	E2 Tuple

	Transforms Matrix4x4
	Material   Material
	Parent     Shape
	Bounds     *BoundingBox
	SavedRay   Ray
}

func NewSmoothTriangle(p1, p2, p3, n1, n2, n3 Tuple) *SmoothTriangle {

	id, err := uuid.NewUUID()
	identityMatix := NewMatrix4x4([16]float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1})

	if err != nil {
		panic("not able to cerate unique id for sphere")
	}
	// e1 = edge 1
	// p1 = point1
	e1 := p2.Subtract(p1)
	e2 := p3.Subtract(p1)

	normal := Normalize(Cross(e2, e1))

	return &SmoothTriangle{
		id: id,

		P1: p1,
		P2: p2,
		P3: p3,

		N1: n1,
		N2: n2,
		N3: n3,

		E1: e1,
		E2: e2,

		Normal:     normal,
		Transforms: identityMatix,
		Material:   DefaultMaterial(),
		Parent:     nil,
		Bounds:     nil,
	}
}

func (smoothTriangle *SmoothTriangle) Equal(other *SmoothTriangle) bool {
	return smoothTriangle.P1.Equal(other.P1) && smoothTriangle.P2.Equal(other.P2) && smoothTriangle.P3.Equal(other.P3)
}

func (smoothTriangle *SmoothTriangle) GetId() uuid.UUID {
	return smoothTriangle.id
}

func (smoothTriangle *SmoothTriangle) GetParent() Shape {
	return smoothTriangle.Parent
}
func (smoothTriangle *SmoothTriangle) SetParent(shape Shape) {
	smoothTriangle.Parent = shape
}

func (smoothTriangle *SmoothTriangle) GetTransforms() Matrix4x4 {
	return smoothTriangle.Transforms
}
func (smoothTriangle *SmoothTriangle) SetTransform(transform *Matrix4x4) Matrix4x4 {
	smoothTriangle.Transforms = transform.Multiply(smoothTriangle.Transforms)
	return smoothTriangle.Transforms
}
func (smoothTriangle *SmoothTriangle) SetTransforms(transform []*Matrix4x4) {
	for i := 0; i < len(transform); i++ {
		smoothTriangle.SetTransform(transform[i])
	}
}

func (smoothTriangle *SmoothTriangle) GetMaterial() *Material {
	return &smoothTriangle.Material
}
func (smoothTriangle *SmoothTriangle) SetMaterial(material Material) {
	smoothTriangle.Material = material
}

func (smoothTriangle *SmoothTriangle) LocalNormalAt(localPoint Tuple, hitPoint *Tuple, intersection *Intersection) Tuple {
	invTransf := smoothTriangle.GetTransforms().Inverse()

	n2U := smoothTriangle.N2.SMultiply(*intersection.U)
	n3V := smoothTriangle.N3.SMultiply(*intersection.V)
	n1UV := smoothTriangle.N1.SMultiply(1 - (*intersection.U) - (*intersection.V))

	n2Un3V := n2U.Add(n3V)
	objectNormal := n2Un3V.Add(n1UV)

	invTransfTransposed := invTransf.Transpose()
	worldNormal := invTransfTransposed.TupleMultiply(objectNormal)
	worldNormal.w = 0
	return Normalize(worldNormal)

}
func (smoothTriangle *SmoothTriangle) NormalAt(localPoint Tuple, hitPoint *Tuple, intersection *Intersection) Tuple {
	return smoothTriangle.LocalNormalAt(localPoint, hitPoint, intersection)
}

func (smoothTriangle *SmoothTriangle) LocalIntersect(ray Ray) Intersections {
	dirCrossE2 := Cross(ray.direction, smoothTriangle.E2)
	det := Dot(smoothTriangle.E1, dirCrossE2)

	if math.Abs(det) < Epsilon {
		return Intersections{intersections: []Intersection{}}
	}

	f := 1.0 / det

	p1ToOrigin := ray.origin.Subtract(smoothTriangle.P1)
	u := f * Dot(p1ToOrigin, dirCrossE2)

	if u < 0 || u > 1 {
		return Intersections{intersections: []Intersection{}}
	}

	originCrossE1 := Cross(p1ToOrigin, smoothTriangle.E1)
	v := f * Dot(ray.direction, originCrossE1)

	if v < 0 || (u+v) > 1 {
		return Intersections{intersections: []Intersection{}}
	}

	t := f * Dot(smoothTriangle.E2, originCrossE1)
	// will change later
	return Intersections{intersections: []Intersection{NewIntersectionWithUV(t, smoothTriangle, &u, &v)}}
}

func (smoothTriangle *SmoothTriangle) Intersect(ray *Ray) Intersections {
	tray := ray.Transform(smoothTriangle.Transforms.Inverse())
	return smoothTriangle.LocalIntersect(tray)
}

func (smoothTriangle *SmoothTriangle) GetSavedRay() Ray {
	return smoothTriangle.SavedRay
}
func (smoothTriangle *SmoothTriangle) SetSavedRay(ray Ray) {
	smoothTriangle.SavedRay = ray
}

func (smoothTriangle *SmoothTriangle) BoundingBox() *BoundingBox {
	if smoothTriangle.Bounds == nil {
		smoothTriangle.Bounds = &BoundingBox{
			Minimum: Point(-1, -1, -1),
			Maximum: Point(1, 1, 1),
		}
	}

	return smoothTriangle.Bounds
}
