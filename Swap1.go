package main

import "fmt"

func Swap1() {
	var num1, num2 int
	fmt.Println("Swapping two numbers without using third variable.")
	fmt.Print("Enter First number:")
	fmt.Scanln(&num1)
	fmt.Print("Enter Second number:")
	fmt.Scanln(&num2)
	fmt.Println("Before Swapping:")
	fmt.Println("First Number:", num1, "Second Number:", num2)
	num1 = num1 + num2
	num2 = num1 - num2
	num1 = num1 - num2
	fmt.Println("After Swapping:")
	fmt.Println("First Number:", num1, "Second Number:", num2)
}
