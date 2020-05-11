package main

import (
	"fmt"
	"github.com/shoppingapp/supermarket"
)

func main() {
	product := make(map[string]interface{})
	product["bread"] = 50.78
	product["butter"] = 40.90
	product["eggs"] = 60.90
	product["milk"] = 90.90

	for i := 0; i < 5; i++ {
		fmt.Println("enter the action to perform:")
		var i string
		fmt.Scan(&i)
		switch i {
		case "get":
			supermarket.GetProduct(product)

		case "add":
			supermarket.AddProduct(product)

		case "update":
			supermarket.UpdateProduct(product)

		case "delete":
			supermarket.DeleteProduct(product)

		case "print":
			supermarket.Print(product)

		default:
			fmt.Println("Invalid Action")
		}
	}
	// supermarket.GetProduct(product)
	// supermarket.AddProduct(product)
	// supermarket.UpdateProduct(product)
	// supermarket.DeleteProduct(product)
}
