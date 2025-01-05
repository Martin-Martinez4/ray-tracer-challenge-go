package shapes

import (
	"fmt"
	"testing"

	pm "github.com/Martin-Martinez4/ray-tracer-challenge-go/primitive_math"
)

func TestRayIntersect(t *testing.T) {

	theSphere := NewSphere()

	tests := []struct {
		name         string
		ray          Ray
		sphere       *Sphere
		intersection Intersections
	}{
		{
			name:   "Ray intersects, should return two values",
			ray:    NewRay([3]float64{0, 0, -5}, [3]float64{0, 0, 1}),
			sphere: theSphere,
			intersection: Intersections{
				NewIntersection(4, theSphere),
				NewIntersection(6, theSphere),
			},
		},
		{
			name:   "Ray intersects, should return two values",
			ray:    NewRay([3]float64{0, 1, -5}, [3]float64{0, 0, 1}),
			sphere: theSphere,
			intersection: Intersections{
				NewIntersection(5, theSphere),
			},
		},
		{
			name:         "Ray misses should return an empty array",
			ray:          NewRay([3]float64{0, 2, -5}, [3]float64{0, 0, 1}),
			sphere:       theSphere,
			intersection: Intersections{},
		},
		{
			name:   "Ray intersects, should return two values",
			ray:    NewRay([3]float64{0, 0, 0}, [3]float64{0, 0, 1}),
			sphere: theSphere,
			intersection: Intersections{
				NewIntersection(-1, theSphere),
				NewIntersection(1, theSphere),
			},
		},
		{
			name:   "Ray intersects, should return two negative values",
			ray:    NewRay([3]float64{0, 0, 5}, [3]float64{0, 0, 1}),
			sphere: theSphere,
			intersection: Intersections{
				NewIntersection(-6, theSphere),
				NewIntersection(-4, theSphere),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ints := Intersections{}

			ints.RaySphereInteresect(tt.ray, tt.sphere)

			if !ints.Equal(tt.intersection) {
				t.Errorf("%s did not pass, lengths are not equal: \nGot: %v \nWanted: %v", tt.name, ints, tt.intersection)
			}

		})
	}
}

func TestRayIntersectWithTransform(t *testing.T) {

	theSphere := NewSphere()

	tests := []struct {
		name         string
		ray          Ray
		sphere       *Sphere
		transform    string
		args         []float64
		want         []float64
		intersection Intersections
	}{
		{
			name:      "Ray intersects, should return two values",
			ray:       NewRay([3]float64{0, 0, -5}, [3]float64{0, 0, 1}),
			sphere:    theSphere,
			transform: "scale",
			args:      []float64{2, 2, 2},
			intersection: Intersections{
				NewIntersection(3, theSphere),
				NewIntersection(7, theSphere),
			},
		},
		{
			name:      "Ray intersects, should return two values",
			ray:       NewRay([3]float64{0, 0, -5}, [3]float64{0, 0, 1}),
			sphere:    theSphere,
			transform: "scale",
			args:      []float64{0.5, 0.5, 0.5},
			intersection: Intersections{
				NewIntersection(4.5, theSphere),
				NewIntersection(5.5, theSphere),
			},
		},

		{
			name:         "Ray intersects, should return two values",
			ray:          NewRay([3]float64{0, 0, -5}, [3]float64{0, 0, 1}),
			sphere:       theSphere,
			transform:    "translate",
			args:         []float64{5, 0, 0},
			intersection: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if tt.transform == "scale" {

				theSphere.SetTransform(pm.Scale(tt.args[0], tt.args[1], tt.args[2]))

			} else if tt.transform == "translate" {
				theSphere.SetTransform(pm.Translate(tt.args[0], tt.args[1], tt.args[2]))

			} else {
				t.Errorf("%s is not a vaild transformation option", tt.transform)
				return
			}

			ints := Intersections{}

			ints.RaySphereInteresect(tt.ray, tt.sphere)

			if !ints.Equal(tt.intersection) {
				t.Errorf("%s did not pass, lengths are not equal: \n\nGot: %v \n\nWanted: %v", tt.name, ints, tt.intersection)
			}

			tt.sphere.Transforms = pm.NewMatrix4x4([16]float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1})

		})
	}
}

// func TestHit(t *testing.T) {

// 	theSphere := NewSphere()

