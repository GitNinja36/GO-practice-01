package main

import "fmt"

// enumerated types

type OrderStatus int

const (
	Received OrderStatus = iota
	Confirmend
	Prepared
	Delivered
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("changing the order status to ", status)
}

func main() {
	changeOrderStatus(Prepared)
}
