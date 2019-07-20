package main

import "fmt"

func View() {
	var Id int
	fmt.Println("Enter the Id of the Account to view Name and Balance{Name Balance}:")
	fmt.Scanln(&Id)
	fmt.Println(m[Id])
}
