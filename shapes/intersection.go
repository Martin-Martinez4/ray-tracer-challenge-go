package shapes

import (
	"math"
	"sort"

	pm "github.com/Martin-Martinez4/ray-tracer-challenge-go/primitive_math"
)

type Intersection struct {
	T float64
	S Shape

	U *float64
	V *float64
}

type Intersections struct {
	Intersections []Intersection
}

type Computations struct {
	T          float64
	Object     Shape
	Point      pm.Tuple
	Eyev       pm.Tuple
	Normalv    pm.Tuple
	OverPoint  pm.Tuple
	UnderPoint pm.Tuple
	ReflectV   pm.Tuple
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

	Intersections := append(inters.Intersections, inter)
	inters.Intersections = Intersections

	sort.Slice(inters.Intersections, func(i, j int) bool {
		return inters.Intersections[i].T < inters.Intersections[j].T
	})
	// remove default values intersection {0 <nil> <nil> <nil>} if present
	if inters.Intersections[0].S == nil {
		inters.Intersections = append(inters.Intersections[:0], inters.Intersections[1:]...)
	}
}

func (inters *Intersections) RaySphereInteresect(ray Ray, s *Sphere) {

	ray = ray.Transform(s.GetTransforms().Inverse())

	sphereToRay := ray.origin.Subtract(pm.Point(0, 0, 0))

	a := pm.Dot(ray.Direction, ray.Direction)
	b := 2 * pm.Dot(ray.Direction, sphereToRay)
	c := pm.Dot(sphereToRay, sphereToRay) - 1

	discriminant := (b * b) - (4*a)*c

	if discriminant < 0 {

	} else {
		d1 := (-b - math.Sqrt(discriminant)) / (2 * a)
		d2 := (-b + math.Sqrt(discriminant)) / (2 * a)

		if !pm.AreFloatsEqual(d1, d2) {

			inters.Add(NewIntersection(d1, s))
			inters.Add(NewIntersection(d2, s))

		} else {

			inters.Add(NewIntersection(d1, s))

		}
	}
}

func (inters *Intersections) RayShapeInteresect(ray Ray, s Shape) {

	Intersections := s.Intersect(&ray).Intersections

	if Intersections == nil {
		return
	}

	for _, intersection := range Intersections {
		inters.Add(intersection)
	}
}

func (inters Intersections) Equal(other Intersections) bool {

	oriInters := inters.Intersections
	otherInters := other.Intersections

	if len(oriInters) != len(otherInters) {
		return false
	}

	for i := 0; i < len(oriInters); i++ {
		if !pm.AreFloatsEqual(oriInters[i].T, otherInters[i].T) || oriInters[i].S != otherInters[i].S {

			return false
		}
	}

	return true
}

func (inters *Intersections) Hit() *Intersection {

	if inters == nil || inters.Intersections == nil || len(inters.Intersections) < 1 {
		return nil
	}

	if inters.Intersections[0].T < 0 && inters.Intersections[len(inters.Intersections)-1].T < 0 {
		return nil
	}

	for i := 0; i < len(inters.Intersections); i++ {

		if inters.Intersections[i].T >= 0 {
			return &inters.Intersections[i]
		}

	}

	return nil
}

func Position(r Ray, distance float64) pm.Tuple {
	add := r.Direction.SMultiply(distance)
	pos := r.origin.Add(add)
	return pos
}

func RaySphereInteresect(ray Ray, s *Sphere) *Intersections {
	ray = ray.Transform(s.GetTransforms().Inverse())

	sphereToRay := ray.origin.Subtract(pm.Point(0, 0, 0))

	a := pm.Dot(ray.Direction, ray.Direction)
	b := 2 * pm.Dot(ray.Direction, sphereToRay)
	c := pm.Dot(sphereToRay, sphereToRay) - 1

	discriminant := (b * b) - (4 * a * c)

	if discriminant < 0 {
		return nil
	} else {
		d1 := (-b - math.Sqrt(discriminant)) / (2 * a)
		d2 := (-b + math.Sqrt(discriminant)) / (2 * a)

		if !pm.AreFloatsEqual(d1, d2) {

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
	comps.Eyev = ray.Direction.Negate()
	// comps.Normalv = NormalAt(shape, comps.pm.Point)
	comps.Normalv = shape.LocalNormalAt(comps.Point, &comps.Point, &intersection)

	if pm.Dot(comps.Normalv, comps.Eyev) < 0 {
		comps.Inside = true
		comps.Normalv = comps.Normalv.Negate()
	} else {
		comps.Inside = false
	}

	nvEp := comps.Normalv.SMultiply(pm.Epsilon)

	comps.ReflectV = ray.Direction.ReflectBy(comps.Normalv)

	comps.OverPoint = comps.Point.Add(nvEp)
	comps.UnderPoint = comps.Point.Subtract(nvEp)

	return comps
}

func (comp *Computations) Equal(other Computations) bool {

	return pm.AreFloatsEqual(comp.T, other.T) && comp.Point.Equal(other.Point) && comp.Eyev.Equal(other.Eyev) && comp.Normalv.Equal(other.Normalv)
}
