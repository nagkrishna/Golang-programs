package main

import "fmt"

func main() {
	var str string
	fmt.Println("Enter the string you want to reverse")
	fmt.Scanln(&str)
	len := len(str)
	slc := make([]string, len)
	for i := 0; i < len-1; i++ {
		slc[i] = str[len-1]
	}
}
