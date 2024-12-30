package shapes

import (
	"fmt"

	pm "github.com/Martin-Martinez4/ray-tracer-challenge-go/primitive_math"
)

type BoundingBox struct {
	Minimum pm.Tuple
	Maximum pm.Tuple
}

func (boundingBox *BoundingBox) Print() string {
	return fmt.Sprintf("\nMinimum: %s\nMaximum: %s\n", boundingBox.Minimum.Print(), boundingBox.Maximum.Print())
}
