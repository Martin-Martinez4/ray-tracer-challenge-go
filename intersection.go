package main

import (
	"fmt"
	"math"
	"sort"
)

type Intersection struct {
	T float64
	S Shape
}

type Intersections struct {
	intersections []Intersection
}

type Computations struct {
	T          float64
	Object     Shape
	Point      Tuple
	Eyev       Tuple
	Normalv    Tuple
	OverPoint  Tuple
	UnderPoint Tuple
	ReflectV   Tuple
	Inside     bool
	N1         float64
	N2         float64
}

func (comps Computations) Print() string {
	return fmt.Sprintf("\nT: %f\nPoint: %s\nEyeV: %s\nNormalV: %s\nOverPoint: %s\nUnderPoint: %s\nN1: %f,\nN2: %f", comps.T, comps.Point.Print(), comps.Eyev.Print(), comps.Normalv.Print(), comps.OverPoint.Print(), comps.UnderPoint.Print(), comps.N1, comps.N2)
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

func (inters *Intersections) RayShapeInteresect(ray Ray, s Shape) {

	intersections := s.Intersect(&ray).intersections

	if intersections == nil {
		return
	}

	for _, intersection := range intersections {
		inters.Add(intersection)
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

	if inters == nil || inters.intersections == nil || len(inters.intersections) < 1 {
		return nil
	}

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

func RaySphereInteresect(ray Ray, s *Sphere) *Intersections {
	ray = ray.Transform(s.GetTransforms().Inverse())

	sphereToRay := ray.origin.Subtract(Point(0, 0, 0))

	a := Dot(ray.direction, ray.direction)
	b := 2 * Dot(ray.direction, sphereToRay)
	c := Dot(sphereToRay, sphereToRay) - 1

	discriminant := (b * b) - (4 * a * c)

	if discriminant < 0 {
		return nil
	} else {
		d1 := (-b - math.Sqrt(discriminant)) / (2 * a)
		d2 := (-b + math.Sqrt(discriminant)) / (2 * a)

		intersections := Intersections{}

		if !AreFloatsEqual(d1, d2) {

			intersections.Add(Intersection{d1, s})
			intersections.Add(Intersection{d2, s})

		} else {
			intersections.Add(Intersection{d1, s})

		}
		return &intersections
	}
}

func PrepareComputationsWithHit(i Intersection, r Ray, xs []Intersection) *Computations {
	comps := PrepareComputations(r, i.S, i)

	containers := make([]Shape, 0)
	for _, item := range xs {
		if i == item {
			if len(containers) == 0 {
				comps.N1 = 1.0
			} else {
				comps.N1 = (containers[len(containers)-1]).GetMaterial().RefractiveIndex
			}
		}

		var itemIndex int = -1
		for index := 0; index < len(containers); index++ {
			if containers[index] == item.S {
				itemIndex = index
			}
		}

		if itemIndex != -1 {
			containers = append(containers[:itemIndex], containers[itemIndex+1:]...)
		} else {
			containers = append(containers, item.S)
		}
		if i == item {
			if len(containers) == 0 {
				comps.N2 = 1.0
			} else {
				comps.N2 = (containers[len(containers)-1]).GetMaterial().RefractiveIndex
			}

			break
		}

	}

	return &comps
}

func PrepareComputations(ray Ray, shape Shape, intersection Intersection) Computations {
	comps := Computations{}
	comps.T = intersection.T
	comps.Object = shape

	comps.Point = Position(ray, intersection.T)
	comps.Eyev = ray.direction.SMultiply(-1)
	comps.Normalv = shape.NormalAt(comps.Point)

	if Dot(comps.Normalv, comps.Eyev) < 0 {
		comps.Inside = true
		comps.Normalv = comps.Normalv.SMultiply(-1)
	} else {
		comps.Inside = false
	}

	nvEp := comps.Normalv.SMultiply(Epsilon)
	comps.OverPoint = comps.Point.Add(nvEp)
	comps.UnderPoint = comps.Point.Subtract(nvEp)

	comps.ReflectV = ray.direction.ReflectBy(comps.Normalv)

	return comps
}

func (comp *Computations) Equal(other Computations) bool {

	return AreFloatsEqual(comp.T, other.T) && comp.Point.Equal(other.Point) && comp.Eyev.Equal(other.Eyev) && comp.Normalv.Equal(other.Normalv)
}
