package main

import (
	"fmt"
	"sort"
)

var IntsSlice []int
var FloatsSlice []float64
var StringsSlice []string

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
	IntsSlice = make([]int, num)
	fmt.Print("Enter", num, "Integers to sort:")
	for i = 0; i < num; i++ {
		fmt.Scanln(&IntsSlice[i])
	}
	fmt.Println("The sorted Integers are:")
	sort.Ints(IntsSlice)
	fmt.Println(IntsSlice)
	/*fmt.Println("The Reverse order is:")
	s := IntsSlice
	sort.Sort(sort.Reverse(sort.IntSlice[s]))
	fmt.Println(IntsSlice)*/
}

func FloatSort() {
	var num, i int
	fmt.Println("How many Float Numbers you want to sort:")
	fmt.Scanln(&num)
	FloatsSlice = make([]float64, num)
	fmt.Println("Enter", num, "Float Numbers to sort:")
	for i = 0; i < num; i++ {
		fmt.Scanln(&FloatsSlice[i])
	}
	fmt.Println("The sorted Float Numbers are:")
	sort.Float64s(FloatsSlice)
	fmt.Println(FloatsSlice)
	/*fmt.Println("The Reverse order is:")
	sort.Sort(sort.Reverse(sort.Float64Slice[FloatsSlice]))
	fmt.Println(FloatsSlice)*/
}

func StringSort() {
	var num, i int
	fmt.Println("How many Strings you want to sort:")
	fmt.Scanln(&num)
	StringsSlice = make([]string, num)
	fmt.Println("Enter", num, "Strings to sort:")
	for i = 0; i < num; i++ {
		fmt.Scanln(&StringsSlice[i])
	}
	fmt.Println("The sorted Strings are:")
	sort.Strings(StringsSlice)
	fmt.Println(StringsSlice)
	//fmt.Println("The Reverse order is:")
	//sort.Sort(sort.Reverse(sort.StringSlice[StringsSlice]))
	//fmt.Println(StringsSlice)
}

func IntegerSearch() {
	var num int
	fmt.Print("Enter an Integer Number to search:")
	fmt.Scanln(&num)
	fmt.Println(sort.SearchInts(IntsSlice, num))
}

func FloatSearch() {
	var num float64
	fmt.Print("Enter a Float Number to search:")
	fmt.Scanln(&num)
	fmt.Println(sort.SearchFloat64s(FloatsSlice, num))
}

func StringSearch() {
	var num string
	fmt.Print("Enter a String to search:")
	fmt.Scanln(&num)
	fmt.Println(sort.SearchStrings(StringsSlice, num))
}
