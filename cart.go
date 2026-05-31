package main

import "fmt"

// Cart stores all shopping items
type Cart struct {
	Items []Item
}

// AddItem adds a new item into the cart
func (c *Cart) AddItem(item Item) {
	c.Items = append(c.Items, item)
	fmt.Println("Item added:", item.Name)
}

// RemoveItem removes an item by name
func (c *Cart) RemoveItem(name string) {
	found := false

	for i, item := range c.Items {
		if item.Name == name {
			c.Items = append(c.Items[:i], c.Items[i+1:]...)
			found = true
			fmt.Println("Item removed:", name)
			break
		}
	}

	if !found {
		fmt.Println("Item not found:", name)
	}
}

// GetTotal calculates total cost of cart
func (c *Cart) GetTotal() float64 {
	total := 0.0

	for _, item := range c.Items {
		total += item.Price * float64(item.Qty)
	}

	return total
}

// ViewCart displays all items in a readable format
func (c *Cart) ViewCart() {
	if len(c.Items) == 0 {
		fmt.Println("Cart is empty")
		return
	}

	fmt.Println("\n--- Your Cart ---")

	for i, item := range c.Items {
		subtotal := item.Price * float64(item.Qty)

		fmt.Printf("%d. %s | Price: %.2f | Qty: %d | Subtotal: %.2f\n",
			i+1,
			item.Name,
			item.Price,
			item.Qty,
			subtotal,
		)
	}

	fmt.Printf("\nTotal: %.2f\n", c.GetTotal())
}