package main

import "fmt"

func main() {
	fmt.Println("Simple Shopping Cart Started")

	cart := Cart{}

	item := Item{
		Name:  "Book",
		Price: 10.5,
		Qty:   2,
	}

	cart.Items = append(cart.Items, item)

	fmt.Println("Cart Items:", cart.Items)
}