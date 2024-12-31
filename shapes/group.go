package shapes

import (
	"math"
	"sort"

	mat "github.com/Martin-Martinez4/ray-tracer-challenge-go/materials"
	pm "github.com/Martin-Martinez4/ray-tracer-challenge-go/primitive_math"
	"github.com/google/uuid"
)

/*
	type Shape interface {
	GetTransforms() pm.Matrix4x4
	SetTransform(transform *pm.Matrix4x4) pm.Matrix4x4
	SetTransforms(transform []*pm.Matrix4x4)

	GetMaterial() *Material
	SetMaterial(material Material)

	Intersect(ray *Ray) Intersections

	NormalAt(point pm.Tuple) pm.Tuple

	GetSavedRay() Ray
	SetSavedRay(ray Ray)
}
*/

type ChildrenMap map[uuid.UUID]Shape

type Group struct {
	*PrimeShape
	Children ChildrenMap
	Parent   Shape
}

func NewGroup() *Group {

	return &Group{
		PrimeShape: CreateDefaultPrimeShape(),
		Children:   ChildrenMap{},
		Parent:     nil,
	}
}

func (group *Group) GetId() uuid.UUID {
	return group.id
}
func (group *Group) GetParent() Shape {
	return group.Parent
}

func (group *Group) SetParent(shape Shape) {
	group.Parent = shape
}

func (group *Group) GetSavedRay() Ray {
	return group.SavedRay
}
func (group *Group) SetSavedRay(ray Ray) {
	group.SavedRay = ray
}

func (group *Group) AddChild(shape Shape) {
	shape.SetParent(group)
	group.Children[shape.GetId()] = shape
}

func (group *Group) GetTransforms() pm.Matrix4x4 {
	return group.Transforms
}
func (group *Group) SetTransform(transform *pm.Matrix4x4) pm.Matrix4x4 {
	group.Transforms = transform.Multiply(group.Transforms)
	return group.Transforms
}
func (group *Group) SetTransforms(transform []*pm.Matrix4x4) {
	for i := 0; i < len(transform); i++ {
		group.SetTransform(transform[i])
	}
}

func (group *Group) GetMaterial() *mat.Material {
	return &group.Material
}
func (group *Group) SetMaterial(material mat.Material) {
	group.Material = material
}

// Could really be improved
func (group *Group) LocalIntersect(ray Ray) Intersections {
	// Def need to change this
	overAllIntersections := Intersections{Intersections: []Intersection{}}
	for _, shape := range group.Children {
		inters := shape.Intersect(&ray)
		overAllIntersections.Intersections = append(overAllIntersections.Intersections, inters.Intersections...)
	}
	sort.Slice(overAllIntersections.Intersections, func(i, j int) bool {
		return overAllIntersections.Intersections[i].T < overAllIntersections.Intersections[j].T
	})

	return overAllIntersections
}

func CheckGroupAxis(origin, direction, min, max float64) (float64, float64) {
	// 	tminNumerator := (min - origin)
	// 	tmaxNumerator := (max - origin)

	// 	var tmin float64
	// 	var tmax float64

	// 	if math.Abs(direction) >= Epsilon {
	// 		tmin = tminNumerator / direction
	// 		tmax = tmaxNumerator / direction
	// 	} else {
	// 		tmin = tminNumerator * math.Inf(1)
	// 		tmax = tmaxNumerator * math.Inf(1)
	// 	}

	// 	if tmin > tmax {
	// 		minTemp := tmin
	// 		tmin = tmax
	// 		tmax = minTemp
	// 	}

	// 	return tmin, tmax
	// }

	tmin := (min - origin) / direction
	tmax := (max - origin) / direction

	if tmin > tmax {
		return tmax, tmin
	}

	return tmin, tmax
}

func (group *Group) Intersect(ray *Ray) Intersections {
	// Check if bounding box is intersected
	tray := ray.Transform(group.Transforms.Inverse())

	// gbdMin := group.BoundingBox().Minimum
	// gbdMax := group.BoundingBox().Maximum

	// xtmin, xtmax := CheckGroupAxis(ray.origin.X, ray.direction.X, gbdMin.X, gbdMax.X)
	// ytmin, ytmax := CheckGroupAxis(ray.origin.Y, ray.direction.Y, gbdMin.Y, gbdMax.Y)
	// ztmin, ztmax := CheckGroupAxis(ray.origin.Z, ray.direction.Z, gbdMin.Z, gbdMax.Z)

	// tmin := math.Max(math.Max(xtmin, ytmin), ztmin)
	// tmax := math.Min(math.Min(xtmax, ytmax), ztmax)

	// if tmin > tmax {
	// 	return Intersections{Intersections: []Intersection{}}

	// }

	return group.LocalIntersect(tray)
}

