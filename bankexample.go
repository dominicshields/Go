package main

import (
	"fmt"
)

type MyBoxItem struct {
	account_holder string
	sort_code      string
	account_no     string
	balance        float64
	status         string
}

type MyBox struct {
	Items []MyBoxItem
}

func (box *MyBox) AddItem(item MyBoxItem) []MyBoxItem {
	box.Items = append(box.Items, item)
	return box.Items
}

func main() {

	item1 := MyBoxItem{Name: "Test Item 1", Account: "0123456789"}
	item2 := MyBoxItem{Name: "Test Item 2", Account: "1234567890"}

	items := []MyBoxItem{}
	box := MyBox{items}

	box.AddItem(item1)
	box.AddItem(item2)

	fmt.Println(box.Items)
}
