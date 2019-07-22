package main

import "fmt"

func Floyd() {
	var num int
	fmt.Println("Enter how many rows of Floyd's Triangle you want:")
	fmt.Scanln(&num)
	fmt.Println("The Floyd's Triangle with", num, "rows:")
	j := 1
	for i := 1; i <= num; i++ {
		for k := 1; k <= i; k++ {
			fmt.Print(j, "\t")
			j += 1
		}
		fmt.Println()
	}

}
