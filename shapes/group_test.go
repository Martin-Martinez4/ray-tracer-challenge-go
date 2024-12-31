package shapes

import (
	"math"
	"testing"

	pm "github.com/Martin-Martinez4/ray-tracer-challenge-go/primitive_math"
)

func TestIntersectingAnEmptyGroup(t *testing.T) {

	group := NewGroup()

	name := "the ray hits an empty group"
	origin := pm.Point(0, 0, 0)
	direction := pm.Vector(0, 0, 1)
	want := []Intersection{}

	normed := pm.Normalize(direction)

	ray := NewRay([3]float64{origin.X, origin.Y, origin.Z}, [3]float64{normed.X, normed.Y, normed.Z})

	got := group.LocalIntersect(ray)

	if !got.Equal(Intersections{Intersections: want}) {
		t.Errorf("%s: \nwant: %v \ngot: %v \ndo not match", name, want, got)
	}

}

func TestIntersectingAGroup(t *testing.T) {

	group := NewGroup()

	name := "the ray hits a group with two shapes"
	origin := pm.Point(0, 0, -5)
	direction := pm.Vector(0, 0, 1)

	s1 := NewSphere()

	s2 := NewSphere()
	s2.SetTransform(pm.Translate(0, 0, -3))

	s3 := NewSphere()
	s3.SetTransform(pm.Translate(5, 0, 0))

	group.AddChild(s1)
	group.AddChild(s2)
	group.AddChild(s3)

	want := []Intersection{NewIntersection(1, s2), NewIntersection(3, s2), NewIntersection(4, s1), NewIntersection(6, s1)}

	normed := pm.Normalize(direction)

	ray := NewRay([3]float64{origin.X, origin.Y, origin.Z}, [3]float64{normed.X, normed.Y, normed.Z})

	got := group.LocalIntersect(ray)

	if !got.Equal(Intersections{Intersections: want}) {
		t.Errorf("%s: \nwant: %v \ngot: %v \ndo not match", name, want, got)
	}

}

func TestIntersectingATransfomredGroup(t *testing.T) {

	group := NewGroup()

	group.SetTransform(pm.Scale(2, 2, 2))

	name := "the ray hits a group with transforms"
	origin := pm.Point(10, 0, -10)
	direction := pm.Vector(0, 0, 1)

	s1 := NewSphere()
	s1.SetTransform(pm.Translate(5, 0, 0))

	group.AddChild(s1)

	want := []Intersection{NewIntersection(8, s1), NewIntersection(12, s1)}

	normed := pm.Normalize(direction)

	ray := NewRay([3]float64{origin.X, origin.Y, origin.Z}, [3]float64{normed.X, normed.Y, normed.Z})

	got := group.Intersect(&ray)

	if !got.Equal(Intersections{Intersections: want}) {
		t.Errorf("%s: \nwant: %v \ngot: %v \ndo not match", name, want, got)
	}

}

func TestObjectToWorld(t *testing.T) {

	group1 := NewGroup()
	group1.SetTransform(pm.RotationAlongY(math.Pi / 2))

	group2 := NewGroup()
	group2.SetTransform(pm.Scale(2, 2, 2))

	group1.AddChild(group2)

	sphere := NewSphere()
	sphere.SetTransform(pm.Translate(5, 0, 0))

	group2.AddChild(sphere)

	name := "the object to world"

	want := pm.Vector(0.2857, 0.4286, -0.8571)

	got := NormalToWorld(sphere, pm.Vector(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3))

	if !got.Equal(got) {
		t.Errorf("%s: \nwant: %v \ngot: %v \ndo not match", name, want, got)
	}

}

func TestNomralOnObjectinGroup(t *testing.T) {

	group1 := NewGroup()
	group1.SetTransform(pm.RotationAlongY(math.Pi / 2))

	group2 := NewGroup()
	group2.SetTransform(pm.Scale(1, 2, 3))

	group1.AddChild(group2)

	sphere := NewSphere()
	sphere.SetTransform(pm.Translate(5, 0, 0))

	group2.AddChild(sphere)

	name := "the object to world"

	want := pm.Vector(0.2857, 0.4286, -0.8571)

	got := NormalAt(sphere, pm.Point(1.7321, 1.1547, -5.5774))

	if !got.Equal(got) {
		t.Errorf("%s: \nwant: %v \ngot: %v \ndo not match", name, want, got)
	}

}

func TestBoundingBoxOnGroup(t *testing.T) {

	group1 := NewGroup()

	leftUp := NewCube()
	leftUp.SetTransform(pm.Translate(-1, 1, 0))

	leftDown := NewCube()
	leftDown.SetTransform(pm.Translate(-1, -1, 0))

	rightUp := NewCube()
	rightUp.SetTransform(pm.Translate(1, 1, 0))

	rightDown := NewCube()
	rightDown.SetTransform(pm.Translate(1, -1, 0))

	group1.AddChild(leftUp)
	group1.AddChild(leftDown)
	group1.AddChild(rightUp)
	group1.AddChild(rightDown)

	name := "bounding box of group of transformed shapes"

	got := Bounds(group1)

	if !got.Maximum.Equal(pm.Point(2, 2, 1)) || !got.Minimum.Equal(pm.Point(-2, -2, -1)) {
		t.Errorf("%s Failed \ngot: %s", name, got.Print())
	}

}
