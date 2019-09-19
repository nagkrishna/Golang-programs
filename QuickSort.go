package main

import "fmt"

var Slicer []int

func QuickSort() {
	var num int
	fmt.Print("How many Integers you want to sort:")
	fmt.Scanln(&num)
	Slicer = make([]int, num)
	fmt.Println("Enter", num, "Integers:")
	for i := 0; i < num; i++ {
		fmt.Scanln(&Slicer[i])
	}
	Sorting1(0, len(Slicer)-1)
	fmt.Println("The Sorted array using QuickSort:")
	fmt.Println(Slicer)
}

func Sorting1(low int, high int) {
	pivot := high
	hIndex := high - 1
Loop:
	for {
		if Slicer[low] > Slicer[pivot] && Slicer[hIndex] < Slicer[pivot] {
			temp := Slicer[low]
			Slicer[low] = Slicer[hIndex]
			Slicer[hIndex] = temp
			continue
		} else if Slicer[low] < Slicer[pivot] {
			low++
			continue
		} else if Slicer[hIndex] > Slicer[pivot] {
			hIndex++
			continue
		} else if low == hIndex || hIndex < low {
			temp := Slicer[pivot]
			Slicer[pivot] = Slicer[low]
			Slicer[low] = temp
			break Loop
		}
	}
	Sorting1(0, low-1)
	Sorting1(low+1, pivot)
}
