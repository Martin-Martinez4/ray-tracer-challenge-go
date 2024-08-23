package main

import (
	"math"
	"sort"

	"github.com/google/uuid"
)

/*
	type Shape interface {
	GetTransforms() Matrix4x4
	SetTransform(transform *Matrix4x4) Matrix4x4
	SetTransforms(transform []*Matrix4x4)

	GetMaterial() *Material
	SetMaterial(material Material)

	Intersect(ray *Ray) Intersections

	NormalAt(point Tuple) Tuple

	GetSavedRay() Ray
	SetSavedRay(ray Ray)
}
*/

type ChildrenMap map[uuid.UUID]Shape

type Group struct {
	id         uuid.UUID
	Children   ChildrenMap
	Transforms Matrix4x4
	Material   Material
	Parent     Shape
	SavedRay   Ray
}

func NewGroup() *Group {
	id, err := uuid.NewUUID()
	identityMatix := NewMatrix4x4([16]float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1})

	if err != nil {
		panic("not able to cerate unique id for cube")
	}
	return &Group{
		id:         id,
		Children:   ChildrenMap{},
		Transforms: identityMatix,
		Material:   DefaultMaterial(),
		Parent:     nil,
	}
}

func (group *Group) GetId() uuid.UUID {
	return group.id
}
func (group *Group) GetParent() Shape {
	return group.Parent
}

func (group *Group) GetSavedRay() Ray {
	return group.SavedRay
}
func (group *Group) SetSavedRay(ray Ray) {
	group.SavedRay = ray
}

func (group *Group) AddChild(shape Shape) {
	group.Children[shape.GetId()] = shape
}

func (group *Group) GetTransforms() Matrix4x4 {
	return group.Transforms
}
func (group *Group) SetTransform(transform *Matrix4x4) Matrix4x4 {
	group.Transforms = transform.Multiply(group.Transforms)
	return group.Transforms
}
func (group *Group) SetTransforms(transform []*Matrix4x4) {
	for i := 0; i < len(transform); i++ {
		group.SetTransform(transform[i])
	}
}

func (group *Group) GetMaterial() *Material {
	return &group.Material
}
func (group *Group) SetMaterial(material Material) {
	group.Material = material
}

// Could really be improved
func (group *Group) LocalIntersect(ray Ray) Intersections {
	// Def need to change this
	overAllIntersections := Intersections{intersections: []Intersection{}}
	for _, shape := range group.Children {
		inters := shape.Intersect(&ray)
		overAllIntersections.intersections = append(overAllIntersections.intersections, inters.intersections...)
	}
	sort.Slice(overAllIntersections.intersections, func(i, j int) bool {
		return overAllIntersections.intersections[i].T < overAllIntersections.intersections[j].T
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

	// xtmin, xtmax := CheckGroupAxis(ray.origin.x, ray.direction.x, gbdMin.x, gbdMax.x)
	// ytmin, ytmax := CheckGroupAxis(ray.origin.y, ray.direction.y, gbdMin.y, gbdMax.y)
	// ztmin, ztmax := CheckGroupAxis(ray.origin.z, ray.direction.z, gbdMin.z, gbdMax.z)

	// tmin := math.Max(math.Max(xtmin, ytmin), ztmin)
	// tmax := math.Min(math.Min(xtmax, ytmax), ztmax)

	// if tmin > tmax {
	// 	return Intersections{intersections: []Intersection{}}

	// }

	return group.LocalIntersect(tray)
}

func WorldToObject(shape Shape, point Tuple) Tuple {
	if shape.GetParent() != nil {
		point = WorldToObject(shape.GetParent(), point)
	}

	return shape.GetTransforms().Inverse().TupleMultiply(point)
}

func NormalToWorld(shape Shape, normal Tuple) Tuple {
	shapTransform := shape.GetTransforms().Inverse().Transpose()
	normal = shapTransform.TupleMultiply(normal)
	normal.w = 0
	normal = Normalize(normal)

	if shape.GetParent() != nil {
		normal = NormalToWorld(shape.GetParent(), normal)
	}

	return normal
}

func (group *Group) LocalNormalAt(localPoint Tuple) Tuple {
	panic("local normal called on group")
}

func NormalAt(shape Shape, point Tuple) Tuple {
	localPoint := WorldToObject(shape, point)
	localNormal := shape.LocalNormalAt(localPoint)
	return NormalToWorld(shape, localNormal)
}

// Create empty bounding box
// For each child
// Create a point for each corner of the bounding box
/*
	For each child transform each of theses points representing each corner of the cube.
	points.append(minimum)
	points.append(.Point(x: minimum.x, y: minimum.y, z: maximum.z))
	points.append(.Point(x: minimum.x, y: maximum.y, z: minimum.z))
	points.append(.Point(x: minimum.x, y: maximum.y, z: maximum.z))
	points.append(.Point(x: maximum.x, y: minimum.y, z: minimum.z))
	points.append(.Point(x: maximum.x, y: minimum.y, z: maximum.z))
	points.append(.Point(x: maximum.x, y: maximum.y, z: minimum.z))
	points.append(maximum)
*/

// get min and max of all childern points as you go

func Bounds(group *Group) *BoundingBox {

	// Multiply the point by the object's transformation matrix
	// Transform all eight of the cube's corner's  and then find a single bounding box that fits them all

	bound := &BoundingBox{
		Minimum: Point(math.Inf(1), math.Inf(1), math.Inf(1)),
		Maximum: Point(math.Inf(-1), math.Inf(-1), math.Inf(-1)),
	}
	for _, v := range group.Children {
		bd := v.BoundingBox()
		points := []Tuple{
			v.GetTransforms().TupleMultiply(bd.Minimum),
			v.GetTransforms().TupleMultiply(Point(bd.Minimum.x, bd.Minimum.y, bd.Maximum.z)),
			v.GetTransforms().TupleMultiply(Point(bd.Minimum.x, bd.Maximum.y, bd.Minimum.z)),
			v.GetTransforms().TupleMultiply(Point(bd.Minimum.x, bd.Maximum.y, bd.Maximum.z)),

			v.GetTransforms().TupleMultiply(Point(bd.Maximum.x, bd.Minimum.y, bd.Minimum.z)),
			v.GetTransforms().TupleMultiply(Point(bd.Maximum.x, bd.Minimum.y, bd.Maximum.z)),
			v.GetTransforms().TupleMultiply(Point(bd.Maximum.x, bd.Maximum.y, bd.Minimum.z)),
			v.GetTransforms().TupleMultiply(bd.Maximum),
		}

		for i := 0; i < len(points); i++ {
			bound.Minimum.x = math.Min(bound.Minimum.x, points[i].x)
			bound.Minimum.y = math.Min(bound.Minimum.y, points[i].y)
			bound.Minimum.z = math.Min(bound.Minimum.z, points[i].z)

			bound.Maximum.x = math.Max(bound.Maximum.x, points[i].x)
			bound.Maximum.y = math.Max(bound.Maximum.y, points[i].y)
			bound.Maximum.z = math.Max(bound.Maximum.z, points[i].z)
		}

	}

	return bound
}

func (group *Group) BoundingBox() *BoundingBox {
	return Bounds(group)
}
