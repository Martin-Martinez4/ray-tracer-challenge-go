package main

import (
	"sort"

	"github.com/google/uuid"
)

var Union = "union"
var Intersect = "intersect"
var Difference = "difference"

// constructive solid geometry
type CSG struct {
	id         uuid.UUID
	Operation  string
	Material   Material
	Transforms Matrix4x4
	SavedRay   Ray
	Parent     Shape
	Bounds     *BoundingBox

	LeftShape  Shape
	RightShape Shape
}

func NewCSG(operation string, leftShape Shape, rightShape Shape) *CSG {
	id, err := uuid.NewUUID()
	identityMatix := NewMatrix4x4([16]float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1})

	if err != nil {
		panic("not able to cerate unique id for cube")
	}
	csg := CSG{
		id:         id,
		Operation:  operation,
		Transforms: identityMatix,
		Material:   DefaultMaterial(),
		Parent:     nil,
		Bounds:     nil,

		LeftShape:  leftShape,
		RightShape: rightShape,
	}
	rightShape.SetParent(&csg)
	leftShape.SetParent(&csg)

	return &csg
}

func (csg *CSG) GetTransforms() Matrix4x4 {
	return csg.Transforms
}
func (csg *CSG) SetTransform(transform *Matrix4x4) Matrix4x4 {
	csg.Transforms = transform.Multiply(csg.Transforms)
	return csg.Transforms
}
func (csg *CSG) SetTransforms(transform []*Matrix4x4) {
	for i := 0; i < len(transform); i++ {
		csg.SetTransform(transform[i])
	}
}

func (csg *CSG) GetMaterial() *Material {
	return &csg.Material
}
func (csg *CSG) SetMaterial(material Material) {
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

	xs := append(leftxs.intersections, rightxs.intersections...)

	sort.Slice(xs, func(i, j int) bool {
		return xs[i].T < xs[j].T
	})
	return Intersections{intersections: xs}

}

func (csg *CSG) Intersect(ray *Ray) Intersections {
	tray := ray.Transform(csg.Transforms.Inverse())
	return csg.LocalIntersect(tray)
}

func (csg *CSG) LocalNormalAt(localPoint Tuple, hitPoint *Tuple, intersection *Intersection) Tuple {
	return Vector(0, 0, 0)
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
