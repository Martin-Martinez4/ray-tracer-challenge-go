package shapes

import (
	"sort"

	pm "github.com/Martin-Martinez4/ray-tracer-challenge-go/primitive_math"
)

var Union = "union"
var Intersect = "intersect"
var Difference = "difference"

// constructive solid geometry
type CSG struct {
	*PrimeShape
	Operation string
	Bounds    *BoundingBox

	LeftShape  Shape
	RightShape Shape
}

func NewCSG(operation string, leftShape Shape, rightShape Shape) *CSG {

	csg := CSG{
		PrimeShape: CreateDefaultPrimeShape(),
		Operation:  operation,
		Bounds:     nil,

		LeftShape:  leftShape,
		RightShape: rightShape,
	}
	rightShape.SetParent(&csg)
	leftShape.SetParent(&csg)

	return &csg
}

func (csg *CSG) IntersectionAllowed(lhit, inl, inr bool) bool {
	switch csg.Operation {
	case Union:
		return (lhit && !inr) || (!lhit && !inl)
	case Intersect:
		return (lhit && inr) || (!lhit && inl)
	case Difference:
		return (lhit && !inr) || (!lhit && inl)
	default:
		return false
	}
}

func (csg *CSG) FilterIntersection(xs []Intersection) []Intersection {
	inl := false
	inr := false

	result := []Intersection{}

	for _, intersection := range xs {
		// Change to be recursive later
		// Going to have to add another function to Shape
		lhit := csg.LeftShape == intersection.S

		if csg.IntersectionAllowed(lhit, inl, inr) {
			result = append(result, intersection)
		}

		if lhit {
			inl = !inl
		} else {
			inr = !inr
		}
	}

	return result
}

func (csg *CSG) LocalIntersect(ray Ray) *Intersections {
	ray = ray.Transform(csg.Transforms.Inverse())

	leftxs := csg.LeftShape.Intersect(&ray)
	rightxs := csg.RightShape.Intersect(&ray)

	xs := append(*leftxs, *rightxs...)

	sort.Slice(xs, func(i, j int) bool {
		return xs[i].T < xs[j].T
	})
	return &xs

}

func (csg *CSG) Intersect(ray *Ray) *Intersections {
	tray := ray.Transform(csg.Transforms.Inverse())
	return csg.LocalIntersect(tray)
}

func (csg *CSG) LocalNormalAt(localPoint pm.Tuple, hitPoint *pm.Tuple, intersection *Intersection) pm.Tuple {
	return pm.Vector(0, 0, 0)
}

func (csg *CSG) BoundingBox() *BoundingBox {

	return &BoundingBox{}
}
