package main

import (
	"fmt"
	"sort"
)

func BinarySearch() {
	var Slicer, Slicer1 []int
	var num, num2, half int
	fmt.Print("How many Integers you want to search from:")
	fmt.Scanln(&num)
	Slicer = make([]int, num)
	fmt.Println("Enter", num, "Integers:")
	for i := 0; i < num; i++ {
		fmt.Scanln(&Slicer[i])
	}
	sort.Ints(Slicer)
	fmt.Print("Enter the number you want to search:")
	fmt.Scanln(&num2)
	Slicer1 = Slicer
Loop:
	for {
		half = len(Slicer1) / 2
		if half == 0 {
			fmt.Println("Number not Found!")
			break Loop
		} else if num2 == Slicer1[half] {
			fmt.Println("Number", num2, "is Found in the Array.")
			break Loop
		} else if num2 < Slicer1[half] {
			Slicer1 = Slicer1[:half]
		} else if num2 > Slicer1[half] {
			Slicer1 = Slicer1[half+1:]
		} else {
			continue
		}
	}
}
