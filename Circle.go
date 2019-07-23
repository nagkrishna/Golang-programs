package main

import "fmt"

func Circle() {
	var radius float64
	const Pi = 3.14
	fmt.Print("Enter radius of the circle:")
	fmt.Scanln(&radius)
	fmt.Println("Area of the Circle is:", Pi*(radius*radius))
	fmt.Println("Circumference of the Circle is:", 2*Pi*radius)
}
