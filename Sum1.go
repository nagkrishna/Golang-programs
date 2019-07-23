package main

import "fmt"

func Sum1() {
	var num, sum, i uint
	fmt.Print("Enter till which number you want to perform Addition operation of natural numbers:")
	fmt.Scanln(&num)
	for i = 1; i <= num; i++ {
		sum += i
	}
	fmt.Println("Addition of Natural numbers till", num, "are:", sum)
}
