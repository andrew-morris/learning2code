package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	m := make(map[string]string)
	m["name"] = "andrew"
	m["hometown"] = "columbia, sc"
	m["pet"] = "vinnie"

	jsonBlerb, _ := json.Marshal(m)

	fmt.Println(string(jsonBlerb))
}
