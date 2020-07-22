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
	var account []accounts
	openacc("Dominic Shields", "33-93-91", "00929999", &account)
	openacc("John Jones", "33-93-91", "00929988", &account)
	fmt.Printf("accounts after:  %v\n", account)

}

func openacc(account_holder string, sort_code string, account_number string, account *[]accounts) {

	*account = append(*account, accounts{account_holder, sort_code, account_number, INIT_BALANCE, INIT_STATUS})
}

func closeacc(sort_code string, account_number string, account *[]accounts) {

	for x := range []account {
		if account.sort_code == sort_code && account.account_number == account_number {
			account.status = CLOSED_STATUS
		}
	}
}
