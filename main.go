package main

import "fmt"

func main() {
	fmt.Println("Simple Shopping Cart Started")

	cart := Cart{}

	item1 := Item{
		Name:  "Book",
		Price: 10.5,
		Qty:   2,
	}

	item2 := Item{
		Name:  "Pen",
		Price: 2.0,
		Qty:   5,
	}

	cart.AddItem(item1)
	cart.AddItem(item2)

	fmt.Println("Cart Items:", cart.Items)
}