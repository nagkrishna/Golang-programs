package main

import "fmt"

func Modify() {
	var Id1, x, Id2 int
	var Name string
	fmt.Println("Enter the Id of the account you want to modify:")
	fmt.Scanln(&Id1)
	fmt.Println("Which field you want to modify:")
	fmt.Printf("1.Name\n2.Id\n3.Balance\n")
	fmt.Scanln(&x)
	switch x {
	case 1:
		fmt.Println("Enter new name for the account with Id:", Id1)
		fmt.Scanln(&Name)
		mod := m[Id1]
		m[Id1] = Account{Name, mod.Balance}
		fmt.Println("Name is successfully changed.")
	case 2:
		fmt.Println("Enter new Id for the account with Id:", Id1)
		fmt.Scanln(&Id2)
		m[Id2] = m[Id1]
		delete(m, Id1)
		fmt.Println("Id is successfully changed.")
	case 3:
		fmt.Println("Enter new Balance for the account with Id:", Id1)
		fmt.Scanln(&x)
		mod := m[Id1]
		m[Id1] = Account{mod.Name, x}
		fmt.Println("Balance is successfully changed.")
	}
}
