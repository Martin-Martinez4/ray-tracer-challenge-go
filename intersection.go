package main

import (
	"math"
	"sort"
)

type Intersection struct {
	T float64
	S *Sphere
}

type Intersections struct {
	intersections []Intersection
}

type Computation struct {
	T       float64
	Object  *Sphere
	Point   Tuple
	Eyev    Tuple
	Normalv Tuple
	Inside  bool
}

func (inters *Intersections) Add(inter Intersection) {

	intersections := append(inters.intersections, inter)
	inters.intersections = intersections

	sort.Slice(inters.intersections, func(i, j int) bool {
		return inters.intersections[i].T < inters.intersections[j].T
	})
}

func (inters *Intersections) RaySphereInteresect(ray Ray, s *Sphere) {

	ray = ray.Transform(s.GetTransforms().Inverse())

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

			inters.Add(Intersection{d1, s})
			inters.Add(Intersection{d2, s})

		} else {

			inters.Add(Intersection{d1, s})

		}
	}
}

func (inters Intersections) Equal(other Intersections) bool {

	oriInters := inters.intersections
	otherInters := other.intersections

	if len(oriInters) != len(otherInters) {
		return false
	}

	for i := 0; i < len(oriInters); i++ {
		if !AreFloatsEqual(oriInters[i].T, otherInters[i].T) || oriInters[i].S != otherInters[i].S {

			return false
		}
	}

	return true
}

func (inters *Intersections) Hit() *Intersection {
	if inters.intersections[0].T < 0 && inters.intersections[len(inters.intersections)-1].T < 0 {
		return nil
	}

	for i := 0; i < len(inters.intersections); i++ {

		if inters.intersections[i].T >= 0 {
			return &inters.intersections[i]
		}

	}

	return nil
}

func Position(r Ray, distance float64) Tuple {
	add := r.direction.SMultiply(distance)
	pos := r.origin.Add(add)
	return pos
}

func RaySphereInteresect(ray Ray, s *Sphere) [](*Intersection) {
	ray = ray.Transform(s.GetTransforms().Inverse())

	sphereToRay := ray.origin.Subtract(Point(0, 0, 0))

	a := Dot(ray.direction, ray.direction)
	b := 2 * Dot(ray.direction, sphereToRay)
	c := Dot(sphereToRay, sphereToRay) - 1

	discriminant := (b * b) - (4*a)*c

	if discriminant < 0 {
		return nil
	} else {
		d1 := (-b - math.Sqrt(discriminant)) / (2 * a)
		d2 := (-b + math.Sqrt(discriminant)) / (2 * a)

		if !AreFloatsEqual(d1, d2) {

			return [](*Intersection){{d1, s}, {d2, s}}

		} else {

			return [](*Intersection){{d1, s}}

		}
	}
}

func prepareComputations(ray Ray, sphere *Sphere, intersection Intersection) Computation {
	comps := Computation{}
	comps.T = intersection.T
	comps.Object = sphere

	comps.Point = Position(ray, intersection.T)
	comps.Eyev = ray.direction.SMultiply(-1)
	comps.Normalv = sphere.NormalAt(comps.Point)

	if Dot(comps.Normalv, comps.Eyev) < 0 {
		comps.Inside = true
		comps.Normalv = comps.Normalv.SMultiply(-1)
	} else {
		comps.Inside = false
	}

	return comps
}

func (comp *Computation) Equal(other Computation) bool {

	return AreFloatsEqual(comp.T, other.T) && comp.Point.Equal(other.Point) && comp.Eyev.Equal(other.Eyev) && comp.Normalv.Equal(other.Normalv)
}
