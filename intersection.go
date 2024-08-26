package main

import (
	"math"
	"sort"
)

type Intersection struct {
	T float64
	S Shape

	U *float64
	V *float64
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

func NewIntersection(T float64, S Shape) Intersection {
	return Intersection{
		T: T,
		S: S,
		U: nil,
		V: nil,
	}
}

func NewIntersectionWithUV(T float64, S Shape, U, V *float64) Intersection {
	return Intersection{
		T: T,
		S: S,
		U: U,
		V: V,
	}
}

func (inters *Intersections) Add(inter Intersection) {

	intersections := append(inters.intersections, inter)
	inters.intersections = intersections

	sort.Slice(inters.intersections, func(i, j int) bool {
		return inters.intersections[i].T < inters.intersections[j].T
	})
	// remove default values intersection {0 <nil> <nil> <nil>} if present
	if inters.intersections[0].S == nil {
		inters.intersections = append(inters.intersections[:0], inters.intersections[1:]...)
	}
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

			inters.Add(NewIntersection(d1, s))
			inters.Add(NewIntersection(d2, s))

		} else {

			inters.Add(NewIntersection(d1, s))

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

		if !AreFloatsEqual(d1, d2) {

			return &Intersections{[]Intersection{NewIntersection(d1, s), NewIntersection(d2, s)}}

		} else {

			return &Intersections{[]Intersection{NewIntersection(d1, s)}}

		}
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
	comps.Eyev = ray.direction.Negate()
	// comps.Normalv = NormalAt(shape, comps.Point)
	comps.Normalv = shape.LocalNormalAt(comps.Point, &comps.Point, &intersection)

	if Dot(comps.Normalv, comps.Eyev) < 0 {
		comps.Inside = true
		comps.Normalv = comps.Normalv.Negate()
	} else {
		comps.Inside = false
	}

	nvEp := comps.Normalv.SMultiply(Epsilon)

	comps.ReflectV = ray.direction.ReflectBy(comps.Normalv)

	comps.OverPoint = comps.Point.Add(nvEp)
	comps.UnderPoint = comps.Point.Subtract(nvEp)

	return comps
}

func (comp *Computations) Equal(other Computations) bool {

	return AreFloatsEqual(comp.T, other.T) && comp.Point.Equal(other.Point) && comp.Eyev.Equal(other.Eyev) && comp.Normalv.Equal(other.Normalv)
}
