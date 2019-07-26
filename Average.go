package main

import "fmt"

func Average() {
	var num int
	var sum float64
	fmt.Print("Enter how many numbers you want to find Average for:")
	fmt.Scanln(&num)
	arr := make([]int, num)
	fmt.Println("Enter", num, "Interger values to find their Average:")
	for i := 0; i < num; i++ {
		fmt.Scanln(&arr[i])
	}
	for i := 0; i < num; i++ {
		sum += float64(arr[i])
	}
	fmt.Print("The Average is:", sum/float64(num))
}
