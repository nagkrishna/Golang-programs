package main

import (
	"fmt"
	"sort"
)

func InterpolationSearch() {
	var Slicer1 []int
	var num, num2, pos, k int
	fmt.Print("How many Integers you want to search from:")
	fmt.Scanln(&num)
	Slicer1 = make([]int, num)
	fmt.Println("Enter", num, "Integers:")
	for i := 0; i < num; i++ {
		fmt.Scanln(&Slicer1[i])
	}
	sort.Ints(Slicer1)
	fmt.Print("Enter the number you want to search:")
	fmt.Scanln(&num2)
	j := 0
	k = len(Slicer1) - 1
Loop:
	for {
		pos = int((float32(num2-Slicer1[j]) / float32(Slicer1[k]-Slicer1[j])) * float32(k-j))
		pos = pos + j
		if pos > len(Slicer1)-1 {
			fmt.Println("Number not Found in the Array.")
			break Loop
		} else if num2 == Slicer1[pos] {
			fmt.Println("Number", num2, "is Found at Index", pos, "in the Array.")
			break Loop
		} else if Slicer1[pos] > num2 {
			j = 0
			k = pos - 1
		} else if Slicer1[pos] < num2 {
			j = pos + 1
			k = len(Slicer1) - 1
		} else if k+1 == 0 || j == len(Slicer1)-1 {
			fmt.Println("Number not Found in the Array.")
			break Loop
		}
	}

}
