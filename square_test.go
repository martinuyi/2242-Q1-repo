package main

import (
	"testing"
)

// Run test for square()
func TestSquare(t *testing.T) {
	//call the return values of area and perimeter with variables aSq, Psq
	aSq, Psq := square(4.0)

	aSq_expected := 16.0
	Psq_expected := 16.0

	//Test for area and perimeter returned
	if (aSq_expected != aSq) || (Psq_expected != Psq) {
		if aSq_expected != aSq {
			t.Errorf("got %f, expected %f", aSq, aSq_expected)
		}
	}
	if Psq_expected != Psq {
		t.Errorf("got %f, expected %f", Psq, Psq_expected)
	}
}
