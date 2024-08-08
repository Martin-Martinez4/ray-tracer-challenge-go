package main

import (
	"fmt"
	strings "strings"
	"testing"
)

func TestNewCanvas(t *testing.T) {
	tests := []struct {
		name   string
		width  int32
		height int32
	}{
		{
			name:   "create canvas 2 width 2 height",
			width:  2,
			height: 2,
		},
		{
			name:   "create canvas 3 width 2 height",
			width:  3,
			height: 2,
		},
		{
			name:   "create canvas 3 width 4 height",
			width:  3,
			height: 4,
		},
	}

	black := NewColor(0, 0, 0)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewCanvas(tt.width, tt.height)

			if len(got.canvas) != int(tt.height) {
				t.Errorf("canvas height not correct")

			}
			if len(got.canvas[0]) != int(tt.width) {
				t.Errorf("canvas width not correct")

			}

			for w := int32(0); w < tt.width; w++ {
				for h := int32(0); h < tt.height; h++ {
					if !got.canvas[h][w].Equal(black) {
						t.Errorf("not all canvas elements were Color(0,0,0)")
					}
				}
			}
		})
	}
}

func TestColorPixel(t *testing.T) {

	canvas1 := NewCanvas(2, 3)
	red := NewColor(1, 0, 0)
	blue := NewColor(0, 0, 1)
	black := NewColor(0, 0, 0)

	tests := []struct {
		name   string
		canvas Canvas
		y      int32
		x      int32
		color  Color
	}{
		{
			name:   "(0,0) should be set to red (1,0,0)",
			canvas: canvas1,
			y:      0,
			x:      0,
			color:  red,
		},
		{
			name:   "(1,2) should be set to blue (0,0,1)",
			canvas: canvas1,
			y:      2,
			x:      1,
			color:  blue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.canvas.ColorPixel(tt.x, tt.y, tt.color)

			if !tt.canvas.canvas[tt.y][tt.x].Equal(tt.color) {
				t.Errorf("canvas was not set to correct color")

			}

			for w := int32(0); w < tt.canvas.width; w++ {
				for h := int32(0); h < tt.canvas.height; h++ {
					if !tt.canvas.canvas[h][w].Equal(black) && h != tt.y && w != tt.x {
						t.Errorf("color at (%v, %v) was not black when it should be", w, h)
					}
				}
			}

			tt.canvas.ColorPixel(tt.x, tt.y, black)
		})
	}
}

func TestPpmheader(t *testing.T) {

	canvas1 := NewCanvas(2, 3)
	canvas2 := NewCanvas(5, 3)

	tests := []struct {
		name   string
		canvas Canvas
		want   string
	}{
		{
			name:   "ppm header for canvas of width 2 height 3",
			canvas: canvas1,
			want:   fmt.Sprintf("P3\n%d %d\n%d", 2, 3, 255),
		},
		{
			name:   "ppm header for canvas of width 5 height 3",
			canvas: canvas2,
			want:   fmt.Sprintf("P3\n%d %d\n%d", 5, 3, 255),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.canvas.ppmheader()

			if strings.Compare(got, tt.want) != 0 {
				t.Errorf("wanted: %s \n got: %s", got, tt.want)
			}

		})
	}
}

