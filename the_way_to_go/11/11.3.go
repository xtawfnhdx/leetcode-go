package main

import "fmt"

type stockPosition struct {
	ticker     string
	sharePrice float32
	count      float32
}

func (p stockPosition) getValue() float32 {
	return p.sharePrice * p.count
}

type car struct {
	make  string
	model string
	price float32
}

func (c car) getValue() float32 {
	return c.price
}

type valuable interface {
	getValue() float32
}

func showValue(asset valuable) {
	fmt.Printf("price is %f\n", asset.getValue())

}
func main() {
	var o valuable = stockPosition{"good", 534.2, 3}
	showValue(o)
	o = car{"ABC", "M4", 564334.43}
	showValue(o)
}
