package main

import "fmt"

func Transfer() {
	var Id1, trans, Id2 int
	fmt.Println("Enter Id of the account from which you want to tranfer money from:")
	fmt.Scanln(&Id1)
	fmt.Println("You have", m[Id1].Balance, "in your", m[Id1].Name, "account.")
	fmt.Println("Enter how much amount you want to tranfer:")
	fmt.Scanln(&trans)
	fmt.Println("Enter Id of the account you want to tranfer your money to:")
	fmt.Scanln(&Id2)
	a := m[Id1]
	b := m[Id2]
	a.Balance = a.Balance - trans
	b.Balance = b.Balance + trans
	m[Id1] = Account{a.Name, a.Balance}
	m[Id2] = Account{b.Name, b.Balance}
	fmt.Println("Amount tranfer is successfully executed.")
}
