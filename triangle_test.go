package main

import (
	"testing"
)

// Run test for Area()
func TestArea(t *testing.T) {
	tr := Triangle{5.0, 12}
	got := tr.area()
	expected := 30.0

	//Compare what I received with wahat I expected
	if got != expected {
		t.Errorf("got %g expected %g", got, expected)
	}

}

//Run test for Perimeter()

func TestPerimeter(t *testing.T) {
	tr := Triangle{5.0, 12}
	got := tr.perimeter()
	expected := 30.0

	//Compare what I received with wahat I expected
	if got != expected {
		t.Errorf("got %g expected %g", got, expected)
	}

}
