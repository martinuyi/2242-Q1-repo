package main

import (
	"fmt"
)

// Create Function square() that returns area and perimeter of a square
func square(side float64) (area float64, perimeter float64) {

	area = side * 4
	perimeter = side * 4
	return area, perimeter
}

func main() {
	//Function call
	area, perimeter := square(4.0)
	fmt.Println(area, perimeter)

}
