package main

type Computer struct {
	ProductDetails
	screen_size string
	storage     string
}

func (c Computer) calculatePrice() float64 {
	return float64(c.price)
}
