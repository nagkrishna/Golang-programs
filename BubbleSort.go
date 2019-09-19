package main

import "fmt"

func BubbleSort() {
	var num int
	var Slicer []int
	fmt.Print("How many Integers you want to sort:")
	fmt.Scanln(&num)
	Slicer = make([]int, num)
	fmt.Println("Enter", num, "Integers:")
	for i := 0; i < num; i++ {
		fmt.Scanln(&Slicer[i])
	}
	for i := 0; i < len(Slicer); i++ {
		for j := 0; j < len(Slicer)-1; j++ {
			if Slicer[j] > Slicer[j+1] {
				temp := Slicer[j+1]
				Slicer[j+1] = Slicer[j]
				Slicer[j] = temp
			}
		}
	}
	fmt.Print("Sorted Array using BubbleSort:")
	fmt.Println(Slicer)
}
