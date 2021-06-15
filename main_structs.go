package main

import "fmt"

func main() {
	mybill := newBill("marios bill")

	mybill.addItem("bawang mentah", 4.58)
	mybill.addItem("telor mentah", 5.60)
	mybill.addItem("sapi mentah", 10.28)
	mybill.updateTip(10)

	fmt.Println(mybill.format())
}
