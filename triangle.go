package main

import (
	"fmt"
	"math"
)

// Create a struct Triangle
type Triangle struct {
	base   float64
	height float64
}

// Create a method area() called on struct Triangle
func (t Triangle) area() float64 {

	return (t.base * t.height) / 2

}

// Create a method perimeter() called on struct Triangle
func (t Triangle) perimeter() float64 {

	return t.base + t.height + (math.Sqrt((t.base * t.base) + (t.height * t.height)))

}

func main() {
	//Initializing a Triangle using SHN
	tr := Triangle{
		base:   5.0,
		height: 12,
	}
	// Functions call
	fmt.Println(tr.area())
	fmt.Println(tr.perimeter())
}
