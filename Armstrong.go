package main

import (
	"fmt"
	"math"
)

var num, num1, num2, num3 uint
var count, sum float64

func Armstrong() {
	fmt.Println("Enter a number to check if its armstrong number or not:")
	fmt.Scanln(&num)
	if num != 0 && num > 0 {
		num1 = num
		for num1 != 0 {
			num1 /= 10
			count++
		}
		num2 = num
		for num2 != 0 {
			num3 = num2 % 10
			num4 := float64(num3)
			num2 /= 10
			sum = sum + math.Pow(num4, count)
		}
	} else if num == 0 || num < 0 {
		fmt.Println("Armstrong Number cannot be 0 or Negative.")
	}
	if sum == float64(num) {
		fmt.Println(num, "is an Armstrong Number.")
	} else {
		fmt.Println(num, "is not an Armstrong Number.")
	}
}
