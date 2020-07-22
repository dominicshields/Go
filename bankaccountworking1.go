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
	fmt.Printf("accounts after:  %v\n", account)

	// Slice passed by value
	var slice1 []int = []int{1, 2, 3}
	//	fmt.Printf("slice1 before: %v\n", slice1)
	//	sliceByVal(slice1)
	//	fmt.Printf("slice1 after:  %v\n", slice1)

	// Slice passed by reference
	var slice2 []int = []int{1, 2, 3}
	fmt.Printf("By ref slice2 before: %v\n", slice2)
	sliceByRef(&slice2)
	fmt.Printf("By ref slice2 after:  %v\n", slice2)

	// Slice backed by array passed by value
	//	var array1 [6]int = [6]int{0, 1, 2, 3, 5, 6}
	//	slice3 := array1[1:4]
	//	fmt.Printf("array1 before: %v\n", array1)
	//	sliceByVal(slice3)
	//	fmt.Printf("array1 after:  %v\n", array1)

	// Slice backed by array passed by reference
	//	var array2 [6]int = [6]int{0, 1, 2, 3, 5, 6}
	//	slice4 := array2[0:6]
	//	fmt.Printf("Slice backed by array passed by reference array2 before: %v\n", array2)
	//	sliceByRef(&slice4)
	//	fmt.Printf("Slice backed by array passed by reference array2 after:  %v\n", array2)

	//Modify the contents of a slice passed by value
	fmt.Printf("Modify the contents of a slice passed by value slice1 before: %v\n", slice1)
	sliceByVal2(slice1)
	fmt.Printf("Modify the contents of a slice passed by value slice1 after:  %v\n", slice1)

	// Slice Array passed by reference
	fmt.Printf("Slice Array By Ref slice2 before: %v\n", slice2)
	sliceByRef2(&slice2)
	fmt.Printf("Slice Array By ref slice2 after:  %v\n", slice2)

}

func openacc(account_holder string, sort_code string, account_number string, account *[]accounts) {

	*account = append(*account, accounts{account_holder, sort_code, account_number, INIT_BALANCE, INIT_STATUS})
	//	for x := range account {
	//		fmt.Printf("%v %v %v %v %v\n", account[x].account_holder, account[x].sort_code, account[x].account_no, account[x].balance, account[x].status)
	//	}
	//return account
}

func sliceByVal(slice []int) {
	slice = append(slice, 4)
}

func sliceByRef(slice *[]int) {
	*slice = append(*slice, 4)
}

func sliceByVal2(slice []int) {
	slice[0] = 42
}

func sliceByRef2(slice *[]int) {
	(*slice)[0] = 1337
}
