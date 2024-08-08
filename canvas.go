package main

import (
	"fmt"
	"strings"
)

type Canvas struct {
	width  int32
	height int32
	canvas [][]Color
}

func NewCanvas(width, height int32) Canvas {

	c := make([][]Color, height)

	for h := int32(0); h < height; h++ {
		c[h] = make([]Color, width)
		for w := int32(0); w < width; w++ {
			c[h][w] = NewColor(0, 0, 0)
		}
	}

	return Canvas{width, height, c}

}

func (c *Canvas) ColorPixel(x int32, y int32, color Color) {
	c.canvas[y][x] = color

}

func (c *Canvas) GetPixel(x int32, y int32) Color {
	return c.canvas[y][x]
}

func (c *Canvas) ppmheader() string {
	return fmt.Sprintf("P3\n%d %d\n%d", c.width, c.height, 255)
}

/*
	Go through canvas and print each color value separated by a space
	If width is 5 there will be 15 values per line for example
	Values of 1 or higher will be clamped to 255 and lower clamped to 0

	Later no line should be greater than 70 numbers wide add a new line where the space would have been
*/

func (c *Canvas) ppmbody() string {
	// have to go through height first and for each member of that height loop through each color value
	var sb strings.Builder

	for h := int32(0); h < c.height; h++ {

		if h > 0 && h != c.height {
			sb.WriteString("\n")
		}

		for w := int32(0); w < c.width; w++ {
			currentColor := c.canvas[h][w]

			// Every 23 add a new line

			if w%11 == 0 && w != 0 {
				sb.WriteString("\n")
			}
			sb.WriteString(fmt.Sprintf("%d %d %d ",
				clamp(int(currentColor.r*255), 0, 255),
				clamp(int(currentColor.g*255), 0, 255),
				clamp(int(currentColor.b*255), 0, 255)),
			)
		}
	}

	return strings.TrimSpace(sb.String())
}

func (c *Canvas) Newppm() string {
	var sb strings.Builder

	sb.WriteString(c.ppmheader())
	sb.WriteString("\n")
	sb.WriteString(c.ppmbody())
	sb.WriteString("\n")

	return strings.TrimSpace(sb.String())

}

func clamp(value int, min, max int) int {
	if value > max {
		return max
	} else if value < min {
		return min
	}

	return value
}
