//custome type definition,like classes

//struct like class ,collection of fields

package main

import (
	"fmt"
	"time"
)

type Order struct {
	id        int
	amount    float32
	status    string
	createdAt time.Time
}
type Products struct {
	name    string
	price   int
	company string
}

// define rule ,return type
func newProducts(name string, price int, company string) Products {
	p := Products{
		name:    name,
		price:   price,
		company: company,
	}
	return p //this is return a copy of the product
}

func main() {

	o := Order{
		id:     123,
		amount: 540,
		status: "Pending",
	}

	//define object for product
	p := Products{
		name:    "Laptop",
		price:   800000,
		company: "Lenvo",
	}
	fmt.Println("Order Struct", o) //Order Struct {123 540 Pending {0 0 <nil>}}
	//how to add value after set vlaue
	o.createdAt = time.Now()
	fmt.Println("order", o) //order {123 540 Pending {14007551585385843844 668501 0x7ff7cd274440}}
	fmt.Println(o.id)       //123

	o2 := Order{
		id:        1234,
		amount:    200,
		status:    "Deliverd",
		createdAt: time.Now(),
	}

	fmt.Println(p)      //{Laptop 800000 Lenvo}
	fmt.Println(p.name) //Laptop
	fmt.Println("My Order two", o2)

}
