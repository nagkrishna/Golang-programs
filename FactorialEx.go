package main

import "fmt"

func FactorialEx() {
	var num int
	fmt.Println("Program to find Factorial of a number.")
	fmt.Print("Enter a number:")
	fmt.Scanln(&num)
	sum := 1
	for i := 1; i <= num; i++ {
		sum *= i
	}
	fmt.Println("Factorial of", num, ":", sum)
}
