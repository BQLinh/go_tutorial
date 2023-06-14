package main

import "fmt"

func main() {
	mybill := newBill("my bill")

	mybill.addItem("onion soup", 4.5)
	mybill.addItem("carrot", 2.5)
	mybill.addItem("chocolate", 3)

	mybill.updateTip(10)
	fmt.Println(mybill.format())
}
