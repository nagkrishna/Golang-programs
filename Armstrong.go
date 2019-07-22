package main

import (
	"fmt"
	"math"
)

func Armstrong() {
	for {
		var num uint
		var count, sum float64
		fmt.Println("Enter a number to check if its armstrong number or not:")
		fmt.Scanln(&num)
		if num >= 0 {
			num1 := num
			for num1 != 0 {
				count++
				num1 /= 10
			}
			num2 := num
			if count >= 3 {
				for num2 != 0 {
					num3 := num2 % 10
					num4 := float64(num3)
					num2 /= 10
					sum = sum + math.Pow(num4, count)
				}
			}
		}
		if num == 0 || num == 1 || num == uint(sum) {
			fmt.Println(num, "is an Armstrong Number.")
		} else if num < 0 {
			fmt.Println("Armstrong Number cannot be Negative.")
		} else {
			fmt.Println(num, "is not an Armstrong Number.")
		}
	}
}
