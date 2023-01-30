package main

import (
	"testing"
)

// Run test for Area()
func TestArea(t *testing.T) {
	circle := Circle{5.0}
	got := circle.area()
	expected := 78.53981633974483

	//Compare what I received with wahat I expected
	if got != expected {
		t.Errorf("got %g expected %g", got, expected)
	}

}

//Run test for Perimeter()

func TestPerimeter(t *testing.T) {
	circle := Circle{5.0}
	got := circle.perimeter()
	expected := 31.41592653589793

	//Compare what I received with wahat I expected
	if got != expected {
		t.Errorf("got %g expected %g", got, expected)
	}

}
