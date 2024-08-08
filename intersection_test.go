package main

import (
	"reflect"
	"testing"
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
			sphere: &theSphere,
			intersection: Intersections{
				intersections: []Intersection{
					{4, &theSphere},
					{6, &theSphere},
				},
			},
		},
		{
			name:   "Ray intersects, should return two values",
			ray:    NewRay([3]float64{0, 1, -5}, [3]float64{0, 0, 1}),
			sphere: &theSphere,
			intersection: Intersections{
				intersections: []Intersection{
					{5, &theSphere},
				},
			},
		},
		{
			name:   "Ray misses should return an empty array",
			ray:    NewRay([3]float64{0, 2, -5}, [3]float64{0, 0, 1}),
			sphere: &theSphere,
			intersection: Intersections{
				intersections: []Intersection{},
			},
		},
		{
			name:   "Ray intersects, should return two values",
			ray:    NewRay([3]float64{0, 0, 0}, [3]float64{0, 0, 1}),
			sphere: &theSphere,
			intersection: Intersections{
				intersections: []Intersection{
					{-1, &theSphere},
					{1, &theSphere},
				},
			},
		},
		{
			name:   "Ray intersects, should return two negative values",
			ray:    NewRay([3]float64{0, 0, 5}, [3]float64{0, 0, 1}),
			sphere: &theSphere,
			intersection: Intersections{
				intersections: []Intersection{
					{-6, &theSphere},
					{-4, &theSphere},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ints := Intersections{}

			ints.RaySphereInteresect(tt.ray, tt.sphere)

			if !ints.Equal(tt.intersection) {
				t.Errorf("%s did not pass, lengths are not equal: \nGot: %v \nWanted: %v", tt.name, ints.intersections, tt.intersection)
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
			sphere:    &theSphere,
			transform: "scale",
			args:      []float64{2, 2, 2},
			intersection: Intersections{
				intersections: []Intersection{
					{3, &theSphere},
					{7, &theSphere},
				},
			},
		},
		{
			name:      "Ray intersects, should return two values",
			ray:       NewRay([3]float64{0, 0, -5}, [3]float64{0, 0, 1}),
			sphere:    &theSphere,
			transform: "scale",
			args:      []float64{0.5, 0.5, 0.5},
			intersection: Intersections{
				intersections: []Intersection{
					{4.5, &theSphere},
					{5.5, &theSphere},
				},
			},
		},

		{
			name:      "Ray intersects, should return two values",
			ray:       NewRay([3]float64{0, 0, -5}, [3]float64{0, 0, 1}),
			sphere:    &theSphere,
			transform: "translate",
			args:      []float64{5, 0, 0},
			intersection: Intersections{
				intersections: nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if tt.transform == "scale" {

				theSphere.Scale(tt.args[0], tt.args[1], tt.args[2])

			} else if tt.transform == "translate" {
				theSphere.Translate(tt.args[0], tt.args[1], tt.args[2])

			} else {
				t.Errorf("%s is not a vaild transformation option", tt.transform)
				return
			}

			ints := Intersections{}

			ints.RaySphereInteresect(tt.ray, tt.sphere)

			if !ints.Equal(tt.intersection) {
				t.Errorf("%s did not pass, lengths are not equal: \n\nGot: %v \n\nWanted: %v", tt.name, ints.intersections, tt.intersection)
			}

			tt.sphere.Transforms = NewMatrix4x4([16]float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1})

		})
	}
}

func TestHit(t *testing.T) {

	theSphere := NewSphere()

	tests := []struct {
		name          string
		sphere        Sphere
		intersections []Intersection
		want          *Intersection
	}{
		{
			name:          "A hit should be returned",
			sphere:        theSphere,
			intersections: []Intersection{{1, &theSphere}, {2, &theSphere}},
			want:          &Intersection{1, &theSphere},
		},
		{
			name:          "A hit should be returned",
			sphere:        theSphere,
			intersections: []Intersection{{-1, &theSphere}, {1, &theSphere}},
			want:          &Intersection{1, &theSphere},
		},
		{
			name:          "An nil hit should be returned",
			sphere:        theSphere,
			intersections: []Intersection{{-1, &theSphere}, {-2, &theSphere}},
			want:          nil,
		},
		{
			name:          "A hit should be returned",
			sphere:        theSphere,
			intersections: []Intersection{{7, &theSphere}, {7, &theSphere}, {-3, &theSphere}, {2, &theSphere}},
			want:          &Intersection{2, &theSphere},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			inters := Intersections{intersections: []Intersection{}}

			for i := 0; i < len(tt.intersections); i++ {
				inters.Add(tt.intersections[i])
			}

			hit := inters.Hit()

			if (hit == nil && tt.want != nil) || (hit != nil && tt.want == nil) {
				t.Errorf("%s did not pass, \nwanted\n%v\ngot%v", tt.name, tt.want, hit)
				return

			}

			if hit == nil && tt.want == nil {
				return

			} else if hit.T != tt.want.T || !reflect.DeepEqual(hit.S, tt.want.S) {
				t.Errorf("%s did not pass, \nwanted\n%v\ngot%v", tt.name, tt.want, hit)
			}

		})
	}
}

func TestPrepareComputations(t *testing.T) {

	theSphere := NewSphere()

	tests := []struct {
		name         string
		ray          Ray
		sphere       Sphere
		intersection Intersection
		want         Computation
	}{
		{
			name:         "A hit should be returned",
			ray:          NewRay([3]float64{0, 0, -5}, [3]float64{0, 0, 1}),
			sphere:       theSphere,
			intersection: Intersection{4, &theSphere},
			want:         Computation{T: 4, Object: &theSphere},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := Intersections{intersections: []Intersection{}}

			if !true {
				t.Errorf("")
				return

			}

		})
	}
}
