package main

import "fmt"

type Cart struct {
	Items []Item
}

func (c *Cart) AddItem(item Item) {
	c.Items = append(c.Items, item)
}

// RemoveItem removes an item by index (simple version for now)
func (c *Cart) RemoveItem(index int) {
	if index < 0 || index >= len(c.Items) {
		fmt.Println("Invalid item index")
		return
	}

	c.Items = append(c.Items[:index], c.Items[index+1:]...)
	fmt.Println("Item removed successfully")
}

// GetTotal calculates total cart value
func (c *Cart) GetTotal() float64 {
	total := 0.0

	for _, item := range c.Items {
		total += item.Price * float64(item.Qty)
	}

	return total
}

// ViewCart prints items + total
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

	fmt.Printf("\nTotal: %.2f\n", c.GetTotal())
}