package supermarket

import "fmt"

var product map[string]interface{}

func GetProduct(product map[string]interface{}) {
	fmt.Println("............Get the price of the item............")
	for {
		fmt.Println("Enter a item:")
		var item string
		_, err := fmt.Scan(&item)

		if err != nil {
			panic(err)
		}
		if element, ok := product[item]; ok {
			fmt.Println("The price of the", item, element)
			break
		} else {
			fmt.Printf("Item %v is not in the list\n", item)
			break
		}
	}
}

func AddProduct(product map[string]interface{}) {
	fmt.Println("............Add the new product to the list............")
	for {
		fmt.Println("Enter a item to add:")
		var item string
		_, err := fmt.Scan(&item)

		if err != nil {
			panic(err)
		}
		if element, ok := product[item]; ok {
			fmt.Println("The product is avalaible in the list with price", element)
			continue

		} else {
			fmt.Printf("The %v item is not present in the list, what is price of the item?\n", item)
			var value string
			fmt.Scan(&value)
			product[item] = value
		}
		fmt.Println("Data added into the list.")
		fmt.Println(product)
		break
	}

}

func UpdateProduct(product map[string]interface{}) {
	fmt.Println("............Update the price of the item............")

	for {
		fmt.Println("Enter a item to update:")
		var item string
		_, err := fmt.Scan(&item)
		if err != nil {
			panic(err)
		}
		if element, ok := product[item]; ok {
			fmt.Println("The price of the", item, element)
			fmt.Printf("Do you want to update the price of the item %v,Enter the price:\n ", item)
			var value string
			fmt.Scan(&value)
			product[item] = value
			fmt.Println(product)
			break

		}
	}
}

func DeleteProduct(product map[string]interface{}) {
	fmt.Println("............Deleting the item............")
	for {
		fmt.Println("Enter a item to delete:")
		var item string
		fmt.Scan(&item)

		delete(product, item)
		fmt.Printf("The item %v is deleted successfully\n", item)
		fmt.Println(product)
		break
	}
}
func Print(product map[string]interface{}) {
	fmt.Println(product)
}
