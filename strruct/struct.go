package main

import (
	"fmt"
	"time"
)

// order struct

type customer struct {
	name  string
	phone string
}

// composition
type order struct {
	id      string
	amount  float32
	status  string
	creatAt time.Time //nanosec precision
	customer
}

func newOrder(id string, amount float32, status string) *order {
	myOrder := order{
		id:     id,
		amount: amount,
		status: status,
	}
	return &myOrder
}

// reciver type
func (o *order) changeStatus(status string) {
	o.status = status
}
func (o *order) getAmount() float32 {
	return o.amount
}

func main() {
	// //if you don't save the value then the default value will be zero value
	// //int -> 0, float -> 0, string -> "", bool -> flase
	// myOrder := order{
	// 	id:     "1",
	// 	amount: 50.00,
	// 	status: "received",
	// }
	// myOrder.changeStatus("confirmend")

	// myOrder.creatAt = time.Now()

	// fmt.Println(myOrder.status)

	// myOrder2 := order{
	// 	id:      "2",
	// 	amount:  100,
	// 	status:  "delivered",
	// 	creatAt: time.Now(),
	// }
	// myOrder.status = "paid"
	// fmt.Println("order struct : ", myOrder)
	// fmt.Println(myOrder.getAmount())
	// fmt.Println("order struct : ", myOrder2)

	//2nd way
	// newOrder := newOrder("1", 30.50, "receuved")
	// fmt.Println(newOrder.amount)

	// language := struct {
	// 	name   string
	// 	isGood bool
	// }{"golang", true}
	// fmt.Println(language)

	newCustomer := customer{
		name:  "jhon",
		phone: "123456789",
	}
	newOrder := order{
		id:       "1",
		amount:   30,
		status:   "received",
		customer: newCustomer,
	}
	newOrder.customer.name = "rohit"
	fmt.Println(newOrder)
}