func TestPpmbody(t *testing.T) {

	type colorchanges struct {
		x     int32
		y     int32
		color Color
	}

	tests := []struct {
		name         string
		width        int32
		height       int32
		colorchanges []colorchanges
		want         string
	}{
		{
			name:         "print ppm body data no color changes width 2 height 3",
			width:        2,
			height:       3,
			colorchanges: []colorchanges{},
			want:         strings.TrimSpace("0 0 0 0 0 0 \n0 0 0 0 0 0 \n0 0 0 0 0 0"),
		},
		{
			name:         "print ppm body data no color changes width 5 height 3",
			width:        5,
			height:       3,
			colorchanges: []colorchanges{},
			want:         strings.TrimSpace("0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 \n0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 \n0 0 0 0 0 0 0 0 0 0 0 0 0 0 0"),
		},
		{
			name:         "print ppm body data no color changes width 5 height 3 pixel (0,0) should be green",
			width:        5,
			height:       3,
			colorchanges: []colorchanges{{0, 0, NewColor(0, 1, 0)}},
			want:         strings.TrimSpace("0 255 0 0 0 0 0 0 0 0 0 0 0 0 0 \n0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 \n0 0 0 0 0 0 0 0 0 0 0 0 0 0 0"),
		},
		{
			name:         "print ppm body data no color changes width 5 height 3 pixel (0,0) should be green (3,1) should be red",
			width:        5,
			height:       3,
			colorchanges: []colorchanges{{0, 0, NewColor(0, 1, 0)}, {3, 1, NewColor(1, 0, 0)}},
			want:         strings.TrimSpace("0 255 0 0 0 0 0 0 0 0 0 0 0 0 0 \n0 0 0 0 0 0 0 0 0 255 0 0 0 0 0 \n0 0 0 0 0 0 0 0 0 0 0 0 0 0 0"),
		},
		{
			name:         "print ppm body data no color changes width 5 height 3 pixel (0,0) should be (0, 255, 0), values should be clamped",
			width:        5,
			height:       3,
			colorchanges: []colorchanges{{0, 0, NewColor(-0.5, 5, -10)}},
			want:         strings.TrimSpace("0 255 0 0 0 0 0 0 0 0 0 0 0 0 0 \n0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 \n0 0 0 0 0 0 0 0 0 0 0 0 0 0 0"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			canvas := NewCanvas(tt.width, tt.height)

			for colorchange := 0; colorchange < len(tt.colorchanges); colorchange++ {
				canvas.ColorPixel(tt.colorchanges[colorchange].x, tt.colorchanges[colorchange].y, tt.colorchanges[colorchange].color)
			}

			got := canvas.ppmbody()

			if strings.Compare(got, tt.want) != 0 {
				t.Errorf("\nwanted: \n%s \ngot:\n%s", tt.want, got)
			}

		})
	}
}

func TestPpmbodylength(t *testing.T) {

	// type colorchanges struct {
	// 	x     int32
	// 	y     int32
	// 	color Color
	// }

	tests := []struct {
		name   string
		width  int32
		height int32
		// colorchanges []colorchanges
		want string
	}{
		{
			name:   "No line should be more than 70 characters long in the ppm body",
			width:  30,
			height: 2,
			want:   strings.TrimSpace("0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 \n0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 \n0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 \n0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 \n0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 \n0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			canvas := NewCanvas(tt.width, tt.height)

			// for colorchange := 0; colorchange < len(tt.colorchanges); colorchange++ {
			// 	canvas.ColorPixel(tt.colorchanges[colorchange].x, tt.colorchanges[colorchange].y, tt.colorchanges[colorchange].color)
			// }

			got := canvas.ppmbody()

			if strings.Compare(got, tt.want) != 0 {
				t.Errorf("\nwanted: \n%s \ngot:\n%s", tt.want, got)
			}

		})
	}
}

func TestPpm(t *testing.T) {

	// type colorchanges struct {
	// 	x     int32
	// 	y     int32
	// 	color Color
	// }

	tests := []struct {
		name   string
		width  int32
		height int32
		// colorchanges []colorchanges
		want string
	}{
		{
			name:   "No line should be more than 70 characters long in the ppm body",
			width:  30,
			height: 2,
			want:   strings.TrimSpace("0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 \n0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 \n0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 \n0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 \n0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 \n0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			canvas := NewCanvas(tt.width, tt.height)

			// for colorchange := 0; colorchange < len(tt.colorchanges); colorchange++ {
			// 	canvas.ColorPixel(tt.colorchanges[colorchange].x, tt.colorchanges[colorchange].y, tt.colorchanges[colorchange].color)
			// }

			got := canvas.ppmbody()

			if strings.Compare(got, tt.want) != 0 {
				t.Errorf("\nwanted: \n%s \ngot:\n%s", tt.want, got)
			}

		})
	}
}