// 	tests := []struct {
// 		name          string
// 		sphere        *Sphere
// 		Intersections []Intersection
// 		want          Intersection
// 	}{
// 		{
// 			name:          "A hit should be returned",
// 			sphere:        theSphere,
// 			Intersections: []Intersection{NewIntersection(1, theSphere), NewIntersection(2, theSphere)},
// 			want:          NewIntersection(1, theSphere),
// 		},
// 		{
// 			name:          "A hit should be returned",
// 			sphere:        theSphere,
// 			Intersections: []Intersection{NewIntersection(-1, theSphere), NewIntersection(1, theSphere)},
// 			want:          NewIntersection(1, theSphere),
// 		},
// 		// {
// 		// 	name:          "An nil hit should be returned",
// 		// 	sphere:        theSphere,
// 		// 	Intersections: []Intersection{NewIntersection(-1, theSphere), NewIntersection(-2, theSphere)},
// 		// 	want:          nil,
// 		// },
// 		{
// 			name:          "A hit should be returned",
// 			sphere:        theSphere,
// 			Intersections: []Intersection{NewIntersection(7, theSphere), NewIntersection(7, theSphere), NewIntersection(-3, theSphere), NewIntersection(2, theSphere)},
// 			want:          NewIntersection(2, theSphere),
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {

// 			inters := Intersections{Intersections: []Intersection{}}

// 			for i := 0; i < len(tt.Intersections); i++ {
// 				inters.Add(tt.Intersections[i])
// 			}

// 			hit := inters.Hit()

// 			if (hit == nil && tt.want != nil) || (hit != nil && tt.want == nil) {
// 				t.Errorf("%s did not pass, \nwanted\n%v\ngot%v", tt.name, tt.want, hit)
// 				return

// 			}

// 			if hit == nil && tt.want == nil {
// 				return

// 			} else if hit.T != tt.want.T || !reflect.DeepEqual(hit.S, tt.want.S) {
// 				t.Errorf("%s did not pass, \nwanted\n%v\ngot%v", tt.name, tt.want, hit)
// 			}

// 		})
// 	}
// }

func TestPrepareComputations(t *testing.T) {

	theSphere := NewSphere()

	tests := []struct {
		name         string
		ray          Ray
		sphere       *Sphere
		intersection Intersection
		want         Computations
	}{
		{
			name:         "computations for a default world, intersection outside",
			ray:          NewRay([3]float64{0, 0, -5}, [3]float64{0, 0, 1}),
			sphere:       theSphere,
			intersection: NewIntersection(4, theSphere),
			want:         Computations{T: 4, Object: theSphere, Point: pm.Point(0, 0, -1), Eyev: pm.Vector(0, 0, -1), Normalv: pm.Vector(0, 0, -1), Inside: false},
		},
		{
			name:         "computations for a default world, intersection inside",
			ray:          NewRay([3]float64{0, 0, 0}, [3]float64{0, 0, 1}),
			sphere:       theSphere,
			intersection: NewIntersection(1, theSphere),
			want:         Computations{T: 1, Object: theSphere, Point: pm.Point(0, 0, 1), Eyev: pm.Vector(0, 0, -1), Normalv: pm.Vector(0, 0, -1), Inside: true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := PrepareComputations(tt.ray, tt.sphere, tt.intersection)

			if !got.Equal(tt.want) {
				t.Errorf("\n%s failed \nwanted %v \ngot %v\n", tt.name, tt.want, got)

			}

		})
	}
}

func TestPrepareComputationsWithinRange(t *testing.T) {

	theSphere := NewSphere()
	theSphere.SetTransform(pm.Translate(0, 0, 1))

	tests := []struct {
		name         string
		ray          Ray
		sphere       *Sphere
		intersection Intersection
		want         Computations
	}{
		{
			name:         "the hit should offset the point",
			ray:          NewRay([3]float64{0, 0, -5}, [3]float64{0, 0, 1}),
			sphere:       theSphere,
			intersection: NewIntersection(5, theSphere),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := PrepareComputations(tt.ray, tt.sphere, tt.intersection)

			if !(got.OverPoint.Z < -pm.Epsilon/2 && got.Point.Z > got.OverPoint.Z) {
				t.Errorf("\n%s failed \nOverPoint value was not within range:\nOverPoint.z < -Epsilon/2 %v \nPoint.z > got.OverPoint.z %v ", tt.name, got.OverPoint.Z < -pm.Epsilon/2, got.Point.Z > got.OverPoint.Z)

			}

		})
	}
}

func TestPrepareComputationsWithHit(t *testing.T) {

	// Setup
	glass1 := NewGlassSphere()
	glass1.GetMaterial().RefractiveIndex = 1.5
	glass1.SetTransform(pm.Scale(2, 2, 2))

	glass2 := NewGlassSphere()
	glass2.GetMaterial().RefractiveIndex = 2.0
	glass2.SetTransform(pm.Translate(0, 0, -0.25))

	glass3 := NewGlassSphere()
	glass3.GetMaterial().RefractiveIndex = 2.5
	glass3.SetTransform(pm.Translate(0, 0, 0.25))

	ray := NewRay([3]float64{0, 0, -4}, [3]float64{0, 0, 1})

	Intersections := []Intersection{
		{T: 2, S: glass1},
		{T: 2.75, S: glass2},
		{T: 3.25, S: glass3},
		{T: 4.75, S: glass2},
		{T: 5.25, S: glass3},
		{T: 6, S: glass1},
	}

	want := []struct {
		n1 float64
		n2 float64
	}{
		{1.0, 1.5},
		{1.5, 2},
		{2, 2.5},
		{2.5, 2.5},
		{2.5, 1.5},
		{1.5, 1.0},
	}

	for i := range Intersections {
		t.Run(fmt.Sprint(i), func(t *testing.T) {

			got := PrepareComputationsWithHit(Intersections[i], ray, Intersections)

			if got.N1 != want[i].n1 || got.N2 != want[i].n2 {
				t.Errorf("#%d failed\nwanted: %f, %f\ngot: %f, %f", i, want[i].n1, want[i].n2, got.N1, got.N2)

			}

		})
	}
}

func TestPrepareComputationsUnderPoint(t *testing.T) {

	tests := []struct {
		name         string
		ray          Ray
		sphere       *Sphere
		intersection Intersection
	}{
		{
			name:         "test underpoint",
			ray:          NewRay([3]float64{0, 0, -5}, [3]float64{0, 0, 1}),
			sphere:       NewGlassSphere(),
			intersection: NewIntersection(5, nil),
		},
	}

	for i, tt := range tests {

		t.Run(fmt.Sprint(i), func(t *testing.T) {

			tt.intersection.S = tt.sphere

			got := PrepareComputationsWithHit(tt.intersection, tt.ray, []Intersection{tt.intersection})

			if got.UnderPoint.Z > pm.Epsilon/2 && got.Point.Z < got.UnderPoint.Z {
				t.Errorf("Failed")

			}

		})
	}
}
