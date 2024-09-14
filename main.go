package main

import "fmt"

type Purchaseable interface {
	calculatePrice() float64
}

type CartItem struct {
	Purchaseable
	quantity int
}

type Cart struct {
	totalItems int
	total      float64
	items      []CartItem
}

var cart []CartItem

func addToCart(products ...CartItem) {
	cart = append(cart, products...)
}

func getCartDetails() Cart {
	var total_items int
	var cost float64

	for _, cartItem := range cart {
		cost += cartItem.calculatePrice() * float64(cartItem.quantity)
		total_items += cartItem.quantity
	}

	return Cart{total_items, cost, cart}
}

func main() {
	myShoe := Shoe{ProductDetails{title: "Nike AIR", brand: "Nike", price: 5000.0}, "43"}
	yoga := Computer{ProductDetails{title: "Yoga", brand: "Thinkpad", price: 8000.00}, "14", "256 SSD"}
	dell := Computer{ProductDetails{title: "Latitude 5490", brand: "Dell", price: 60000.00}, "14", "256 SSD"}

	addToCart(CartItem{myShoe, 2}, CartItem{yoga, 20}, CartItem{dell, 5})

	fmt.Printf("%#v", getCartDetails())
}
