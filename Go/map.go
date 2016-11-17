package main

import (
	"fmt"
)

func main() {
	m := make(map[string]string)
	m["name"] = "andrew"

	fmt.Println(m["name"])
}
