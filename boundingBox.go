package main

import "fmt"

type BoundingBox struct {
	Minimum Tuple
	Maximum Tuple
}

func (boundingBox *BoundingBox) Print() string {
	return fmt.Sprintf("\nMinimum: %s\nMaximum: %s\n", boundingBox.Minimum.Print(), boundingBox.Maximum.Print())
}
