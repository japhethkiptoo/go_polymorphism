package main

type Shoe struct {
	ProductDetails
	size string
}

func (s Shoe) calculatePrice() float64 {
	return float64(s.price)
}
