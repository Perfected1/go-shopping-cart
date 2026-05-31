package main

import "fmt"

func main() {
	fmt.Println("Simple Shopping Cart Started")

	cart := Cart{}

	cart.AddItem(Item{"Book", 10.5, 2})
	cart.AddItem(Item{"Pen", 2.0, 5})
	cart.AddItem(Item{"Laptop Sticker", 3.0, 1})

	cart.ViewCart()

	fmt.Println("\nRemoving item at index 1...")
	cart.RemoveItem(1)

	cart.ViewCart()
}
