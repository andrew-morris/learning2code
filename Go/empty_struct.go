package main

import (
	"encoding/json"
	"fmt"
)

type Pizza struct {
	Toppings []string
	Size     string
	Inches   int
	Vendor   string
	Delivery bool
	Other	 map[string]string
}

func main() {

	m := map[string]string{"ooga":"booga"}

	Dinner := Pizza{
		Toppings: []string{"pepperoni", "cheese", "pineapple"},
		//Size:		"Large",
		Inches:   32,
		Vendor:   "PizzaHut",
		Delivery: true,
		Other:	  m,
	}

	JsonBlerb, _ := json.Marshal(Dinner)

	fmt.Println(string(JsonBlerb))
}
