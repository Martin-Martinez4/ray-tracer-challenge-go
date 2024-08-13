package main

import "testing"

func TestStripeAt(t *testing.T) {
	tests := []struct {
		name    string
		pattern *Stripes
		point   Tuple
		want    Color
	}{
		{
			name:    "a stripe pattern is constant in y",
			pattern: NewStripe(BLACK, WHITE),
			point:   Point(0, 0, 0),
			want:    WHITE,
		},
		{
			name:    "a stripe pattern is constant in y",
			pattern: NewStripe(BLACK, WHITE),
			point:   Point(0, 1, 0),
			want:    WHITE,
		},
		{
			name:    "a stripe pattern is constant in y",
			pattern: NewStripe(BLACK, WHITE),
			point:   Point(0, 10, 0),
			want:    WHITE,
		},
		{
			name:    "a stripe pattern is constant in z",
			pattern: NewStripe(BLACK, WHITE),
			point:   Point(0, 0, 1),
			want:    WHITE,
		},
		{
			name:    "a stripe pattern is constant in z",
			pattern: NewStripe(BLACK, WHITE),
			point:   Point(0, 1, 10),
			want:    WHITE,
		},

		{
			name:    "a stripe pattern changes in x",
			pattern: NewStripe(BLACK, WHITE),
			point:   Point(0, 0, 0),
			want:    WHITE,
		},
		{
			name:    "a stripe pattern changes in x",
			pattern: NewStripe(BLACK, WHITE),
			point:   Point(0.9, 0, 00),
			want:    WHITE,
		},
		{
			name:    "a stripe pattern changes in x",
			pattern: NewStripe(BLACK, WHITE),
			point:   Point(1, 0, 00),
			want:    BLACK,
		},
		{
			name:    "a stripe pattern changes in x",
			pattern: NewStripe(BLACK, WHITE),
			point:   Point(-1, 0, 00),
			want:    BLACK,
		},
		{
			name:    "a stripe pattern changes in x",
			pattern: NewStripe(BLACK, WHITE),
			point:   Point(-0.1, 0, 00),
			want:    BLACK,
		},
		{
			name:    "a stripe pattern changes in x",
			pattern: NewStripe(BLACK, WHITE),
			point:   Point(-1.1, 0, 00),
			want:    WHITE,
		},
		{
			name:    "a stripe pattern changes in x",
			pattern: NewStripe(BLACK, WHITE),
			point:   Point(9.9, 0, 00),
			want:    BLACK,
		},
		{
			name:    "a stripe pattern changes in x",
			pattern: NewStripe(BLACK, WHITE),
			point:   Point(1.5, 0, 0),
			want:    BLACK,
		},
	}

	for i, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			got := tt.pattern.StripeAt(tt.point)

			if !got.Equal(tt.want) {
				t.Errorf("\n%d %s failed:\nwanted: %s\ngot: %s\n", i, tt.name, tt.want.Print(), got.Print())
			}
		})

	}
}
