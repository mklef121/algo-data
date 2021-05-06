package main

import "fmt"

var taxTotal float64 = .0

// nil is not a type but a special value in Go. It represents an empty value of no type.
func main() {
	var message *string
	if message == nil {
		fmt.Println("error, unexpected nil value")
		// return
	}

	// fmt.Println(*message)
	products := loadProducts()
	for _, product := range products {
		taxTotal += salesTax(product.price, product.percentage)
	}

	fmt.Println("The total Tax is ", taxTotal)
}

// In this activity, we create a shopping cart application, where sales tax must be added to calculate the total:
// 1. Create a calculator that calculates the sales tax for a single item.
// 2. The calculator must take the items cost and its sales tax rate.
// 3. Sum the sales tax and print the total amount of sales tax required for the following items:

type productPrice struct {
	name       string
	price      float64
	percentage float64
}

func loadProducts() []productPrice {

	allProducts := []productPrice{}

	fmt.Println("enter", allProducts)
	// hi := new(productPrice)
	//cake
	cake := productPrice{name: "Cake", price: 0.99, percentage: 7.5}

	//milk
	milk := productPrice{name: "Milk", price: 2.75, percentage: 1.5}

	//Butter
	butter := productPrice{name: "Butter", price: 0.87, percentage: 2}

	allProducts = append(allProducts, cake)
	allProducts = append(allProducts, milk)
	allProducts = append(allProducts, butter)
	// allProducts = append(allProducts, productPrice{name: "Baller", price: 45.66})

	return allProducts
}

func salesTax(cost float64, taxRate float64) float64 {
	return cost * taxRate
}
