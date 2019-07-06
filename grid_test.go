package grid

import (
	"testing"
)

func TestSize(t *testing.T) {
	cases := []struct {
		in, want int
	}{
		{12, 12},
		{123, 124},
		{-12, 12},
		{-11, 12},
		{0, 0},
	}
	for _, c := range cases {
		got := New(c.in)
		radius := got.Size()
		if radius != c.want {
			t.Errorf("%d == %d, want %d", c.in, radius, c.want)
		}
	}
}

func TestRadius(t *testing.T) {
	cases := []struct {
		in, want int
	}{
		{12, 6},
		{123, 62},
		{-12, 6},
		{-11, 6},
		{0, 0},
	}
	for _, c := range cases {
		got := New(c.in)
		radius := got.Radius()
		if radius != c.want {
			t.Errorf("%d == %d, want %d", c.in, radius, c.want)
		}
	}
}
