package polyclip

import (
	"reflect"
	"testing"
)

func TestClipLine(t *testing.T) {
	subject := Polygon{{
		{0, 1}, {1.25, 1}, {1.5, 1.1}, {1.75, 1}, {5, 1},
		{5, 2}, {0, 2},
	}}
	clipping := Polygon{
		{{1, 0}, {4, 0}, {4, 3}, {1, 3}},
		{{2, 0.5}, {3, 0.5}, {3, 2.5}, {2, 2.5}},
	}
	r := subject.Construct(CLIPLINE, clipping)
	want := Polygon{
		Contour{Point{X: 2, Y: 1}, Point{X: 1.75, Y: 1}, Point{X: 1.5, Y: 1.1}, Point{X: 1.25, Y: 1}, Point{X: 1, Y: 1}},
		Contour{Point{X: 2, Y: 2}, Point{X: 1, Y: 2}},
		Contour{Point{X: 4, Y: 1}, Point{X: 3, Y: 1}},
		Contour{Point{X: 4, Y: 2}, Point{X: 3, Y: 2}},
	}
	if !reflect.DeepEqual(r, want) {
		t.Errorf("%+v != %+v", r, want)
	}
}
