package shapes

import (
	"sort"

	mat "github.com/Martin-Martinez4/ray-tracer-challenge-go/materials"
	pm "github.com/Martin-Martinez4/ray-tracer-challenge-go/primitive_math"
	"github.com/google/uuid"
)

var Union = "union"
var Intersect = "intersect"
var Difference = "difference"

// constructive solid geometry
type CSG struct {
	*PrimeShape
	Operation string
	Parent    Shape
	Bounds    *BoundingBox

	LeftShape  Shape
	RightShape Shape
}

func NewCSG(operation string, leftShape Shape, rightShape Shape) *CSG {

	csg := CSG{
		PrimeShape: CreateDefaultPrimeShape(),
		Operation:  operation,
		Parent:     nil,
		Bounds:     nil,

		LeftShape:  leftShape,
		RightShape: rightShape,
	}
	rightShape.SetParent(&csg)
	leftShape.SetParent(&csg)

	return &csg
}

func (csg *CSG) GetTransforms() pm.Matrix4x4 {
	return csg.Transforms
}
func (csg *CSG) SetTransform(transform *pm.Matrix4x4) pm.Matrix4x4 {
	csg.Transforms = transform.Multiply(csg.Transforms)
	return csg.Transforms
}
func (csg *CSG) SetTransforms(transform []*pm.Matrix4x4) {
	for i := 0; i < len(transform); i++ {
		csg.SetTransform(transform[i])
	}
}

func (csg *CSG) GetMaterial() *mat.Material {
	return &csg.Material
}
func (csg *CSG) SetMaterial(material mat.Material) {
	csg.Material = material
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

func (csg *CSG) LocalIntersect(ray Ray) Intersections {
	ray = ray.Transform(csg.Transforms.Inverse())

	leftxs := csg.LeftShape.Intersect(&ray)
	rightxs := csg.RightShape.Intersect(&ray)

	xs := append(leftxs.Intersections, rightxs.Intersections...)

	sort.Slice(xs, func(i, j int) bool {
		return xs[i].T < xs[j].T
	})
	return Intersections{Intersections: xs}

}

func (csg *CSG) Intersect(ray *Ray) Intersections {
	tray := ray.Transform(csg.Transforms.Inverse())
	return csg.LocalIntersect(tray)
}

func (csg *CSG) LocalNormalAt(localPoint pm.Tuple, hitPoint *pm.Tuple, intersection *Intersection) pm.Tuple {
	return pm.Vector(0, 0, 0)
}

func (csg *CSG) GetSavedRay() Ray {
	return csg.SavedRay
}
func (csg *CSG) SetSavedRay(ray Ray) {
	csg.SavedRay = ray
}

func (csg *CSG) GetParent() Shape {
	return csg.Parent
}

func (csg *CSG) SetParent(shape Shape) {
	csg.Parent = shape
}

func (csg *CSG) GetId() uuid.UUID {
	return csg.id
}

func (csg *CSG) BoundingBox() *BoundingBox {

	return &BoundingBox{}
}
