package main

import (
	"fmt"
	"sort"
)

var IntSlice []int
var FloatSlice []float64
var StringSlice []string

func SortEx() {
Loop:
	for {
		var num int
		fmt.Print("1.Sorting\n2.Searching\n3.Exit\n")
		fmt.Print("Which operation you want to perform:")
		fmt.Scanln(&num)
		switch num {
		case 1:
			Sorting()
		case 2:
			Searching()
		case 3:
			break Loop
		default:
			fmt.Println("Wrong Input!")
		}
	}
}

func Sorting() {
	var num int
	fmt.Println("1.Sort Integers\n2.Sort Floats\n3.Sort Strings")
	fmt.Print("Which Sorting operation you want to perform:")
	fmt.Scanln(&num)
	switch num {
	case 1:
		IntegerSort()
	case 2:
		FloatSort()
	case 3:
		StringSort()
	default:
		fmt.Println("Wrong Input!")
	}

}

func Searching() {
	var num int
	fmt.Println("1.Search an Integer\n2.Search a Float Number\n3.Search a String")
	fmt.Print("Which Search operation you want to perform:")
	fmt.Scanln(&num)
	switch num {
	case 1:
		IntegerSearch()
	case 2:
		FloatSearch()
	case 3:
		StringSearch()
	default:
		fmt.Println("Wrong Input!")
	}
}

func IntegerSort() {
	var num, i int
	fmt.Print("How many Integers you want to sort:")
	fmt.Scanln(&num)
	IntSlice = make([]int, num)
	fmt.Print("Enter ", num, " Integers to sort:")
	for i = 0; i < num; i++ {
		fmt.Scanln(&IntSlice[i])
	}
	fmt.Println("The sorted Integers are:")
	sort.Ints(IntSlice)
	fmt.Println(IntSlice)
}

func FloatSort() {
	var num, i int
	fmt.Println("How many Float Numbers you want to sort:")
	fmt.Scanln(&num)
	FloatSlice = make([]float64, num)
	fmt.Println("Enter", num, "Float Numbers to sort:")
	for i = 0; i < num; i++ {
		fmt.Scanln(&FloatSlice[i])
	}
	fmt.Println("The sorted Float Numbers are:")
	sort.Float64s(FloatSlice)
	fmt.Println(FloatSlice)
}

func StringSort() {
	var num, i int
	fmt.Println("How many Strings you want to sort:")
	fmt.Scanln(&num)
	StringSlice = make([]string, num)
	fmt.Println("Enter", num, "Strings to sort:")
	for i = 0; i < num; i++ {
		fmt.Scanln(&StringSlice[i])
	}
	fmt.Println("The sorted Strings are:")
	sort.Strings(StringSlice)
	fmt.Println(StringSlice)
}

func IntegerSearch() {
	var num int
	fmt.Print("Enter an Integer Number to search:")
	fmt.Scanln(&num)
	fmt.Println(sort.SearchInts(IntSlice, num))
}

func FloatSearch() {
	var num float64
	fmt.Print("Enter a Float Number to search:")
	fmt.Scanln(&num)
	fmt.Println(sort.SearchFloat64s(FloatSlice, num))
}

func StringSearch() {
	var num string
	fmt.Print("Enter an Integer Number to search:")
	fmt.Scanln(&num)
	fmt.Println(sort.SearchStrings(StringSlice, num))
}
