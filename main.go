package main

import "fmt"

func main() {
	// Starting point of our shopping cart application
	fmt.Println("Simple Shopping Cart Started")

	// Create a new cart instance
	cart := Cart{}

	// Add sample items into cart (for testing purposes)
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

	cart.AddItem(Item{
		Name:  "Laptop Sticker",
		Price: 3.0,
		Qty:   1,
	})

	// View cart before removal
	cart.ViewCart()

	fmt.Println("\nRemoving Pen...")

	// Remove item by name
	cart.RemoveItem("Pen")

	// View updated cart
	cart.ViewCart()
}