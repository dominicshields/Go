package main

import (
	"fmt"
)

type accounts struct {
	account_holder string
	sort_code      string
	account_no     string
	balance        float64
	status         string
}

const INIT_STATUS = "Open"
const CLOSED_STATUS = "Closed"
const INIT_BALANCE = 0.00

func main() {
	account := []accounts{}
	// Open account with a balance of zero
	account = openacc("Dominic Shields", "33-93-91", "00929999", &account)
	//	account = openacc("John Jones", "33-93-91", "00929988")
	//	fmt.Println(account)

	// Close account
	closeacc("33-93-91", "00929999", &account)
	//	fmt.Println(account)

}

func openacc(account_holder string, sort_code string, account_number string, account *[]accounts) {

	*account = append(*account, accounts{account_holder, sort_code, account_number, INIT_BALANCE, INIT_STATUS})
	//	for x := range account {
	//		fmt.Printf("%v %v %v %v %v\n", account[x].account_holder, account[x].sort_code, account[x].account_no, account[x].balance, account[x].status)
	//	}
	//	return *account
}

func closeacc(sort_code string, account_number string, account *[]accounts) {
	fmt.Println(account)
	//	for x := range *account {
	//		if account.sort_code == sort_code && account.account_number == account_number {
	//			account.status = CLOSED_STATUS
	//		}
	//	}
}
