package main

import (
	"fmt"
)

type Account struct {
	Name    string
	Balance int
}

var m map[int]Account

func main() {
	m = make(map[int]Account)
	var Name string
	var x, Id, Bal int
	fmt.Println("Create Default data for activating few functionalities of the program.")
	fmt.Println("Enter Name of your default account:")
	fmt.Scanln(&Name)
	fmt.Println("Enter Id of your default account(only intergers):")
	fmt.Scanln(&Id)
	fmt.Println("Enter Balance of your default account:")
	fmt.Scanln(&Bal)
	m[Id] = Account{Name, Bal}
Loop:
	for {
		fmt.Printf("\n1.Register/Create New Account\n2.View Account\n3.Modify Existing Account\n4.Transfer Balance\n5.Exit\n")
		fmt.Println("Enter your choice:")
		fmt.Scanln(&x)
		switch x {
		case 1:
			Register()
		case 2:
			View()
		case 3:
			Modify()
		case 4:
			Transfer()
		case 5:
			break Loop
		default:
			fmt.Println("Wrong Input!")
		}
	}
}
