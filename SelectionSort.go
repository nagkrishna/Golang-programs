package main

import "fmt"

func SelectionSort() {
	var num, Index int
	var SelectSlicer []int
	fmt.Print("How many Integers you want to sort:")
	fmt.Scanln(&num)
	SelectSlicer = make([]int, num)
	fmt.Println("Enter", num, "Integers:")
	for i := 0; i < num; i++ {
		fmt.Scanln(&SelectSlicer[i])
	}
	for i := 0; i < len(SelectSlicer); i++ {
		less := SelectSlicer[i]
		for j := i; j < len(SelectSlicer); j++ {
			if SelectSlicer[j] < less {
				less = SelectSlicer[j]
				Index = j
			}
		}
		temp := SelectSlicer[i]
		SelectSlicer[i] = SelectSlicer[Index]
		SelectSlicer[Index] = temp
	}
	fmt.Print("The sorted array using SelectionSort:")
	fmt.Println(SelectSlicer)
}
