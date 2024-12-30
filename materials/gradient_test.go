package materials

// import (
// 	"testing"

// 	pm "github.com/Martin-Martinez4/ray-tracer-challenge-go/primitive_math"
// 	"github.com/Martin-Martinez4/ray-tracer-challenge-go/shapes"
// )

// func TestGradientPatternAt(t *testing.T) {
// 	tests := []struct {
// 		name              string
// 		pattern           Pattern
// 		shape             shapes.Shape
// 		shapeTransforms   []*pm.Matrix4x4
// 		patternTransforms []*pm.Matrix4x4
// 		point             pm.Tuple
// 		want              Color
// 	}{
// 		{
// 			name:              "pm.Point(0,0,0) should be white",
// 			pattern:           NewGradient(WHITE, BLACK),
// 			shape:             shapes.NewSphere(),
// 			shapeTransforms:   []*pm.Matrix4x4{},
// 			patternTransforms: []*pm.Matrix4x4{},
// 			point:             pm.Point(0, 0, 0),
// 			want:              WHITE,
// 		},
// 		{
// 			name:              "pm.Point(0.25,0,0) should be a bit darker than white",
// 			pattern:           NewGradient(WHITE, BLACK),
// 			shape:             shapes.NewSphere(),
// 			shapeTransforms:   []*pm.Matrix4x4{},
// 			patternTransforms: []*pm.Matrix4x4{},
// 			point:             pm.Point(0.25, 0, 0),
// 			want:              NewColor(0.75, 0.75, 0.75),
// 		},
// 		{
// 			name:              "pm.Point(0.25,0,0) should be gray",
// 			pattern:           NewGradient(WHITE, BLACK),
// 			shape:             shapes.NewSphere(),
// 			shapeTransforms:   []*pm.Matrix4x4{},
// 			patternTransforms: []*pm.Matrix4x4{},
// 			point:             pm.Point(0.5, 0, 0),
// 			want:              NewColor(0.5, 0.5, 0.5),
// 		},
// 	}

// 	for i, tt := range tests {

// 		t.Run(tt.name, func(t *testing.T) {

// 			if len(tt.shapeTransforms) > 0 {

// 				tt.shape.SetTransforms(tt.shapeTransforms)
// 			}

// 			if len(tt.patternTransforms) > 0 {
// 				tt.pattern.SetTransforms(tt.patternTransforms)
// 			}

// 			got := tt.pattern.PatternAtShape(tt.shape, tt.point)

// 			if !got.Equal(tt.want) {
// 				t.Errorf("\n%d %s failed:\nwanted: %s\ngot: %s\n", i, tt.name, tt.want.Print(), got.Print())
// 			}
// 		})

// 	}
// }
