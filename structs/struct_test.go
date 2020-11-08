package structs

import (
    "testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("Got - %.2f want - %.2f", got, want)
	}

}


func TestArea(t *testing.T) {
    // t.Parallel()
    areaTest := []struct {
        shape Shape
        want float64
    }{
        {shape: Rectangle{12,6}, want: 72.0},
        {shape: Circle{10}, want: 314.1592653589793},
        {shape: Triangle{12, 6}, want: 36.0},
    }

    for _, tt := range areaTest{
        got := tt.shape.Area()
        
        if got != tt.want {
            t.Errorf("got %g want %g", got, tt.want)
        }
    }

}
