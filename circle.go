package main

import (
	"fmt"
	"math"
)

// Create a struct Circle
type Circle struct {
	radius float64
}

// Create a method area() called on struct Circle
func (c Circle) area() float64 {

	return math.Pi * c.radius * c.radius

}

// Create a method perimeter() called on struct Circle
func (t Circle) perimeter() float64 {

	return 2 * math.Pi * t.radius
}

func main() {
	//Initialize a Circle using SHN
	circle := Circle{
		radius: 5.0,
	}
	// Functions call
	fmt.Println(circle.area())
	fmt.Println(circle.perimeter())
}
