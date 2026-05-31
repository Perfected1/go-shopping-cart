package main

import "fmt"

type Cart struct {
	Items []Item
}

func (c *Cart) AddItem(item Item) {
	c.Items = append(c.Items, item)
}

// ViewCart prints all items in a readable format
func (c *Cart) ViewCart() {
	if len(c.Items) == 0 {
		fmt.Println("Cart is empty")
		return
	}

	fmt.Println("\n--- Your Cart ---")
	for i, item := range c.Items {
		fmt.Printf("%d. %s | Price: %.2f | Qty: %d | Subtotal: %.2f\n",
			i+1,
			item.Name,
			item.Price,
			item.Qty,
			item.Price*float64(item.Qty),
		)
	}
}