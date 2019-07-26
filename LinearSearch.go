package main

import "fmt"

func LinearSearch() {
	var Slicer []int
	var num, num2, count int
	fmt.Print("How many Integers you want to search from:")
	fmt.Scanln(&num)
	Slicer = make([]int, num)
	fmt.Println("Enter", num, "Integers:")
	for i := 0; i < num; i++ {
		fmt.Scanln(&Slicer[i])
	}
	fmt.Print("Enter the number you want to search:")
	fmt.Scanln(&num2)
	for i := 0; i < len(Slicer); i++ {
		if Slicer[i] == num2 {
			fmt.Println("The Number", num2, "is found at Index", i, ".")
		} else {
			count++
		}
	}
	if count == len(Slicer) {
		fmt.Println("Number not Found!")
	}

}
