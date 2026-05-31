package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Entry point of the shopping cart app
	fmt.Println("Welcome to Simple Shopping Cart CLI")

	cart := Cart{}
	reader := bufio.NewReader(os.Stdin)

	for {
		// Show menu options
		fmt.Println("\nChoose an option:")
		fmt.Println("1 - Add Item")
		fmt.Println("2 - Remove Item")
		fmt.Println("3 - View Cart")
		fmt.Println("4 - Exit")
		fmt.Print("Enter choice: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {

		case "1":
			// Add item flow
			fmt.Print("Enter item name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Print("Enter price: ")
			priceInput, _ := reader.ReadString('\n')
			priceInput = strings.TrimSpace(priceInput)
			price, _ := strconv.ParseFloat(priceInput, 64)

			fmt.Print("Enter quantity: ")
			qtyInput, _ := reader.ReadString('\n')
			qtyInput = strings.TrimSpace(qtyInput)
			qty, _ := strconv.Atoi(qtyInput)

			cart.AddItem(Item{
				Name:  name,
				Price: price,
				Qty:   qty,
			})

		case "2":
			// Remove item flow
			fmt.Print("Enter item name to remove: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			cart.RemoveItem(name)

		case "3":
			// View cart
			cart.ViewCart()

		case "4":
			// Exit app
			fmt.Println("Exiting... Goodbye 👋")
			return

		default:
			fmt.Println("Invalid option. Try again.")
		}
	}
}