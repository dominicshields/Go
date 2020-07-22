package main

import (
	"fmt"
)

type Account struct {
	account_holder string
	sort_code      string
	account_no     string
	balance        float64
	status         string
}

type Accounts struct {
	Items []Account
}

func main() {
	item1 := Account{account_holder: "Dominic", sort_code: "30-90-98", account_no: "0123456789", balance: 0.00}
	item2 := Account{account_holder: "Fred", account_no: "1234567890"}

	items := []Account{}
	NewAccount := Accounts{items}

	NewAccount.AddAccount(item1)
	NewAccount.AddAccount(item2)

	fmt.Println(NewAccount.Items)

	//	fmt.Println(NewAccount.Items[1].account_holder)

}
func (NewAccount *Accounts) AddAccount(item Account) []Account {
	NewAccount.Items = append(NewAccount.Items, item)
	return NewAccount.Items
}
