package main

type Cart struct {
	Items []Item
}

// AddItem adds a new item into the cart
func (c *Cart) AddItem(item Item) {
	c.Items = append(c.Items, item)
}