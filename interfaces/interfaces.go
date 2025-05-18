package main

import "fmt"

// interface
type paymenter interface {
	pay(amount float32)
}

type payment struct {
	gateway paymenter
}

func (p payment) makePayment(amount float32) {
	p.gateway.pay(amount)
}

// razorpay struct
type razorpay struct{}

// razorpay method
func (r razorpay) pay(amount float32) {
	fmt.Println("making payment using razorpay ", amount)
}

type strip struct{}

func (s strip) pay(amount float32) {
	fmt.Println("making payment using strip ", amount)
}

type paypal struct{}

func (p paypal) pay(amount float32) {
	fmt.Println("making payment using paypal ", amount)

}

func main() {
	// stripPaymnetGw := strip{}
	// razorpayPaymnetGw := razorpay{}
	paypalPaymnetGw := paypal{}
	newPayment := payment{
		gateway: paypalPaymnetGw,
	}
	newPayment.makePayment(100)
}
