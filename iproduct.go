package main

type iProduct interface {
	calculatePrice() int
}

type ProductDetails struct {
	title string
	price float64
	brand string
}
