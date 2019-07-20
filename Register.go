package main

import "fmt"

func Register() {
	var Name string
	var Id, Bal int
	fmt.Println("Enter the Name of the new account:")
	fmt.Scanln(&Name)
	fmt.Println("Enter Id of the new account:")
	fmt.Scanln(&Id)
	fmt.Println("Enter Balance of the new account:")
	fmt.Scanln(&Bal)
	m[Id] = Account{Name, Bal}
}
