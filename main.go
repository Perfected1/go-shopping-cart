package main

import "fmt"

func main() {
	fmt.Println("Simple Shopping Cart Started")

	cart := Cart{}

	cart.AddItem(Item{
		Name:  "Book",
		Price: 10.5,
		Qty:   2,
	})

	cart.AddItem(Item{
		Name:  "Pen",
		Price: 2.0,
		Qty:   5,
	})

	cart.ViewCart()
}