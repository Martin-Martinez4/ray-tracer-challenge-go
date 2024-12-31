package shapes

import (
	"math"

	pm "github.com/Martin-Martinez4/ray-tracer-challenge-go/primitive_math"
)

type SmoothTriangle struct {
	*PrimeShape
	P1     pm.Tuple
	P2     pm.Tuple
	P3     pm.Tuple
	Normal pm.Tuple

	N1 pm.Tuple
	N2 pm.Tuple
	N3 pm.Tuple

	E1 pm.Tuple
	E2 pm.Tuple

	Bounds *BoundingBox
}

func NewSmoothTriangle(p1, p2, p3, n1, n2, n3 pm.Tuple) *SmoothTriangle {

	// e1 = edge 1
	// p1 = point1
	e1 := p2.Subtract(p1)
	e2 := p3.Subtract(p1)

	normal := pm.Normalize(pm.Cross(e2, e1))

	return &SmoothTriangle{
		PrimeShape: CreateDefaultPrimeShape(),
		P1:         p1,
		P2:         p2,
		P3:         p3,

		N1: n1,
		N2: n2,
		N3: n3,

		E1: e1,
		E2: e2,

		Normal: normal,
		Bounds: nil,
	}
}

func (smoothTriangle *SmoothTriangle) Equal(other *SmoothTriangle) bool {
	return smoothTriangle.P1.Equal(other.P1) && smoothTriangle.P2.Equal(other.P2) && smoothTriangle.P3.Equal(other.P3)
}

func (smoothTriangle *SmoothTriangle) LocalNormalAt(localPoint pm.Tuple, hitPoint *pm.Tuple, intersection *Intersection) pm.Tuple {
	invTransf := smoothTriangle.GetInverseTransforms()

	n2U := smoothTriangle.N2.SMultiply(*intersection.U)
	n3V := smoothTriangle.N3.SMultiply(*intersection.V)
	n1UV := smoothTriangle.N1.SMultiply(1 - (*intersection.U) - (*intersection.V))

	n2Un3V := n2U.Add(n3V)
	objectNormal := n2Un3V.Add(n1UV)

	invTransfTransposed := invTransf.Transpose()
	worldNormal := invTransfTransposed.TupleMultiply(objectNormal)
	worldNormal.W = 0
	return pm.Normalize(worldNormal)

}
func (smoothTriangle *SmoothTriangle) NormalAt(localPoint pm.Tuple, hitPoint *pm.Tuple, intersection *Intersection) pm.Tuple {
	return smoothTriangle.LocalNormalAt(localPoint, hitPoint, intersection)
}

func (smoothTriangle *SmoothTriangle) LocalIntersect(ray Ray) Intersections {
	dirCrossE2 := pm.Cross(ray.Direction, smoothTriangle.E2)
	det := pm.Dot(smoothTriangle.E1, dirCrossE2)

	if math.Abs(det) < pm.Epsilon {
		return Intersections{Intersections: []Intersection{}}
	}

	f := 1.0 / det

	p1ToOrigin := ray.origin.Subtract(smoothTriangle.P1)
	u := f * pm.Dot(p1ToOrigin, dirCrossE2)

	if u < 0 || u > 1 {
		return Intersections{Intersections: []Intersection{}}
	}

	originCrossE1 := pm.Cross(p1ToOrigin, smoothTriangle.E1)
	v := f * pm.Dot(ray.Direction, originCrossE1)

	if v < 0 || (u+v) > 1 {
		return Intersections{Intersections: []Intersection{}}
	}

	t := f * pm.Dot(smoothTriangle.E2, originCrossE1)
	// will change later
	return Intersections{Intersections: []Intersection{NewIntersectionWithUV(t, smoothTriangle, &u, &v)}}
}

func (smoothTriangle *SmoothTriangle) Intersect(ray *Ray) Intersections {
	tray := ray.Transform(smoothTriangle.GetInverseTransforms())
	return smoothTriangle.LocalIntersect(tray)
}

func (smoothTriangle *SmoothTriangle) BoundingBox() *BoundingBox {
	if smoothTriangle.Bounds == nil {
		smoothTriangle.Bounds = &BoundingBox{
			Minimum: pm.Point(-1, -1, -1),
			Maximum: pm.Point(1, 1, 1),
		}
	}

	return smoothTriangle.Bounds
}