func WorldToObject(shape Shape, point pm.Tuple) pm.Tuple {
	if shape.GetParent() != nil {
		point = WorldToObject(shape.GetParent(), point)
	}

	return shape.GetTransforms().Inverse().TupleMultiply(point)
}

func NormalToWorld(shape Shape, normal pm.Tuple) pm.Tuple {
	shapTransform := shape.GetTransforms().Inverse().Transpose()
	normal = shapTransform.TupleMultiply(normal)
	normal.W = 0
	normal = pm.Normalize(normal)

	if shape.GetParent() != nil {
		normal = NormalToWorld(shape.GetParent(), normal)
	}

	return normal
}

func (group *Group) LocalNormalAt(localPoint pm.Tuple, hitPoint *pm.Tuple, intersection *Intersection) pm.Tuple {
	panic("local normal called on group")
}

func NormalAt(shape Shape, point pm.Tuple) pm.Tuple {
	localPoint := WorldToObject(shape, point)
	localNormal := shape.LocalNormalAt(localPoint, nil, nil)
	return NormalToWorld(shape, localNormal)
}

// Create empty bounding box
// For each child
// Create a point for each corner of the bounding box
/*
	For each child transform each of theses points representing each corner of the cube.
	points.append(minimum)
	points.append(.pm.Point(x: minimum.X, y: minimum.Y, z: maximum.Z))
	points.append(.pm.Point(x: minimum.X, y: maximum.Y, z: minimum.Z))
	points.append(.pm.Point(x: minimum.X, y: maximum.Y, z: maximum.Z))
	points.append(.pm.Point(x: maximum.X, y: minimum.Y, z: minimum.Z))
	points.append(.pm.Point(x: maximum.X, y: minimum.Y, z: maximum.Z))
	points.append(.pm.Point(x: maximum.X, y: maximum.Y, z: minimum.Z))
	points.append(maximum)
*/

// get min and max of all childern points as you go

func Bounds(group *Group) *BoundingBox {

	// Multiply the point by the object's transformation matrix
	// Transform all eight of the cube's corner's  and then find a single bounding box that fits them all

	bound := &BoundingBox{
		Minimum: pm.Point(math.Inf(1), math.Inf(1), math.Inf(1)),
		Maximum: pm.Point(math.Inf(-1), math.Inf(-1), math.Inf(-1)),
	}
	for _, v := range group.Children {
		bd := v.BoundingBox()
		points := []pm.Tuple{
			v.GetTransforms().TupleMultiply(bd.Minimum),
			v.GetTransforms().TupleMultiply(pm.Point(bd.Minimum.X, bd.Minimum.Y, bd.Maximum.Z)),
			v.GetTransforms().TupleMultiply(pm.Point(bd.Minimum.X, bd.Maximum.Y, bd.Minimum.Z)),
			v.GetTransforms().TupleMultiply(pm.Point(bd.Minimum.X, bd.Maximum.Y, bd.Maximum.Z)),

			v.GetTransforms().TupleMultiply(pm.Point(bd.Maximum.X, bd.Minimum.Y, bd.Minimum.Z)),
			v.GetTransforms().TupleMultiply(pm.Point(bd.Maximum.X, bd.Minimum.Y, bd.Maximum.Z)),
			v.GetTransforms().TupleMultiply(pm.Point(bd.Maximum.X, bd.Maximum.Y, bd.Minimum.Z)),
			v.GetTransforms().TupleMultiply(bd.Maximum),
		}

		for i := 0; i < len(points); i++ {
			bound.Minimum.X = math.Min(bound.Minimum.X, points[i].X)
			bound.Minimum.Y = math.Min(bound.Minimum.Y, points[i].Y)
			bound.Minimum.Z = math.Min(bound.Minimum.Z, points[i].Z)

			bound.Maximum.X = math.Max(bound.Maximum.X, points[i].X)
			bound.Maximum.Y = math.Max(bound.Maximum.Y, points[i].Y)
			bound.Maximum.Z = math.Max(bound.Maximum.Z, points[i].Z)
		}

	}

	return bound
}

func (group *Group) BoundingBox() *BoundingBox {
	return Bounds(group)
}
